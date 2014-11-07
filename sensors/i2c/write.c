
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
 * Write data to a slave
 * @param slave Slave
 * @param b Buffer to send
 * @param s bytes to send
 * @return 0 on success, 1 on error
 */
int i2c_write(struct i2c_slave *slave, char *b, int s) {
    return write(slave->file, b, s) != s;
}
