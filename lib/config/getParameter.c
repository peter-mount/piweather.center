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
#include <strings.h>
#include <string.h>
#include "lib/config.h"

/**
 * Retrieves the named section
 * 
 * @param config
 * @param name
 * @return null if not found
 */
CONFIG_PARAM *config_getParameter(CONFIG_SECTION *sect, const char *name) {
    return (CONFIG_PARAM *) list_findNode(&sect->parameters, name);
}
