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

int config_getHexParameter(CONFIG_SECTION *sect, const char *name, uint32_t *val) {
    return config_scanParameter(sect, name, "%x", (void *) val);
}
