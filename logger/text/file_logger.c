/**
 * A logger which appends to a file.
 * 
 * In this logger configuration consists of the path to a directory where the logging is to be placed.
 * 
 * Then, for each entry the path is augmented with the sensor name then followed by year/month. The day of the month
 * is then the final log file
 */
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <errno.h>
#include <unistd.h>
#include "lib/config.h"
#include "lib/file.h"
#include "lib/list.h"
#include "logger/logger.h"

struct state {
    struct logger logger;
    // The path to where logging is to be placed
    char *path;
};

static void update(struct logger *logger, struct log_entry *entry) {
    struct state *state = (struct state *) logger;

    char tmp[256];
    struct tm timeinfo;
    localtime_r(&entry->time, &timeinfo);

    // Make sure the directory exists
    snprintf(tmp, sizeof (tmp),
            "%s/%s/%04d/%02d",
            state->path,
            entry->node.name,
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1
            );
    mkdirs(tmp, S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH);

    // Now the file
    snprintf(tmp, sizeof (tmp),
            "%s/%s/%04d/%02d/%02d",
            state->path,
            entry->node.name,
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1,
            timeinfo.tm_mday
            );
    FILE *file = fopen(tmp, "a");

    if (file) {
        fprintf(file,
                "%04d:%02d:%02d %02d:%02d:%02d %7s %s %d %s\n",
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
        fclose(file);
        sync();
    } else {
        fprintf(stderr, "cannot open log %s: %s", tmp, strerror(errno));
    }
}

void register_file_logger() {
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    state->logger.node.name = strdup("file");
    state->logger.update = update;

    CONFIG_SECTION *sect = config_getSection("logging");
    if (sect) {
        config_getCharParameter(sect, "file-path", &state->path);

        state->logger.enabled = state->path && *state->path;
    }

    logger_register((struct logger *) state);
}
