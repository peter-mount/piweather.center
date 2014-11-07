/**
 * Adafruit SI1145 Digital UV Index / IR / Visible Light Sensor
 * 
 * This sensor consists of the core which does nothing but control the sensor.
 * It then supplies several sub-sensor's which contain's the measurements.
 * 
 * ===========================================================================
 * 
 * The core of this is based on Adafruit's Arduino library:
 * https://github.com/adafruit/Adafruit_SI1145_Library/
 * 
 * so this sensor is under the same BSD license:
 * 
 *  This is a library for the Si1145 UV/IR/Visible Light Sensor
 * 
 *  Designed specifically to work with the Si1145 sensor in the adafruit shop
 *   ----> https://www.adafruit.com/products/1777
 * 
 *   These sensors use I2C to communicate, 2 pins are required to interface
 *   Adafruit invests time and resources providing this open source code, 
 *   please support Adafruit and open-source hardware by purchasing 
 *   products from Adafruit!
 * 
 *   Written by Limor Fried/Ladyada for Adafruit Industries.  
 *   BSD license, all text above must be included in any redistribution
 */

#include <stdlib.h>
#include <lib/config.h>
#include <string.h>
#include <unistd.h>
#include "sensors/sensors.h"
#include "sensors/i2c/i2c.h"
#include "logger/logger.h"

/* COMMANDS */
#define SI1145_PARAM_QUERY 0x80
#define SI1145_PARAM_SET 0xA0
#define SI1145_NOP 0x0
#define SI1145_RESET    0x01
#define SI1145_BUSADDR    0x02
#define SI1145_PS_FORCE    0x05
#define SI1145_ALS_FORCE    0x06
#define SI1145_PSALS_FORCE    0x07
#define SI1145_PS_PAUSE    0x09
#define SI1145_ALS_PAUSE    0x0A
#define SI1145_PSALS_PAUSE    0xB
#define SI1145_PS_AUTO    0x0D
#define SI1145_ALS_AUTO   0x0E
#define SI1145_PSALS_AUTO 0x0F
#define SI1145_GET_CAL    0x12

/* Parameters */
#define SI1145_PARAM_I2CADDR 0x00
#define SI1145_PARAM_CHLIST   0x01
#define SI1145_PARAM_CHLIST_ENUV 0x80
#define SI1145_PARAM_CHLIST_ENAUX 0x40
#define SI1145_PARAM_CHLIST_ENALSIR 0x20
#define SI1145_PARAM_CHLIST_ENALSVIS 0x10
#define SI1145_PARAM_CHLIST_ENPS1 0x01
#define SI1145_PARAM_CHLIST_ENPS2 0x02
#define SI1145_PARAM_CHLIST_ENPS3 0x04

#define SI1145_PARAM_PSLED12SEL   0x02
#define SI1145_PARAM_PSLED12SEL_PS2NONE 0x00
#define SI1145_PARAM_PSLED12SEL_PS2LED1 0x10
#define SI1145_PARAM_PSLED12SEL_PS2LED2 0x20
#define SI1145_PARAM_PSLED12SEL_PS2LED3 0x40
#define SI1145_PARAM_PSLED12SEL_PS1NONE 0x00
#define SI1145_PARAM_PSLED12SEL_PS1LED1 0x01
#define SI1145_PARAM_PSLED12SEL_PS1LED2 0x02
#define SI1145_PARAM_PSLED12SEL_PS1LED3 0x04

#define SI1145_PARAM_PSLED3SEL   0x03
#define SI1145_PARAM_PSENCODE   0x05
#define SI1145_PARAM_ALSENCODE  0x06

#define SI1145_PARAM_PS1ADCMUX   0x07
#define SI1145_PARAM_PS2ADCMUX   0x08
#define SI1145_PARAM_PS3ADCMUX   0x09
#define SI1145_PARAM_PSADCOUNTER   0x0A
#define SI1145_PARAM_PSADCGAIN 0x0B
#define SI1145_PARAM_PSADCMISC 0x0C
#define SI1145_PARAM_PSADCMISC_RANGE 0x20
#define SI1145_PARAM_PSADCMISC_PSMODE 0x04

#define SI1145_PARAM_ALSIRADCMUX   0x0E
#define SI1145_PARAM_AUXADCMUX   0x0F

#define SI1145_PARAM_ALSVISADCOUNTER   0x10
#define SI1145_PARAM_ALSVISADCGAIN 0x11
#define SI1145_PARAM_ALSVISADCMISC 0x12
#define SI1145_PARAM_ALSVISADCMISC_VISRANGE 0x20

#define SI1145_PARAM_ALSIRADCOUNTER   0x1D
#define SI1145_PARAM_ALSIRADCGAIN 0x1E
#define SI1145_PARAM_ALSIRADCMISC 0x1F
#define SI1145_PARAM_ALSIRADCMISC_RANGE 0x20

#define SI1145_PARAM_ADCCOUNTER_511CLK 0x70

#define SI1145_PARAM_ADCMUX_SMALLIR  0x00
#define SI1145_PARAM_ADCMUX_LARGEIR  0x03

/* REGISTERS */
#define SI1145_REG_PARTID  0x00
#define SI1145_REG_REVID  0x01
#define SI1145_REG_SEQID  0x02

#define SI1145_REG_INTCFG  0x03
#define SI1145_REG_INTCFG_INTOE 0x01
#define SI1145_REG_INTCFG_INTMODE 0x02

#define SI1145_REG_IRQEN  0x04
#define SI1145_REG_IRQEN_ALSEVERYSAMPLE 0x01
#define SI1145_REG_IRQEN_PS1EVERYSAMPLE 0x04
#define SI1145_REG_IRQEN_PS2EVERYSAMPLE 0x08
#define SI1145_REG_IRQEN_PS3EVERYSAMPLE 0x10


#define SI1145_REG_IRQMODE1 0x05
#define SI1145_REG_IRQMODE2 0x06

#define SI1145_REG_HWKEY  0x07
#define SI1145_REG_MEASRATE0 0x08
#define SI1145_REG_MEASRATE1  0x09
#define SI1145_REG_PSRATE  0x0A
#define SI1145_REG_PSLED21  0x0F
#define SI1145_REG_PSLED3  0x10
#define SI1145_REG_UCOEFF0  0x13
#define SI1145_REG_UCOEFF1  0x14
#define SI1145_REG_UCOEFF2  0x15
#define SI1145_REG_UCOEFF3  0x16
#define SI1145_REG_PARAMWR  0x17
#define SI1145_REG_COMMAND  0x18
#define SI1145_REG_RESPONSE  0x20
#define SI1145_REG_IRQSTAT  0x21
#define SI1145_REG_IRQSTAT_ALS  0x01

#define SI1145_REG_ALSVISDATA0 0x22
#define SI1145_REG_ALSVISDATA1 0x23
#define SI1145_REG_ALSIRDATA0 0x24
#define SI1145_REG_ALSIRDATA1 0x25
#define SI1145_REG_PS1DATA0 0x26
#define SI1145_REG_PS1DATA1 0x27
#define SI1145_REG_PS2DATA0 0x28
#define SI1145_REG_PS2DATA1 0x29
#define SI1145_REG_PS3DATA0 0x2A
#define SI1145_REG_PS3DATA1 0x2B
#define SI1145_REG_UVINDEX0 0x2C
#define SI1145_REG_UVINDEX1 0x2D
#define SI1145_REG_PARAMRD 0x2E
#define SI1145_REG_CHIPSTAT 0x30

#define SI1145_ADDR 0x60

struct state {
    struct sensor sensor;
    // Format used in logging the sensor
    const char *format;
    // ============================
    // Internal use from this point
    // UV index
    struct sensor uv;
    // Visible+IR light levels
    struct sensor vis;
    // IR light level
    struct sensor ir;
    // poximity - assumes an IR LED is attached to LED pin on sensor board
    struct sensor proximity;
};

// ======================================================
// i2c api equivalents of methods in the Adafruit library

static uint8_t read8(uint8_t reg) {
    unsigned char buf[1];
    buf[0] = reg;
    i2c_sendCommand2(SI1145_ADDR, buf, 1, buf, 1);
    return buf[0];
}

static int read16(uint8_t a) {
    unsigned char buf[2];

    buf[0] = a;
    i2c_sendCommand2(SI1145_ADDR, buf, 1, buf, 2);
    return (int) buf[0] + ((int) buf[1] << 8);
}

static void write8(uint8_t reg, uint8_t val) {
    unsigned char buf[2];
    buf[0] = reg;
    buf[1] = val;
    i2c_sendCommand2(SI1145_ADDR, (char *) &buf, 2, NULL, 0);
}

static uint8_t readParam(uint8_t p) {
    write8(SI1145_REG_COMMAND, p | SI1145_PARAM_QUERY);
    return read8(SI1145_REG_PARAMRD);
}

static uint8_t writeParam(uint8_t p, uint8_t v) {
    write8(SI1145_REG_PARAMWR, v);
    write8(SI1145_REG_COMMAND, p | SI1145_PARAM_SET);
    return read8(SI1145_REG_PARAMRD);
}

// ==================================================

static void reset() {
    // Perform a software reset of the sensor's firmware
    write8(SI1145_REG_COMMAND, SI1145_RESET);
    i2c_sleep(2000L);

    // enable UVindex measurement coefficients!
    write8(SI1145_REG_UCOEFF0, 0x29);
    write8(SI1145_REG_UCOEFF1, 0x89);
    write8(SI1145_REG_UCOEFF2, 0x02);
    write8(SI1145_REG_UCOEFF3, 0x00);

    // enable UV sensor
    writeParam(SI1145_PARAM_CHLIST, SI1145_PARAM_CHLIST_ENUV |
            SI1145_PARAM_CHLIST_ENALSIR | SI1145_PARAM_CHLIST_ENALSVIS |
            SI1145_PARAM_CHLIST_ENPS1);
    // enable interrupt on every sample
    write8(SI1145_REG_INTCFG, SI1145_REG_INTCFG_INTOE);
    write8(SI1145_REG_IRQEN, SI1145_REG_IRQEN_ALSEVERYSAMPLE);

    // program LED current
    write8(SI1145_REG_PSLED21, 0x03); // 20mA for LED 1 only
    writeParam(SI1145_PARAM_PS1ADCMUX, SI1145_PARAM_ADCMUX_LARGEIR);
    // prox sensor #1 uses LED #1
    writeParam(SI1145_PARAM_PSLED12SEL, SI1145_PARAM_PSLED12SEL_PS1LED1);
    // fastest clocks, clock div 1
    writeParam(SI1145_PARAM_PSADCGAIN, 0);
    // take 511 clocks to measure
    writeParam(SI1145_PARAM_PSADCOUNTER, SI1145_PARAM_ADCCOUNTER_511CLK);
    // in prox mode, high range
    writeParam(SI1145_PARAM_PSADCMISC, SI1145_PARAM_PSADCMISC_RANGE |
            SI1145_PARAM_PSADCMISC_PSMODE);

    writeParam(SI1145_PARAM_ALSIRADCMUX, SI1145_PARAM_ADCMUX_SMALLIR);
    // fastest clocks, clock div 1
    writeParam(SI1145_PARAM_ALSIRADCGAIN, 0);
    // take 511 clocks to measure
    writeParam(SI1145_PARAM_ALSIRADCOUNTER, SI1145_PARAM_ADCCOUNTER_511CLK);
    // in high range mode
    writeParam(SI1145_PARAM_ALSIRADCMISC, SI1145_PARAM_ALSIRADCMISC_RANGE);
    // fastest clocks, clock div 1
    writeParam(SI1145_PARAM_ALSVISADCGAIN, 0);
    // take 511 clocks to measure
    writeParam(SI1145_PARAM_ALSVISADCOUNTER, SI1145_PARAM_ADCCOUNTER_511CLK);
    // in high range mode (not normal signal)
    writeParam(SI1145_PARAM_ALSVISADCMISC, SI1145_PARAM_ALSVISADCMISC_VISRANGE);

    // measurement rate for auto
    write8(SI1145_REG_MEASRATE0, 0xFF); // 255 * 31.25uS = 8ms

    // auto run
    write8(SI1145_REG_COMMAND, SI1145_PSALS_AUTO);
    
    i2c_sleep(100000L);
}

static void init(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;
        reset();
        i2c_unlock();
    }
}

static void update(struct sensor *sensor) {
    if (!i2c_lock()) {
        struct state *state = (struct state *) sensor;

        reset();
        
        // UV
        int uv = read16(SI1145_REG_UVINDEX0);
        state->uv.last_update = sensor->last_update;
        sensor_log(&state->uv, uv, "UV %.2f", (double) uv / 100.0);

        int vis = read16(SI1145_REG_ALSVISDATA0);
        state->vis.last_update = sensor->last_update;
        sensor_log(&state->vis, vis, "V/IR %d", vis);

        int ir = read16(SI1145_REG_ALSIRDATA0);
        state->ir.last_update = sensor->last_update;
        sensor_log(&state->ir, ir, "IR %d", ir);

        if (state->proximity.enabled) {
            int prox = read16(SI1145_REG_PS1DATA0);
            state->proximity.last_update = sensor->last_update;
            sensor_log(&state->proximity, prox, "Proximity %d", prox);
        }

        i2c_unlock();
    }
}

void register_si1145_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = sect->node.name;
    state->sensor.init = init;
    state->sensor.update = update;

    // The title is used in debug
    config_getCharParameter(sect, "title", (char **) &state->sensor.title);
    if (!state->sensor.title) {
        char tmp[256];
        snprintf(tmp, sizeof (tmp), "SI1145 I2C Digital UV Index / IR / Visible Light Sensor");
        state->sensor.title = strdup(tmp);
    }

    config_getCharParameter(sect, "format", (char **) &state->format);
    if (!state->format)
        state->format = "";

    /*
        config_getDoubleParameter(sect, "unit", &state->unit);
     */

    // i2c configuration
    struct i2c_slave *slave = i2c_getSlave(SI1145_ADDR);
    config_getLongParameter(sect, "i2c-rw-delay", &slave->rw_delay);
    config_getLongParameter(sect, "i2c-post-delay", &slave->post_delay);

    // Standard config
    sensor_configure(sect, &state->sensor);
    // Disallow parent from being annotated as it's a wrapper
    state->sensor.annotate = 0;
    sensor_register(&state->sensor);

    // Attach subsensors
    char tmp[256];

    // UV
    state->uv.enabled = state->sensor.enabled;
    snprintf(tmp, sizeof (tmp), "%s/uv", state->sensor.name);
    state->uv.name = strdup(tmp);
    config_getBooleanParameter(sect, "uv.annotate", &state->uv.annotate);
    sensor_register(&state->uv);

    // Visible
    state->vis.enabled = state->sensor.enabled;
    snprintf(tmp, sizeof (tmp), "%s/visible", state->sensor.name);
    state->vis.name = strdup(tmp);
    config_getBooleanParameter(sect, "visible.annotate", &state->vis.annotate);
    sensor_register(&state->vis);

    // IR
    state->ir.enabled = state->sensor.enabled;
    snprintf(tmp, sizeof (tmp), "%s/ir", state->sensor.name);
    state->ir.name = strdup(tmp);
    config_getBooleanParameter(sect, "ir.annotate", &state->ir.annotate);
    sensor_register(&state->ir);

    // Proximity, as it's optional hardware it has it's own enable
    snprintf(tmp, sizeof (tmp), "%s/proximity", state->sensor.name);
    state->proximity.name = strdup(tmp);
    config_getBooleanParameter(sect, "proximity.enable", &state->proximity.enabled);
    config_getBooleanParameter(sect, "proximity.annotate", &state->proximity.annotate);
    sensor_register(&state->proximity);
}