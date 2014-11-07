#include "lib/list.h"
#include "lib/threadpool.h"
#include "logger/logger.h"

/**
 * Our thread worker, this will pass data on to the logger
 * @param n
 * @param d
 */
static void *worker(struct Node *n, void *d) {
    struct log_entry *e = (struct log_entry *) n;
    struct logger *l = (struct logger *) d;

    l->update(l, e);

    node_free((struct Node *) e);
}

/**
 * Start all enabled loggers
 * @param station
 */
void logger_start() {
    struct Node *n = list_getHead(&loggers.loggers);
    while (list_isNode(n)) {
        struct logger *logger = (struct logger *) n;
        n = n->n_succ;

        if (logger->enabled) {
            threadpool_init(&logger->threadPool, 1);
            threadpool_start(&logger->threadPool, worker, (void *) logger);

            if (logger->init)
                logger->init(logger);
        }
    }
}
