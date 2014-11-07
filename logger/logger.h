/* 
 * File:   logger.h
 * Author: Peter T Mount
 *
 * Created on March 28, 2014, 11:55 AM
 */

#ifndef LOGGER_H
#define	LOGGER_H

#include <stdlib.h>
#include <time.h>
#include "lib/config.h"
#include "lib/list.h"
#include "lib/threadpool.h"

#define LOG_TEXT_SIZE 32

/**
 * Structure passed to loggers.
 */
struct log_entry {
    // Node used by job queue, must be first
    // name field is the sensor name
    struct Node node;
    // The time of this entry
    time_t time;
    // Text from the sensor
    char text[LOG_TEXT_SIZE];
    // Sensor raw value
    int value;
    // Has sensor value changed or is it stable
    int updated;
};

/**
 * Definition of a logger
 */
struct logger {
    struct Node node;
    // Is this logger enabled
    int enabled;
    // =================
    // Internal use only
    // =================
    // All loggers are handled sequentially with a single thread
    struct thread_pool threadPool;
    // Optional, allows a logger to set itself up
    void (*init)(struct logger *);
    // Optional shutdown hook
    void (*stop)(struct logger *);
    // Hook that will handle an incoming log request
    void (*update) (struct logger *, struct log_entry *);
};

extern struct loggers loggers;

/**
 * Global struct for the logging framework
 */
struct loggers {
    struct List loggers;
    // Unique host id
    char *hostid;
};

extern void logger_free_entry(struct log_entry *e);
extern struct log_entry *logger_create_entry(time_t time, char *name, char *text, int value, int updated);
extern void logger_init();
extern void logger_log(time_t time, char *name, char *text, int value, int updated);
extern void logger_register(struct logger *logger);
extern void logger_start();

#endif	/* LOGGER_H */

