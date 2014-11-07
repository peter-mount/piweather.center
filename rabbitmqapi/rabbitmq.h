/* 
 * File:   rabbitmq.h
 * Author: Peter T Mount
 *
 * Created on April 4, 2014, 11:58 AM
 */

#ifndef RABBITMQ_H
#define	RABBITMQ_H

#include <stdint.h>
#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>
#include "lib/charbuffer.h"
#include "lib/bytebuffer.h"

struct rabbitmq {
    char *hostname;
    int port;
    char *virtualhost;
    char *username;
    char *password;
    char *exchange;
    // =================
    // Internal use only
    // =================
    struct charbuffer buffer;
    amqp_socket_t *socket;
    amqp_connection_state_t conn;
};

extern void rabbitmq_connect(struct rabbitmq *mq);
extern int rabbitmq_publish(struct rabbitmq *mq, char *routingKey, char *message);
extern int rabbitmq_publish_bytes(struct rabbitmq *mq, char *routingKey, void *message, int len);
extern int rabbitmq_publish_bytebuffer(struct rabbitmq *mq, char *routingKey, struct bytebuffer *buffer);
extern int rabbitmq_publish_charbuffer(struct rabbitmq *mq, char *routingKey, struct charbuffer *buffer);
extern void rabbitmq_init(struct rabbitmq *mq);

#endif	/* RABBITMQ_H */

