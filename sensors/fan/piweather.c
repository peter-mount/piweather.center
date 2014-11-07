/*
 * Fan control handler for the pi weather board high power outputs
 */

#include <stdlib.h>
#include <pthread.h>
#include <string.h>
#include <stdio.h>
#include "lib/config.h"
#include "lib/string.h"
#include "sensors/i2c/i2c.h"

extern int verbose;

struct state {
    int port;
    struct i2c_slave *slave;
};

void *fan_piweather_init(CONFIG_SECTION *sect) {
    int port = 0;
    config_getIntParameter(sect, "fan-port", &port);
    if (port < 1 || port > 4)
        fatalError("fan-port %d is invalid, must be 1..4 in %s", port, sect->node.name);

    struct state *s = (struct state *) malloc(sizeof (struct state));
    memset(s, 0, sizeof (struct state));
    s->port = port;
    s->slave = i2c_getSlave(0x4e);
    return (void *) s;
}

void fan_piweather_control(void *arg, int state) {
    char cmd[2];
    struct state *s = (struct state *) arg;

    cmd[0] = state ? 0x25 : 0x26;
    cmd[1] = (char) s->port;

    i2c_sendSlaveCommand2(s->slave, cmd, 2, NULL, 0);
}