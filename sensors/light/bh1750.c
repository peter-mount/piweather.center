/**
 * BH1750FVI I2C Ambient Light Sensor
 */

#include <stdlib.h>
#include <string.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "sensors/i2c/i2c.h"

// BH1750 command set (not all used here)
#define CMD_POWER_OFF   0x00
#define CMD_POWER_ON    0x01
#define CMD_RESET       0x07
// continuous samples
#define CMD_CONT_HIRES  0x10
#define CMD_CONT_HIRES2 0x11
#define CMD_CONT_LORES  0x12
// Single sample
#define CMD_SAMP_HIRES  0x20
#define CMD_SAMP_HIRES2 0x21
#define CMD_SAMP_LORES  0x23

// Sample speeds in microseconds
#define SPEED_HIRES     180000
#define SPEED_HIRES2    180000
#define SPEED_LORES     24000

// Our logging limit
#define UNIT            0.1

struct state {
    struct sensor sensor;
    // Format used in logging the sensor
    const char *format;
    // The unit. This is multiplied to the result before formatting
    double unit;
    // ============================
    // Internal use from this point
    struct i2c_slave *slave;
    // 0 for low level, 1 for high level
    int level;
};

static void update(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        uint8_t cmd[4];
        int r = 0;
        int v;

        cmd[0] = state->level ? CMD_SAMP_HIRES2 : CMD_SAMP_LORES;
        r = i2c_write(state->slave, cmd, 1);

        if (!r) {
            i2c_sleep(state->level ? SPEED_HIRES2 : SPEED_LORES);
            r = i2c_read(state->slave, cmd, 2);
        }

        if (!r) {
            v = (cmd[0] << 8) | cmd[1];
            
            // Scale is 0.1 but we need to convert to the correct value
            // So datasheet says / 1.2 but we use 0.12 as thats 10/1.2 giving
            // us the correct unit
            v = (int) ((double) v / 0.12);
        }

        // log the value
        if (!r) {
            sensor_log(sensor, v, "%s %.1f", state->format, ((double) v) * UNIT);
        }

        i2c_unlock();
    }
}

void register_bh1750_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = sect->node.name;
    state->sensor.update = update;
    
    // Use HIRES mode
    state->level = 1;

    int address = 0x23;
    config_getHexParameter(sect, "i2c-address", &address);
    if (address == -1) {
        fprintf(stderr, "%s: i2c 0x%02x i2c-address is mandatory\n", sect->node.name);
        exit(1);
    } else if (address < 0 || address > 127) {
        fprintf(stderr, "%s: i2c 0x%02x address is invalid\n", sect->node.name, address);
        exit(1);
    }

    // The title is used in debug
    config_getCharParameter(sect, "title", (char **) &state->sensor.title);
    if (!state->sensor.title) {
        char tmp[256];
        snprintf(tmp, sizeof (tmp), "BH1750 I2C Ambient Light, address 0x%02x", address);
        state->sensor.title = strdup(tmp);
    }

    config_getCharParameter(sect, "format", (char **) &state->format);
    if (!state->format)
        state->format = "Lux";

    config_getDoubleParameter(sect, "unit", &state->unit);

    // i2c configuration
    state->slave = i2c_getSlave(address);
    config_getLongParameter(sect, "i2c-rw-delay", &state->slave->rw_delay);
    config_getLongParameter(sect, "i2c-post-delay", &state->slave->post_delay);

    // Finally standard config & register it in the system
    sensor_configure(sect, &state->sensor);
    sensor_register((struct sensor *) state);
}