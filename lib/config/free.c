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
 * Frees up memory used by configuration
 * @param config
 */
void config_free() {

    struct Node *sn = list_getHead(&config->sections);
    while (list_isNode(sn)) {
        CONFIG_SECTION *sect = (CONFIG_SECTION *) sn;
        sn = list_remove(sn);

        struct Node *pn = list_getHead(&sect->parameters);
        while (list_isNode(pn)) {
            CONFIG_PARAM *param = (CONFIG_PARAM *) pn;
            pn = list_remove(pn);

            free((void*) param->value);
            node_free((struct Node *) param);
        }

        node_free((struct Node *) sect);
    }

    free(config);
}
