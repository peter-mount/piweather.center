/* 
 * File:   interrupts.h
 * Author: peter
 *
 * Created on 08 April 2014, 18:08
 */

#ifndef COUNTER_H
#define	COUNTER_H

#include "sensors/sensors.h"

#define GPIO_SYSFD_MAX 64

struct gpio {
    // Map file descriptor for /sys/class/gpio/gpioX/value
    int sysfd[GPIO_SYSFD_MAX];
};

/*
 * Base for interrupt based sensors
 */
struct interrupt_sensor {
    struct sensor sensor;
    // The gpio pin we are monitoring state transitions
    int pin;
    // The handler method
    void (*handler)(struct interrupt_sensor *sensor);
};

/*
 * Base for counters like anemometers or rain gauges
 */
struct counter_sensor {
    struct interrupt_sensor sensor;
    // The current counter value
    unsigned long counter;
};

extern struct gpio gpio;

extern void gpio_init();
extern void gpio_configure_interrupt(CONFIG_SECTION *sect, struct interrupt_sensor *s, void (*handler)(struct interrupt_sensor *sensor));
extern void *gpio_interruptHandler(void *arg);
extern int gpio_waitForInterrupt(int pin, int timeout);

#endif	/* COUNTER_H */
