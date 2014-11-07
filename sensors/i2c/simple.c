
#include <pthread.h>

#include "i2c.h"

int i2c_sendSlaveCommand2(struct i2c_slave *slave, char *cmd, int cmdSize, char *response, int respSize) {
    int ret = i2c_write(slave, cmd, cmdSize);

    if (!ret && response) {
        i2c_sleep(slave->rw_delay);
        ret = i2c_read(slave, response, respSize);
    }

    if (!ret)
        i2c_sleep(slave->post_delay);

    return ret;
}

int i2c_sendCommand2(int address, char *cmd, int cmdSize, char *response, int respSize) {
    struct i2c_slave *slave = i2c_getSlave(address);
    return i2c_sendSlaveCommand2(slave, cmd, cmdSize, response, respSize);
}

int i2c_sendSlaveCommand(struct i2c_slave *slave, int cmd, char *response, int size) {
    char buf[1];
    buf[0] = cmd;

    return i2c_sendSlaveCommand2(slave, buf, 1, response, size);
}

int i2c_sendCommand(int address, int cmd, char *response, int size) {
    char buf[1];
    buf[0] = cmd;

    return i2c_sendCommand2(address, buf, 1, response, size);
}
