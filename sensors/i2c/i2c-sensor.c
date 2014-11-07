/*
 * A dynamic sensor for handling basic i2c based sensors
 */


#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "sensors/i2c/i2c.h"
#include "weatherstation/main.h"

struct state {
    struct sensor sensor;
    // The i2c command
    int cmd;
    // Format used in logging the sensor
    const char *format;
    // The unit. This is multiplied to the result before formatting
    double unit;
    // ============================
    // Internal use from this point
    struct i2c_slave *slave;
};

// ======================================================================
// Sensor reading implementations

/*
 * Formats the response and logs it
 */
static void logSensor(struct state *state, int value) {
    if (state->unit != 0.0) {
        // We have a unit so multiply value by it and pass the double to the format
        double dVal = ((double) value) * state->unit;
        sensor_log(&state->sensor, value, (char *) state->format, dVal);
    } else {
        // Pass the raw integer to the format
        sensor_log(&state->sensor, value, (char *) state->format, value);
    }
}

static void update_byte(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        unsigned char response[1];
        if (!i2c_sendSlaveCommand(state->slave, state->cmd, response, 1)) {
            int value = (int) response[0];
            logSensor(state, value);
        }
        i2c_unlock();
    }
}

static void update_signed_byte(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        char response[1];
        if (!i2c_sendSlaveCommand(state->slave, state->cmd, response, 1)) {
            int value = (int) response[0];
            logSensor(state, value);
        }
        i2c_unlock();
    }
}

static void update_word(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        unsigned char response[2];
        if (!i2c_sendSlaveCommand(state->slave, state->cmd, response, 2)) {
            int value = (int) response[0] + ((int) response[1] << 8);
            logSensor(state, value);
        }
        i2c_unlock();
    }
}

static void update_signed_word(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        char response[2];
        if (!i2c_sendSlaveCommand(state->slave, state->cmd, response, 2)) {
            int value = (int) response[0] + ((int) response[1] << 8);
            logSensor(state, value);
        }
        i2c_unlock();
    }
}

static void update_signed_word_bit31(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        char response[2];
        if (!i2c_sendSlaveCommand(state->slave, state->cmd, response, 2)) {
            int value = (int) response[0] + ((int) response[1] << 8);
            int temp = value & 0x7ff;
            if (value & 0x8000)
                temp = -temp;
            logSensor(state, temp);
        }
        i2c_unlock();
    }
}

// Lookup table used to select the correct sensor implementation by name

struct sensor_type {
    char *value;
    void (*update)(struct sensor *sensor);
};
static struct sensor_type sensor_types[] = {
    { "byte", update_byte},
    { "signed byte", update_signed_byte},
    { "word", update_word},
    { "signed word", update_signed_word},
    { "signed word bit 31", update_signed_word_bit31},
    // End of list marker
    { NULL, NULL}
};

// End of Sensor reading implementations
// ======================================================================

/**
 * Register an i2c sensor based on the supplied configuration section
 * @param camera
 * @param sect
 */
void register_i2c_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = sect->node.name;

    int address = -1;
    config_getHexParameter(sect, "i2c-address", &address);
    if (address == -1) {
        fprintf(stderr, "%s: i2c 0x%02x i2c-address is mandatory\n", sect->node.name);
        exit(1);
    } else if (address < 0 || address > 127) {
        fprintf(stderr, "%s: i2c 0x%02x address is invalid\n", sect->node.name, address);
        exit(1);
    }

    state->cmd = -1;
    config_getHexParameter(sect, "i2c-command", &state->cmd);
    if (state->cmd == -1) {
        fprintf(stderr, "%s: i2c 0x%02x i2c-command is mandatory\n", sect->node.name, address);
        exit(1);
    } else if (state->cmd > 255) {
        fprintf(stderr, "%s: i2c 0x%02x command %02x is invalid\n", sect->node.name, address, state->cmd);
        exit(1);
    }

    // The title is used in debug
    config_getCharParameter(sect, "title", (char **) &state->sensor.title);
    if (!state->sensor.title) {
        char tmp[256];
        snprintf(tmp, sizeof (tmp), "I2C sensor, address 0x%02x, cmd 0x%02x", address, state->cmd);
        state->sensor.title = strdup(tmp);
    }

    config_getCharParameter(sect, "format", (char **) &state->format);
    if (!state->format)
        state->format = "%d";

    config_getDoubleParameter(sect, "unit", &state->unit);

    // update is the only mandatory hook, here we pick the relevant method
    // based on our config
    char *type;
    config_getCharParameter(sect, "i2c-response-type", &type);
    if (type) {
        int i = 0;
        while (!state->sensor.update && sensor_types[i].value != NULL) {
            if (strcmp(sensor_types[i].value, type) == 0)
                state->sensor.update = sensor_types[i].update;
            i++;
        }
        if (!state->sensor.update) {
            fprintf(stderr, "%s: i2c-response-type %s is unsupported\n", sect->node.name, type);
            exit(1);
        }
    } else {
        fprintf(stderr, "%s: i2c-response-type is mandatory\n", sect->node.name);
        exit(1);
    }

    // i2c configuration
    state->slave = i2c_getSlave(address);
    config_getLongParameter(sect, "i2c-rw-delay", &state->slave->rw_delay);
    config_getLongParameter(sect, "i2c-post-delay", &state->slave->post_delay);

    // Finally standard config & register it in the system
    sensor_configure(sect, &state->sensor);
    sensor_register((struct sensor *) state);
}