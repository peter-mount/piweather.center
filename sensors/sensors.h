/* 
 * File:   temp.h
 * Author: peter
 *
 * Created on February 12, 2014, 3:12 PM
 */

#ifndef SENSORS_H
#define	SENSORS_H

#include <time.h>
#include <pthread.h>
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "weatherstation/main.h"

#define SENSOR_TEXT_SIZE 64

/**
 * Defined in main.c this holds the available sensor registries
 */
struct sensor_registry {
    const char *type;
    void (*registry)(CONFIG_SECTION *sect);
    const char *desc;
};

/**
 * The definition of a sensor
 */
struct sensor {
    // The name of this sensor
    const char *name;
    // The descriptive title for this sensor
    const char *title;
    // !0 to enable this sensor
    int enabled;
    // !0 to show on the annotated image
    int annotate;
    // Optional, Initialise this sensor
    void (*init)(struct sensor *sensor);
    void (*postinit)(struct sensor *sensor);
    // Update this sensor
    void (*update)(struct sensor *sensor);
    // Optional, show debug
    void (*debug)(struct sensor *sensor);
    // =========================
    // Internal to sensor.c only
    // =========================
    struct sensor *next;
    // Optional, if present a chain of listeners which will be notified of an update
    struct sensor_listener *listeners;
    // When we last performed an update
    time_t last_update;
    // Last value, used to check for changes
    int value;
    // Text value used by renderers
    char text[SENSOR_TEXT_SIZE];
};

/*
 * Sensor listener
 */
struct sensor_listener {
    struct sensor_listener *next;
    struct sensor *listener;
    void (*update)(struct sensor *sensor, struct sensor *listener);
};

/*
 * Our working collection of sensors
 */
struct sensors {
    // Our registered sensors
    struct sensor *sensors;
};
extern struct sensors *sensors;

extern void sensor_init();
extern void sensor_postinit();
extern void sensor_log(struct sensor *sensor, int value, char *fmt, ...);
extern void sensor_register(struct sensor *sensor);
extern void sensor_registerAll(struct sensor_registry *registries);
extern void sensor_configure(CONFIG_SECTION *sect, struct sensor *sensor);
extern void sensor_start();
extern void sensor_update();
extern void *sensor_loop();

extern struct sensor *sensors_get(const char *name);

extern void sensor_add_listener(
        struct sensor *sensor,
        struct sensor *listener,
        void (*update)(struct sensor *sensor, struct sensor *listener)
        );
#endif	/* SENSORS_H */

