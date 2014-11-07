/* 
 * File:   interruptHandler.c
 * Author: peter
 *
 * Created on 08 April 2014, 18:07
 */

#include <stdio.h>
#include <stdlib.h>
#include "lib/thread.h"
#include "sensors/gpio/gpio.h"

void *gpio_interruptHandler(void *arg) {
    struct interrupt_sensor *s = (struct interrupt_sensor *) arg;

    // Increase the priority, only works if we are root
    thread_setPriority(55);
    while (1)
        if (gpio_waitForInterrupt(s->pin, -1) > 0)
            s->handler(s);

    return NULL;
}
