
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lib/config.h"
#include "lib/list.h"
#include "rabbitmqapi/rabbitmq.h"
#include "lib/bytebuffer.h"

int rabbitmq_publish_bytebuffer(struct rabbitmq *mq, char *routingKey, struct bytebuffer *buffer) {
    if (buffer->buffer && buffer->pos)
        return rabbitmq_publish_bytes(mq, routingKey, (void *) buffer->buffer, buffer->pos);
    else
        return -1;
}
