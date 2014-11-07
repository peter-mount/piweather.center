
#include "i2c.h"

void i2c_sleep(long t) {
    if (t > 0)
        usleep(t);
}