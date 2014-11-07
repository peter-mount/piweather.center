/**
 * A sensor which logs the server (i.e. the PI) uptime and the program uptime.
 * 
 * Including this as I saw this on the Rochester Lodge weather station.
 * 
 * http://www.jpsc.co.uk/weather/
 * 
 * 
 */

#include <string.h>
#include <stdlib.h>
#include <sys/sysinfo.h>
#include <time.h>
#include "lib/config.h"
#include "sensors/sensors.h"

struct uptime_sensor {
    struct sensor sensor;
    time_t start;
};

static void logUptime(struct sensor *sensor, int uptime) {
    int m = (uptime / 60) % 60;
    int h = (uptime / 3600) % 24;
    int d = (uptime / 86400);

    if (d > 0)
        sensor_log(sensor, uptime, "%dd %dh %dm", d, h, m);
    else if (h > 0)
        sensor_log(sensor, uptime, "%dh %dm", h, m);
    else
        sensor_log(sensor, uptime, "%dm", m);
}

/**
 * Sensor update for system uptime
 * @param sensor
 */
static void uptime(struct sensor *sensor) {
    struct sysinfo info;
    if (!sysinfo(&info)) {
        // Seconds since boot. It's a long but we'd not have a machine up that long
        logUptime(sensor, (int) info.uptime);
    }
}

// Init prog uptime so we know when we started

static void prog_init(struct sensor *sensor) {
    struct uptime_sensor *s = (struct uptime_sensor *) sensor;
    time(&s->start);
}

static void prog_uptime(struct sensor *sensor) {
    struct uptime_sensor *s = (struct uptime_sensor *) sensor;
    time_t tm;
    time(&tm);

    tm = tm - s->start;
    logUptime(sensor, (int) tm);
}

void register_uptime_sensor(CONFIG_SECTION *sect) {
    // System uptime logger
    struct sensor *s = (struct sensor *) malloc(sizeof (struct sensor));
    memset(s, 0, sizeof (struct sensor));
    s->name = "uptime/server";
    s->title = "Server uptime";
    s->update = uptime;
    sensor_configure(sect, s);
    sensor_register(s);

    // Program uptime logger
    struct uptime_sensor *us = (struct uptime_sensor *) malloc(sizeof (struct uptime_sensor));
    memset(us, 0, sizeof (struct uptime_sensor));
    us->sensor.name = "uptime/program";
    us->sensor.title = "Program uptime";
    us->sensor.init = prog_init;
    us->sensor.update = prog_uptime;
    sensor_configure(sect, &us->sensor);
    sensor_register(&us->sensor);
}