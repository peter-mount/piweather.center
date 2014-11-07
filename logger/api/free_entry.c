#include <stdlib.h>
#include "lib/list.h"
#include "logger/logger.h"

/**
 * Free a log_entry
 * 
 * @param e log_entry
 */
void logger_free_entry(struct log_entry *e) {
    if (e)
        node_free((struct Node *) e);
}