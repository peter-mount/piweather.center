/**
 * This is the ADCPI Version 1 board
 */

#include <stdlib.h>
#include <string.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "sensors/i2c/i2c.h"

// Addresses of both ADC chips on this board
#define ADDRESS1 0x68
#define ADDRESS2 0x69

// Commands for the 4 channels on each chip
#define CHANNEL_1 0x98
#define CHANNEL_2 0xb8
#define CHANNEL_3 0xd8
#define CHANNEL_4 0xf8

// Sleep for 1ms (0.01s in original python source)
#define SLEEP 1000

static int channels[] = {
    CHANNEL_1, CHANNEL_2, CHANNEL_3, CHANNEL_4,
    CHANNEL_1, CHANNEL_2, CHANNEL_3, CHANNEL_4
};

struct state {
    struct sensor sensor;
    const char *format;
    double unit;
    struct i2c_slave *slave;
    char channel;
};

static void update(struct sensor *sensor) {
    if(!i2c_lock()) {
        struct state *state = (struct state *)sensor;
        uint8_t cmd[4];
        int r;
        
        cmd[0] = state->channel;
        r = i2c_write(state->slave,cmd,1);
        
        if(!r) {
            i2c_sleep( SLEEP );
            r = i2c_read(state->slave,cmd,3);
        }
        
        if(!r) {
            // The raw reading
            int t = (cmd[0]<<8) | cmd[1];
            if(t>=32768)
                t=65536-t;
            
            // The voltage
            double v = t * 0.000154;
            if( v>=5.5)
                v=0.0;
            
            sensor_log( sensor, t, "%s %.1fV", state->format, v);
        }
        
        i2c_unlock();
    }
}

void register_adcpi1_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = sect->node.name;
    state->sensor.update = update;
    
    int port = -1;
    config_getIntParameter(sect,"port",&port);
    if(port<0 || port>7) {
        fprintf(stderr, "%s: ADC PI Port must be 0..7, got %d for %s\n", port, sect->node.name);
        exit(1);
    }
    
    int address = port>4 ? ADDRESS2:ADDRESS1;
    state->channel = channels[port];

    // The title is used in debug
    config_getCharParameter(sect, "title", (char **) &state->sensor.title);
    if (!state->sensor.title) {
        char tmp[256];
        snprintf(tmp, sizeof (tmp), "BH1750 I2C Ambient Light, address 0x%02x", address);
        state->sensor.title = strdup(tmp);
    }

    config_getCharParameter(sect, "format", (char **) &state->format);
    if (!state->format)
        state->format = "ADC";

    // i2c configuration
    state->slave = i2c_getSlave(address);
    config_getLongParameter(sect, "i2c-rw-delay", &state->slave->rw_delay);
    config_getLongParameter(sect, "i2c-post-delay", &state->slave->post_delay);

    // Finally standard config & register it in the system
    sensor_configure(sect, &state->sensor);
    sensor_register((struct sensor *) state);
}
