
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "rabbitmqapi/rabbitmq.h"

int rabbitmq_publish(struct rabbitmq *mq, char *routingKey, char *message) {
    return rabbitmq_publish_bytes(mq, routingKey, (void *) message, strlen(message));
}
