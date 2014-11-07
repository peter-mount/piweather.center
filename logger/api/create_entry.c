
#include <stdlib.h>
#include <string.h>
#include "lib/list.h"
#include "logger/logger.h"

/**
 * Free a log_entry
 * 
 * @param e log_entry
 */
struct log_entry *logger_create_entry(time_t time, char *name, char *text, int value, int updated) {
    struct log_entry *e = (struct log_entry *) malloc(sizeof (struct log_entry));
    if (e) {
        memset(e, 0, sizeof (struct log_entry));

        e->node.name = strdup(name);

        e->time = time;

        strncpy(e->text, text, LOG_TEXT_SIZE);
        e->text[LOG_TEXT_SIZE - 1] = '\0';

        e->value = value;
        e->updated = updated;
    }
    return e;
}