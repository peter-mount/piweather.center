/*
 * 
 */
#include "sensors/gpio/gpio.h"

struct gpio gpio;

void gpio_init() {
    int i;

    for (i = 0; i < GPIO_SYSFD_MAX; i++)
        gpio.sysfd[i] = -1;

}