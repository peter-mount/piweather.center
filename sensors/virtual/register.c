
#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "sensors/sensors.h"

// The supported virtual sensors

extern void register_virtual_dewpoint(CONFIG_SECTION *sect);
extern void register_virtual_heatindex(CONFIG_SECTION *sect);
extern void register_virtual_pressure_trend(CONFIG_SECTION *sect);
extern void register_virtual_windchill(CONFIG_SECTION *sect);
extern void register_fan_sensor(CONFIG_SECTION *sect);
extern void register_virtual_cloudbase(CONFIG_SECTION *sect);

static struct sensor_registry virtual_sensors[] = {
    // Calculates values based on other sensors
    {"cloudbase", register_virtual_cloudbase, "Calculates cloud base altitude"},
    {"dewpoint", register_virtual_dewpoint, "Calculates dewpoint from temp & humidity"},
    {"heatindex", register_virtual_heatindex, "Calculates the heat index from temp & humidity"},
    {"pressure-trend", register_virtual_pressure_trend, "Calculates pressure trend from last x hours"},
    {"wind chill", register_virtual_windchill, "Calculates wind chill from temp & wind speed"},
    // Controls external devices based on sensor values
    {"fan", register_fan_sensor, "Controls a Fan/Relay based on temperature"},
    // List terminator
    {NULL, NULL, NULL}
};

/**
 * Registers all virtual sensors
 */
void register_virtual_sensor(CONFIG_SECTION *sect) {

    char *type = NULL;
    config_getCharParameter(sect, "virtual-type", &type);

    if (!type) {
        fprintf(stderr, "Sensor %s is declared virtual but has no virtual-type\n", sect->node.name);
        exit(1);
    }

    int i = 0;
    void (*registry)(CONFIG_SECTION * sect) = NULL;
    while (!registry && virtual_sensors[i].type) {
        if (strcmp(virtual_sensors[i].type, type) == 0) {
            registry = virtual_sensors[i].registry;
        }
        i++;
    }

    if (registry) {
        registry(sect);
    } else {
        fprintf(stderr, "Unsupported virtual-type %s in section %s\n", type, sect->node.name);
        exit(1);
    }
}