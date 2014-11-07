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

void config_getBooleanParameter(CONFIG_SECTION *sect, const char *name, int *val) {
    CONFIG_PARAM *p = config_getParameter(sect, name);
    if (p) {
        const char *s = p->value;
        if (strcmp(s, "true") == 0 || strcmp(s, "enabled") == 0)
            *val = 1;
        else if (strcmp(s, "false") == 0 || strcmp(s, "disabled") == 0)
            *val = 0;
    }
};
