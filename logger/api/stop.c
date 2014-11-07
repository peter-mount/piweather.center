#include "lib/list.h"
#include "lib/threadpool.h"
#include "logger/logger.h"

/**
 * Start all enabled loggers
 * @param station
 */
void logger_stop() {
    struct Node *n = list_getHead(&loggers.loggers);
    while (list_isNode(n)) {
        struct logger *logger = (struct logger *) n;
        n = n->n_succ;

        if (logger->enabled && logger->stop)
            logger->stop(logger);
    }
}
