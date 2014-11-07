/**
 * Manages the configuration files
 * 
 * The config is in a directory, usually /etc/weatherstation and consists of one file per unit.
 * 
 * For example, the camera is defined as /etc/weatherstation/camera
 * 
 * The format is specific for each module but is usually a set of key value pairs delimited with whitespace
 */

#include <stdlib.h>
#include "lib/config.h"
#include "lib/list.h"

/**
 * Retrieves the named section
 * 
 * @param config
 * @param name
 * @return null if not found
 */
CONFIG_SECTION *config_getSection(const char *name) {
    // If no config loaded then return NULL
    if (config)
        return (CONFIG_SECTION *) list_findNode(&config->sections, name);
    else
        return NULL;
}
