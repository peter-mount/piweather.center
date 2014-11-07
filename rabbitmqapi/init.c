
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lib/config.h"
#include "lib/list.h"
#include "rabbitmqapi/rabbitmq.h"
#include "lib/charbuffer.h"

#define AMQP_PORT 5672

void rabbitmq_init(struct rabbitmq *mq) {
    CONFIG_SECTION *sect = config_getSection("rabbitmq");
    if (sect) {
        config_getCharParameter(sect, "hostname", &mq->hostname);
        config_getIntParameter(sect, "port", &mq->port);
        config_getCharParameter(sect, "virtualhost", &mq->virtualhost);
        config_getCharParameter(sect, "username", &mq->username);
        config_getCharParameter(sect, "password", &mq->password);
        config_getCharParameter(sect, "exchange", &mq->exchange);

        if (mq->port < 1 || mq->port > 65535)
            mq->port = AMQP_PORT;

        if (!mq->virtualhost)
            mq->virtualhost = "/";

        if (!mq->username)
            mq->username = "guest";

        if (!mq->password)
            mq->password = "guest";

        if (!mq->exchange)
            mq->exchange = "amq.topic";

        charbuffer_init(&mq->buffer);
        rabbitmq_connect(mq);
    }
}
