
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "rabbitmqapi/rabbitmq.h"

int rabbitmq_publish_bytes(struct rabbitmq *mq, char *routingKey, void *message, int len) {
    amqp_bytes_t msg;

    msg.len = len;
    msg.bytes = message;


    return amqp_basic_publish(
            mq->conn,
            1,
            amqp_cstring_bytes(mq->exchange),
            amqp_cstring_bytes(routingKey),
            0,
            0,
            NULL,
            msg
            );
}
