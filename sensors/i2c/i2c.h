/* 
 * File:   i2c.h
 * Author: peter
 *
 * Created on February 27, 2014, 7:44 PM
 */

#ifndef I2C_H
#define	I2C_H

#include "global_config.h"
#ifdef HAVE_I2C

#include "lib/list.h"
#include <pthread.h>

struct i2c {
    struct List slaves;
    pthread_mutex_t mutex;
};
extern struct i2c i2c_slaves;

struct i2c_slave {
    struct Node node;
    // The i2c host address
    int address;
    // Delay in nano-seconds between the write and read, used for slow devices
    long rw_delay;
    // Delay in nano-seconds after a used for slow devices
    long post_delay;
    // =================================
    // Internal use only from this point
    // =================================
    // File socket to the slave
    int file;
};

extern int i2c_lock();
extern int i2c_unlock();

/**
 * Send a single byte command to a slave
 * @param slave i2c_slave
 * @param cmd Command to send
 * @param response where to write the response
 * @param size The expected size of the response
 * @return 0 success, 1 faulre
 */
extern int i2c_sendSlaveCommand(struct i2c_slave *slave, int cmd, char *response, int size);
/**
 * Send a single byte command to a slave
 * @param i2c core i2c structure
 * @param address i2c device address
 * @param cmd Command to send
 * @param response where to write the response
 * @param size The expected size of the response
 * @return 0 success, 1 faulre
 */
extern int i2c_sendCommand(int address, int cmd, char *response, int size);
/**
 * Send a complex command to a slave
 * @param slave
 * @param cmd buffer to send
 * @param cmdSize size of command
 * @param response buffer to recive, NULL for no data
 * @param respSize size of expected response
 * @return 
 */
extern int i2c_sendSlaveCommand2(struct i2c_slave *slave, char *cmd, int cmdSize, char *response, int respSize);
extern int i2c_sendCommand2(int address, char *cmd, int cmdSize, char *response, int respSize);
/**
 * Initialise the i2c framework
 * @param i2c
 */
extern void i2c_init();
/**
 * Lookup and create as required an i2c_slave
 * @param i2c core i2c structure
 * @param address i2c device address
 * @return i2c_slave
 */
extern struct i2c_slave *i2c_getSlave(int address);
/**
 * Read bytes from a slave
 * @param slave Slave
 * @param b buffer to write to
 * @param s bytes expected
 * @return 1 on error, 0 on success
 */
extern int i2c_read(struct i2c_slave *slave, char *b, int s);
/**
 * Write data to a slave
 * @param slave Slave
 * @param b Buffer to send
 * @param s bytes to send
 * @return 0 on success, 1 on error
 */
extern int i2c_write(struct i2c_slave *slave, char *b, int s);
extern void i2c_sleep(long t);

#endif

#endif	/* I2C_H */

