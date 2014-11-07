
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>
#include "lib/bytebuffer.h"
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/list.h"
#include "logger/logger.h"
#include "logger/rest/rest-utils.h"
#include "rabbitmqapi/rabbitmq.h"

struct rabbitmq_logger {
    struct logger logger;
    struct rabbitmq mq;
    char *key_prefix;
    char *key_suffix;
};

/**
 * Our thread worker, this will pass data on to rabbitMQ
 * @param n
 * @param d
 */
static void update(struct logger *logger, struct log_entry *e) {
    struct rabbitmq_logger *mq = (struct rabbitmq_logger *) logger;

    charbuffer_reset(&mq->mq.buffer);

    if (mq->key_prefix)
        charbuffer_append(&mq->mq.buffer, mq->key_prefix);

    charbuffer_append(&mq->mq.buffer, e->node.name);

    if (mq->key_suffix)
        charbuffer_append(&mq->mq.buffer, mq->key_suffix);

    int len = 0;
    char *routingKey = charbuffer_toarray(&mq->mq.buffer, &len);

    // Map / in sensor names to . in the routing key
    char *p = routingKey;
    while (*p) {
        if (*p == '/') *p = '.';
        p++;
    }

    log_entry_to_json(&mq->mq.buffer,e);

    rabbitmq_publish_charbuffer(&mq->mq, routingKey, &mq->mq.buffer);
    
    free(routingKey);
}

void register_rabbitmq_logger() {
    struct rabbitmq_logger *logger = (struct rabbitmq_logger *) malloc(sizeof (struct rabbitmq_logger));
    memset(logger, 0, sizeof (struct rabbitmq_logger));

    logger->logger.node.name = strdup("rabbitmq");
    logger->logger.update = update;

    CONFIG_SECTION *sect = config_getSection("rabbitmq");
    if (sect) {
        config_getCharParameter(sect, "routingkey.prefix", &logger->key_prefix);
        config_getCharParameter(sect, "routingkey.suffix", &logger->key_suffix);

        // Enabled under logging?
        sect = config_getSection("logging");
        if (sect)
            config_getBooleanParameter(sect, "rabbitmq", &logger->logger.enabled);
    }

    rabbitmq_init(&logger->mq);
    logger_register(&logger->logger);
}