/* 
 * File:   config.h
 * Author: peter
 *
 * Created on February 27, 2014, 12:35 PM
 */

#ifndef CONFIG_H
#define	CONFIG_H

#ifdef	__cplusplus
extern "C" {
#endif

#include <stdint.h>
#include "lib/list.h"

    /*
     * Structure used to define the parameters in a config file.
     * 
     */
    typedef struct {
        // The ID of this key
        int id;
        // The name of the parameter
        const char *key;
        // Does it have an argument
        int argc;
        // Optional?
        int optional;
    } CONFIG_ENTRY;

    /*
     * A parsed config parameter
     */
    typedef struct config_param CONFIG_PARAM;

    struct config_param {
        struct Node node;
        const char *value;
    };

    typedef struct config_section CONFIG_SECTION;

    struct config_section {
        struct Node node;
        struct List parameters;
    };

    typedef struct {
        char *name;
        struct List sections;
    } CONFIG;
    
    extern CONFIG *config;

    /**
     * Frees up memory used by configuration
     * @param config
     */
    extern void config_free();

    /**
     * Retrieves the named section
     * 
     * @param config
     * @param name
     * @return null if not found
     */
    extern CONFIG_SECTION *config_getSection(const char *name);

    extern CONFIG_PARAM *config_getParameter(CONFIG_SECTION *sect, const char *name);
    extern int config_scanParameter(CONFIG_SECTION *sect, const char *name, const char *fmt, void *val);
    extern int config_getIntParameter(CONFIG_SECTION *sect, const char *name, int *val);
    extern int config_getCharParameter(CONFIG_SECTION *sect, const char *name, char **val);
    extern void config_getBooleanParameter(CONFIG_SECTION *sect, const char *name, int *val);
    extern int config_getHexParameter(CONFIG_SECTION *sect, const char *name, uint32_t *val);
    extern int config_getHexLongParameter(CONFIG_SECTION *sect, const char *name, uint64_t *val);
    extern int config_getLongParameter(CONFIG_SECTION *sect, const char *name, long *val);
    extern int config_getFloatParameter(CONFIG_SECTION *sect, const char *name, float *val);
    extern int config_getDoubleParameter(CONFIG_SECTION *sect, const char *name, double *val);
    extern void config_parse_dir(char *name);
#ifdef	__cplusplus
}
#endif

#endif	/* CONFIG_H */

