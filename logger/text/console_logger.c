/**
 * A simple logger which logs to stdout, useful in debugging.
 * 
 * To enable, in the config:
 * 
 * [logging]
 * console enabled
 * 
 */
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lib/config.h"
#include "lib/list.h"
#include "logger/logger.h"

static void update(struct logger *logger, struct log_entry *entry) {
    struct tm timeinfo;
    localtime_r(&entry->time, &timeinfo);

    printf("%04d:%02d:%02d %02d:%02d:%02d %7s %32s %8d %s\n",
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1,
            timeinfo.tm_mday,
            timeinfo.tm_hour,
            timeinfo.tm_min,
            timeinfo.tm_sec,
            entry->updated ? "CHANGED" : "STABLE",
            entry->node.name,
            entry->value,
            entry->text
            );
}

void register_console_logger() {
    struct logger *logger = (struct logger *) malloc(sizeof (struct logger));
    memset(logger, 0, sizeof (struct logger));

    logger->node.name = strdup("console");
    logger->update = update;

    CONFIG_SECTION *sect = config_getSection("logging");
    if (sect) {
        config_getBooleanParameter(sect, "console", &logger->enabled);
    }

    logger_register(logger);
}
