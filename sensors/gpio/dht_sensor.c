/**
 * dht_sensor.c
 * 
 * Provides support for the DHT11/DHT22 temperature/humidity sensor directly connected to a GPIO pin
 * 
 */

#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <wiringPi.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "weatherstation/main.h"

#define MAXTIMINGS 85

struct state {
    // This must be the first entry, it's what the state engine will see this struct as
    struct sensor sensor;
    // -----------------
    // INTERNAL USE ONLY
    // -----------------
    // Sub-sensors that are used to perform actual logging
    struct sensor temp_sensor;
    struct sensor humidity_sensor;
    // GPIO Pin sensor is attached to
    int pin;
    // Received data
    float temp, humidity;
    // Temporary buffer
    int data[5];
};

static uint8_t sizecvt(const int read) {
    /* digitalRead() and friends from wiringpi are defined as returning a value
    < 256. However, they are returned as int() types. This is a safety function */

    if (read > 255 || read < 0) {
        fprintf(stderr, "Invalid data from wiringPi library\n");
        exit(EXIT_FAILURE);
    }
    return (uint8_t) read;
}

static int read_dht22_dat(struct state *state) {
    uint8_t laststate = HIGH;
    uint8_t counter = 0;
    uint8_t j = 0, i;

    state->data[0] = state->data[1] = state->data[2] = state->data[3] = state->data[4] = 0;

    // pull pin down for 18 milliseconds
    pinMode(state->pin, OUTPUT);
    digitalWrite(state->pin, HIGH);
    delay(10);
    digitalWrite(state->pin, LOW);
    delay(18);
    // then pull it up for 40 microseconds
    digitalWrite(state->pin, HIGH);
    delayMicroseconds(40);
    // prepare to read the pin
    pinMode(state->pin, INPUT);

    // detect change and read data
    for (i = 0; i < MAXTIMINGS; i++) {
        counter = 0;
        while (sizecvt(digitalRead(state->pin)) == laststate) {
            counter++;
            delayMicroseconds(1);
            if (counter == 255) {
                break;
            }
        }
        laststate = sizecvt(digitalRead(state->pin));

        if (counter == 255) break;

        // ignore first 3 transitions
        if ((i >= 4) && (i % 2 == 0)) {
            // shove each bit into the storage bytes
            state->data[j / 8] <<= 1;
            if (counter > 16)
                state->data[j / 8] |= 1;
            j++;
        }
    }

    // check we read 40 bits (8bit x 5 ) + verify checksum in the last byte
    // print it out if data is good
    if ((j >= 40) && (state->data[4] == ((state->data[0] + state->data[1] + state->data[2] + state->data[3]) & 0xFF))) {

        state->humidity = ((float) state->data[0] * 256 + (float) state->data[1]) / 10.0;

        state->temp = ((float) (state->data[2] & 0x7F)* 256 + (float) state->data[3]) / 10.0;

        if ((state->data[2] & 0x80) != 0)
            state->temp *= -1;

        return 1;
    } else {
        return 0;
    }
}

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    //sensor_log(camera, sensor, value, "CPU %0.1fC", (double) value / 1000.0);
}

void register_dht22_sensor(CONFIG_SECTION *sect) {
    char temp[256];

    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The primary sensor which does the actual measurement
    state->sensor.name = sect->node.name;
    state->sensor.title = "DHT22 Sensor";
    state->sensor.update = update;

    // The sub-sensors which are actually updated by the primary
    snprintf(temp, sizeof (temp), "%s/temp", sect->node.name);
    state->temp_sensor.name = strdup(temp);
    state->temp_sensor.title = "DHT22 Temperature Sensor";

    snprintf(temp, sizeof (temp), "%s/humidity", sect->node.name);
    state->temp_sensor.name = strdup(temp);
    state->temp_sensor.title = "DHT22 Humidity Sensor";

    // Common config on all three sensors
    // FIXME is this needed for sub-sensors?
    sensor_configure(sect, &state->sensor);
    sensor_configure(sect, &state->temp_sensor);
    sensor_configure(sect, &state->humidity_sensor);

    // Now register the sensors, primary first
    sensor_register(&state->sensor);
    sensor_register(&state->temp_sensor);
    sensor_register(&state->humidity_sensor);
}
