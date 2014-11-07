
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lib/config.h"
#include "lib/list.h"
#include "logger/logger.h"
#include "rabbitmqapi/rabbitmq.h"
#include "lib/charbuffer.h"

extern int verbose;

void rabbitmq_connect(struct rabbitmq *mq) {
    if (!mq->conn) {
        
        if (verbose)
            fprintf(stderr, "RabbitMQ connecting to %s %s\n",
                    mq->hostname,
                    mq->virtualhost);

        mq->conn = amqp_new_connection();

        mq->socket = amqp_tcp_socket_new(mq->conn);

        if (!mq->socket) {
            fprintf(stderr, "Unable to connect to socket\n");
            amqp_destroy_connection(mq->conn);
            mq->socket = NULL;
            mq->conn = NULL;
            return;
        }

        if (amqp_socket_open(mq->socket, mq->hostname, mq->port )) {
            fprintf(stderr, "failed to connect\n");
            amqp_destroy_connection(mq->conn);
            mq->socket = NULL;
            mq->conn = NULL;
            return;
        }
        
        amqp_rpc_reply_t reply = amqp_login(mq->conn,
                mq->virtualhost,
                0, 131072, 0,
                AMQP_SASL_METHOD_PLAIN,
                mq->username,
                mq->password);

        
        if (reply.reply_type != AMQP_RESPONSE_NORMAL) {
            fprintf(stderr, "failed to login\n");
            amqp_destroy_connection(mq->conn);
            mq->socket = NULL;
            mq->conn = NULL;
            return;
        }

        amqp_channel_open(mq->conn, 1);

        if (amqp_get_rpc_reply(mq->conn).reply_type != AMQP_RESPONSE_NORMAL) {
            fprintf(stderr, "failed to open channel");
            amqp_destroy_connection(mq->conn);
            mq->socket = NULL;
            mq->conn = NULL;
            return;
        }

        if (verbose)
            fprintf(stderr, "RabbitMQ connected\n");
    }
}
