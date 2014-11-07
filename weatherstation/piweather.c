/* 
 * File:   main.c
 * Author: peter
 *
 * Created on February 5, 2014, 5:58 PM
 */

#include <signal.h>
#include "main.h"
#include "camera/camera.h"
#include "sensors/sensors.h"
#include "sensors/gpio/gpio.h"
#include "webserver/webserver.h"
#include "global_config.h"
#include "renderers/imagerenderer.h"
#include "clouds/clouds.h"
#include "scheduler/scheduler.h"
#include "logger/logger.h"
#include "piweather_build.h"

#ifdef HAVE_CURL
#include "curl/curl.h"
#endif

// These are the deployed sensor types.
extern void register_cpu_sensor(CONFIG_SECTION *sect);
extern void register_uptime_sensor(CONFIG_SECTION *sect);
extern void register_virtual_sensor(CONFIG_SECTION *sect);
extern void register_cloud_sensor(CONFIG_SECTION *sect);

#ifdef HAVE_I2C
extern void register_i2c_sensor(CONFIG_SECTION *sect);
extern void register_bh1750_sensor(CONFIG_SECTION *sect);
extern void register_si1145_sensor(CONFIG_SECTION *sect);
extern void register_adcpi1_sensor(CONFIG_SECTION *sect);
#endif

// Our available sensors
static struct sensor_registry deployable_sensors[] = {

    // The PI's CPU Temperature sensor
    {"cloudcoverage", register_cloud_sensor, "Cloud Coverage"},

    // The PI's CPU Temperature sensor
    {"cpu", register_cpu_sensor, "PI CPU Temperature"},

    // wiringPi based sensors
#ifdef HAVE_WIRING_PI
    {"dht22", register_dht22_sensor, "DHT22 Temperature/Humidity sensor"},
#endif

    // Generic i2c Sensor support
#ifdef HAVE_I2C
    {"i2c", register_i2c_sensor, "Generic i2c based sensor"},
    {"bh1750", register_bh1750_sensor, "BH1750 Ambient Light sensor"},
    {"si1145", register_si1145_sensor, "SI1145 UV/IR/Ambient Light sensor"},
    {"adcpiv1", register_adcpi1_sensor, "ADC PI Version 1 board"},
#endif

    // Uptime sensor
    {"uptime", register_uptime_sensor, "PI & Program uptime"},

    // Virtual sensors - this must be last as they refer to other sensors
    {"virtual", register_virtual_sensor, "Virtual (calculated) sensors"},

    // Terminates the list
    {NULL, NULL}
};

// Our loggers
extern void register_console_logger();
extern void register_file_logger();
extern void register_rest_logger();

#ifdef HAVE_CURL
extern void register_iotonl_logger();
#endif

#ifdef HAVE_RABBITMQ
extern void register_rabbitmq_logger();
#endif

// Mandatory renderers
extern struct image_renderer *create_annotatedrenderer();
extern struct image_renderer *create_rawrenderer();
extern struct image_renderer *create_thumbnailrenderer();

int verbose;

/**
 * Handler for sigint signals
 *
 * @param signal_number ID of incoming signal.
 *
 */
void signal_handler(int signal_number) {
    if (signal_number == SIGUSR1) {
        // Handle but ignore - prevents us dropping out if started in none-signal mode
        // and someone sends us the USR1 signal anyway
    } else {
        // Going to abort on all other signals
        fprintf(stderr, "Aborting program\n");
        exit(130);
    }
}

// Shows default help, mainly the available sensor's in this binary

static void showHelp() {
    fprintf(stderr, "Pi Weather Station\n\nCommand line arguments:\n");
    fprintf(stderr, "%8s %s\n", "-?", "Show help");
    fprintf(stderr, "%8s %s\n", "-v", "Verbose output on stderr");
    fprintf(stderr, "%8s %s\n", "-vv", "Debugging output on stderr");

    fprintf(stderr, "\nAvailable sensors:\n");
    struct sensor_registry *s = deployable_sensors;
    while (s->registry) {
        fprintf(stderr, " %16s %s\n", s->type, s->desc);
        s++;
    }
}

static void parseArgs(const int argc, const char** argv) {
    int i = 1;
    while (i < argc) {
        const char *arg = argv[i];
        if (strcmp("-?", arg) == 0) {
            showHelp();
            exit(0);
        } else if (strcmp("-v", arg) == 0) {
            verbose = 1;
            i++;
        } else if (strcmp("-vv", arg) == 0) {
            verbose = 2;
            i++;
        } else {
            fprintf(stderr, "Unsupported option: %s\n", arg);
            exit(1);
        }
    }
}

int main(const int argc, const char** argv) {

    // Our only parameters
    verbose = 0;

    parseArgs(argc, argv);

    config_parse_dir("/etc/weather");

    // Verbose give build details
    if (verbose > 1)
        fprintf(stderr, "Version " PIWEATHER_VERSION " (" PIWEATHER_BUILD_HOST ") " PIWEATHER_BUILD_TIME);

    astro_init();
    scheduler_init();

    // Register any loggers
    logger_init();
    register_console_logger();
    register_file_logger();
    register_rest_logger();

#ifdef HAVE_CURL
    curl_global_init(CURL_GLOBAL_ALL);
    register_iotonl_logger();
#endif

#ifdef HAVE_RABBITMQ
    register_rabbitmq_logger();
#endif

#ifdef HAVE_I2C
    i2c_init();
#endif

    gpio_init();

#ifdef HAVE_CAMERA
    camera_init();
    imagerenderer_initialise();
#endif

    // Register the sensors
    sensor_registerAll(deployable_sensors);
    sensor_init();

#ifdef HAVE_CAMERA
    // Our mandatory renderers, ensures they are run first being registered last
    struct Node *n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        // Put annotated last
        list_addTail(&camera->renderers.renderers, &create_annotatedrenderer()->node);
        // These go first
        list_addHead(&camera->renderers.renderers, &create_thumbnailrenderer()->node);
        list_addHead(&camera->renderers.renderers, &create_rawrenderer()->node);
    }
#endif

    signal(SIGINT, signal_handler);

    // Disable USR1 for the moment - may be reenabled if go in to signal capture mode
    signal(SIGUSR1, SIG_IGN);

    // Now start the system up
    webserver_initialise(config);

    logger_start();

    // Finish off configuring the webserver, default port etc
    webserver_set_defaults();

#ifdef HAVE_CAMERA
    n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        imagerenderer_init(camera);
        imagerenderer_postinit(camera);
        // Initial render ensures the test page is present until we get a real camera image
        imagerenderer_render(camera);
        imagerenderer_start(camera);
    }
#endif

    sensor_postinit();

    // Start everything up
    webserver_start();

    // Now go to sleep by waiting for the scheduler thread to die - which it shouldn't
    pthread_join(schedule.thread, NULL);

    // Shutdown - we shouldn't actually get here
    webserver_stop();
    logger_stop();

#ifdef HAVE_CURL
    curl_global_cleanup();
#endif

#ifdef HAVE_CAMERA
    camera_stop();
#endif

    return 0;
}

