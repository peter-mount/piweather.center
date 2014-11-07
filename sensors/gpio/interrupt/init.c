
#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "sensors/gpio/gpio.h"

void gpio_configure_interrupt(CONFIG_SECTION *sect, struct interrupt_sensor *s, void (*handler)(struct interrupt_sensor *sensor)) {
    int pin = -1;

    config_getIntParameter(sect, "pin", &pin);
    if (pin < 0 || pin > GPIO_SYSFD_MAX) {
        fprintf(stderr, "Invalid GPIO Pin %d in %s\n", pin, sect->node.name);
        exit(1);
    }

    s->pin = pin;
    s->handler = hander;
}