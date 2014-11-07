
#include <pthread.h>
#include <stdint.h>
#include <string.h>
#include <unistd.h>  
#include <errno.h>  
#include <stdio.h>  
#include <stdlib.h>  
#include <linux/i2c-dev.h> 
#include <sys/ioctl.h>  
#include <fcntl.h> 
#include <unistd.h>
#include "sensors/i2c/i2c.h"

/**
 * Read bytes from a slave
 * @param slave Slave
 * @param b buffer to write to
 * @param s bytes expected
 * @return 1 on error, 0 on success
 */
int i2c_read(struct i2c_slave *slave, char *b, int s) {
    return read(slave->file, b, s) != s;
}
