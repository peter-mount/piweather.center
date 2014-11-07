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
#include <dirent.h>
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/list.h"
#include "lib/string.h"

CONFIG *config = NULL;

static char *newSection(CONFIG_SECTION **sect, char *p) {
    // Skip leading [
    p++;
    char *q = p;
    while (*q && *q != ']' && *q != '\n' && *q != '\r')
        q++;
    if (*q != ']') {
        fprintf(stderr, "Invalid section name found in %s\n", config->name);
        exit(1);
    }

    char *name = extractString(p, &q);

    *sect = config_getSection(name);

    if (!*sect) {
        *sect = (CONFIG_SECTION *) malloc(sizeof (CONFIG_SECTION));
        memset(*sect, 0, sizeof (CONFIG_SECTION));

        (*sect)->node.name = name;

        list_init(&(*sect)->parameters);

        list_addTail(&config->sections, (struct Node *) (*sect));

    }

    return findNextLine(q);
}

/**
 * Append new config parameter
 * @param config
 * @param param
 */
static int appendParameter(CONFIG_SECTION *sect, char *key, char *value) {
    CONFIG_PARAM *param = (CONFIG_PARAM *) malloc(sizeof (CONFIG_PARAM));
    if (!param) return 1;

    memset(param, 0, sizeof (CONFIG_PARAM));
    param->node.name = key;
    param->value = value;

    list_addTail(&sect->parameters, (struct Node *) param);

    return 0;
}

static char *parseParameter(CONFIG_SECTION *sect, char *p) {
    char *key, *value;

    // Find the key
    p = findNonWhitespace(p);
    key = p;
    if (!*p) return p;

    p = findWhitespace(p);
    if (!*p) return p;

    *p++ = '\0';

    // Now the value
    p = findNonWhitespace(p);
    value = p;
    if (!*p) return p;

    p = findEndOfLine(p);
    if (*p)
        *p++ = '\0';

    key = strdup(key);
    if (!key) return p;

    value = strdup(value);
    if (!value) {
        free(key);
        return p;
    }

    if (appendParameter(sect, key, value)) {

        free(value);
        free(key);
    }

    return p;
}

static void config_parse_charbuffer(struct charbuffer * b) {

    CONFIG_SECTION *sect = NULL;

    char *p = b->buffer;
    while (*p) {
        switch (*p) {
                // Comment
            case '#':
                p = findNextLine(p);
                break;

                // Skip blank lines
            case '\n':
            case '\r':
                p++;
                break;

                // [section] marker
            case '[':
                p = newSection(&sect, p);
                break;

                // Property
            default:
                if (!sect) {
                    fprintf(stderr, "Parameter outside of section in %s\n", config->name);
                    exit(1);
                }
                p = parseParameter(sect, p);

                break;
        }
    }
}

static void config_parse_file(struct charbuffer *buffer, char *name) {
    FILE *f = fopen(name, "r");
    if (!f) {
        fprintf(stderr, "Unable to read %s\n", name);
        exit(1);
    }

    charbuffer_reset(buffer);
    charbuffer_read(buffer, f);
    fclose(f);

    config_parse_charbuffer(buffer);
}

void config_parse_dir(char *name) {
    struct charbuffer buffer;
    charbuffer_init(&buffer);

    if (config)
        config_free();

    config = (CONFIG *) malloc(sizeof (CONFIG));
    memset(config, 0, sizeof (CONFIG));
    config->name = name;
    list_init(&config->sections);

    DIR *dir;
    struct dirent *ent;
    if ((dir = opendir(name)) != NULL) {
        while ((ent = readdir(dir)) != NULL) {
            if (ent->d_name[0] != '.') {
                char *f = (char *) malloc(strlen(name) + strlen(ent->d_name) + 2);
                sprintf(f, "%s/%s", name, ent->d_name);
                config_parse_file(&buffer, f);
                free(f);
            }
        }
        closedir(dir);
    } else {
        fprintf(stderr, "No config found in %s\n", name);
        exit(1);
    }

    charbuffer_free(&buffer);
}
