#include "lib/list.h"
#include "logger/logger.h"

/**
 * Log an entry
 * 
 * @param e log_entry
 */
void logger_log(time_t time, char *name, char *text, int value, int updated) {

    struct Node *n = list_getHead(&loggers.loggers);
    while (list_isNode(n)) {
        struct logger *l = (struct logger *) n;
        n = n->n_succ;

        if (l->enabled) {
            struct log_entry *e = logger_create_entry(time, name, text, value, updated);
            if (e)
                threadpool_submit(&l->threadPool, &e->node);
        }
    }
}
