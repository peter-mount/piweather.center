#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "astro/observatory.h"

OBSERVATORY observatory;

void astro_init() {
    CONFIG_SECTION *sect = config_getSection("observatory");
    if (!sect) {
        fprintf(stderr, "Unable to locate station configuration section\n");
        exit(1);
    }

    // Initialise observatory with some defaults
    observatory.name = "Unknown";
    observatory.altitude = observatory.latitude = observatory.longitude = 0.0;

    // Read in the parameters
    config_getCharParameter(sect, "name", &observatory.name);
    config_getDoubleParameter(sect, "latitude", &observatory.latitude);
    config_getDoubleParameter(sect, "longitude", &observatory.longitude);
    config_getDoubleParameter(sect, "altitude", &observatory.altitude);
}