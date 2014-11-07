/**
 * Manages the configuration files
 * 
 * The config is in a directory, usually /etc/weatherstation and consists of one file per unit.
 * 
 * For example, the camera is defined as /etc/weatherstation/camera
 * 
 * The format is specific for each module but is usually a set of key value pairs delimited with whitespace
 */

#include "lib/config.h"

int config_getCharParameter(CONFIG_SECTION *sect, const char *name, char **val) {
    CONFIG_PARAM *p = config_getParameter(sect, name);
    if (p) {
        *val = (char *) p->value;
        return 0;
    }
    return 1;
}
