/**
 * A sensor logger to submit sensor readings to iot.onl
 * 
 */
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/list.h"
#include "lib/string.h"
#include "logger/logger.h"
#include "logger/iot/iot.h"

extern int verbose;

struct iot_logger {
    struct logger logger;
    struct iot iot;
    struct charbuffer buffer;
};

static void update(struct logger *logger, struct log_entry *entry) {
    struct iot_logger *l = (struct iot_logger *) logger;

    int sensorId = iot_lookup_sensorId(&l->iot, &l->buffer, entry->node.name);

    if (sensorId)
        iot_update(&l->iot, &l->buffer, sensorId, entry->text, entry->value, entry->time);
}

void register_iotonl_logger() {
    struct iot_logger *logger = (struct iot_logger *) malloc(sizeof (struct iot_logger));
    memset(logger, 0, sizeof (struct iot_logger));

    logger->logger.node.name = strdup("iot");
    logger->logger.update = update;

    CONFIG_SECTION *sect = config_getSection("iot.onl");
    if (sect) {
        if (!iot_configure(sect, &logger->iot))
            fatalError("iot.onl config is incorrect\n");

        charbuffer_init(&logger->buffer);
        logger->logger.enabled = 1;
    }

    logger_register(&logger->logger);
}
