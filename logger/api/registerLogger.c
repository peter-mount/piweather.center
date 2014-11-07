#include "lib/list.h"
#include "logger/logger.h"

void logger_register(struct logger *logger) {
    list_addTail(&loggers.loggers, &logger->node);
}
