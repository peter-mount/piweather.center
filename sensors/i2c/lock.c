#include "i2c.h"

int i2c_lock() {
    return pthread_mutex_lock(&i2c_slaves.mutex);
}

int i2c_unlock() {
    return pthread_mutex_unlock(&i2c_slaves.mutex);
}
