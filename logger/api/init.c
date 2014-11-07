#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lib/config.h"
#include "lib/list.h"
#include "logger/logger.h"

struct loggers loggers;

void logger_init() {
    memset(&loggers, 0, sizeof (loggers));
    list_init(&loggers.loggers);

    // The hostid, either defined in config or taken from the local host name
    loggers.hostid = NULL;
    CONFIG_SECTION *sect = config_getSection("logging");
    if (sect)
        config_getCharParameter(sect, "hostid", &loggers.hostid);

    if (!loggers.hostid) {
        char tmp[1024];
        gethostname(tmp, 1024);
        loggers.hostid = strdup(tmp);
    }
}
