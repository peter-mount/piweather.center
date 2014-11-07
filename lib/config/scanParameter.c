/**
 * Manages the configuration files
 * 
 * The config is in a directory, usually /etc/weatherstation and consists of one file per unit.
 * 
 * For example, the camera is defined as /etc/weatherstation/camera
 * 
 * The format is specific for each module but is usually a set of key value pairs delimited with whitespace
 */

#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <string.h>
#include "lib/charbuffer.h"
#include "lib/config.h"

int config_scanParameter(CONFIG_SECTION *sect, const char *name, const char *fmt, void *val) {
    CONFIG_PARAM *p = config_getParameter(sect, name);

    return p == NULL || 0 == sscanf(p->value, fmt, val);
}
