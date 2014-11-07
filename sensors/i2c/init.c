
#include "i2c.h"

struct i2c i2c_slaves;

void i2c_init() {
    list_init(&i2c_slaves.slaves);
    pthread_mutex_init(&i2c_slaves.mutex, NULL);
}
