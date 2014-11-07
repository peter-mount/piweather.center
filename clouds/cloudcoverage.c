
/*
 * A sensor which calculates the amount of cloud cover in the current image.
 * 
 * Unlike other sensors, this utilises the renderer api to analyse the image,
 * generating a 640x480 3 colour image:
 *   white - cloud
 *   blue  - sky
 *   black - none of the above (if detectable)
 * 
 * The sensor value will be the percentage of white/blue pixels which are white.
 */

#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "clouds/clouds.h"
#include "camera/camera.h"
#include "lib/string.h"
#include <gd.h>
#include <math.h>

/**
 * Used to create the sensor struct defining this sensor
 * @return 
 */
void register_cloud_sensor(CONFIG_SECTION *sect) {

    // Create the cloud virtual sensor
    struct cloudcoverage_sensor *cloud_sensor = (struct cloudcoverage_sensor *) malloc(sizeof (struct cloudcoverage_sensor));
    memset(cloud_sensor, 0, sizeof (struct cloudcoverage_sensor));
    cloud_sensor->sensor.name = genurl(sect->node.name, "/percent");
    cloud_sensor->sensor.title = "Cloud Coverage Percent";
    sensor_configure(sect, &cloud_sensor->sensor);

    // Create the okta virtual sensor
    struct sensor *okta_sensor = (struct sensor *) malloc(sizeof (struct sensor));
    memset(okta_sensor, 0, sizeof (struct sensor));
    okta_sensor->name = genurl(sect->node.name, "/okta");
    okta_sensor->title = "Cloud Coverage Okta";
    // Configure to get the schedule
    sensor_configure(sect, okta_sensor);
    // Override annotate
    config_getBooleanParameter(sect, "okta.annotate", &okta_sensor->annotate);

    // Now register them
    sensor_register(&cloud_sensor->sensor);
    sensor_register(okta_sensor);

    // Register the renderer only if sensor is enabled
    if (cloud_sensor->sensor.enabled) {
        // FIXME for now enable for every camera - probably not what we want
        struct Node *n = cameras.l_head;
        while (list_isNode(n)) {
            CAMERA camera = (CAMERA) n;
            n = n->n_succ;

            struct cloud_renderer *renderer = (struct cloud_renderer *) malloc(sizeof (struct cloud_renderer));
            memset(renderer, 0, sizeof (struct cloud_renderer));

            renderer->cloud_sensor = cloud_sensor;
            renderer->okta_sensor = okta_sensor;
            renderer->renderer.postinit = clouds_render_postinit;
            renderer->renderer.render = clouds_render;

            renderer->black = 0;
            config_getHexParameter(sect, "black", &renderer->black);

            renderer->cloud = 0x00FFFFFF;
            config_getHexParameter(sect, "cloud", &renderer->cloud);

            renderer->sky = 0x000000FF;
            config_getHexParameter(sect, "sky", &renderer->sky);

            renderer->rblim = 0.84;
            config_getDoubleParameter(sect, "rblim", &renderer->rblim);

            // This is optional
            config_getIntParameter(sect, "sun", &renderer->sunLimit);
            if (renderer->sunLimit < 0 || renderer->sunLimit > 255) {
                fprintf(stderr, "Invalid sun in %s: %d\n", sect->node.name, renderer->sunLimit);
                exit(1);
            }

            char *s = NULL;
            config_getCharParameter(sect, "mask", &s);
            if (s) {
                FILE *f = fopen(s, "r");
                if (f) {
                    if (strendswith(s, ".png"))
                        renderer->mask = gdImageCreateFromPng(f);
                    else
                        renderer->mask = gdImageCreateFromJpeg(f);
                    fclose(f);
                }
            }

            imagerenderer_register(camera, &renderer->renderer);
        }
    }
}