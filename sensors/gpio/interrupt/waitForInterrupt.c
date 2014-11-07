
#include <stdlib.h>
#include <stdio.h>
#include <poll.h>
#include "sensors/gpio/gpio.h"

int gpio_waitForInterrupt(int pin, int timeout) {
    int fd = gpio.sysfd[pin];
    if (fd == -1)
        return -2;

    struct pollfd polls;
    polls.fd = fd;
    polls.events = POLLPRI; // Urgent data!

    int x = poll(&polls, 1, timeout);

    uint8_t c;
    (void) read(fd, &c, 1);

    return x;
}


