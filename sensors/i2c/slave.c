
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

// The I2C bus: This is for V2 pi's. For V1 Model B you need i2c-0
static const char *devName = "/dev/i2c-1";

struct i2c_slave *i2c_getSlave(int address) {

    struct Node *n = i2c_slaves.slaves.l_head;
    while (list_isNode(n) && ((struct i2c_slave *) n)->address != address)
        n = n->n_succ;

    struct i2c_slave *s;
    if (list_isNode(n))
        s = (struct i2c_slave *) n;
    else {
        s = (struct i2c_slave *) malloc(sizeof (struct i2c_slave));
        memset(s, 0, sizeof (struct i2c_slave));

        s->address = address;
        
        s->file = open(devName, O_RDWR);
        if (s->file < 0) {
            fprintf(stderr, "Failed to access I2C for address 0x%02x\n", address);
            exit(1);
        }

        if (ioctl(s->file, I2C_SLAVE, address) < 0) {
            fprintf(stderr, "Failed to acquire bus access/talk to I2C address 0x%02x\n", address);
            exit(1);
        }

        list_addTail(&i2c_slaves.slaves, &s->node);
    }

    return s;
}

