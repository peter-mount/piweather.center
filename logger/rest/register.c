
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <inttypes.h>
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/hashmap.h"
//#include "lib/history.h"
#include "logger/logger.h"
#include "logger/rest/rest.h"
#include "logger/rest/rest-utils.h"
#include "webserver/webserver.h"
#include "sensors/sensors.h"
#include "lib/history.h"

static const char *REST_BASE = "sensor/current/";

/*
static const char *HIST_BASE = "sensor/history/";
 */

static const char *addRest(struct rest_logger *l, const char *n, const char *suffix) {
    charbuffer_reset(&l->buffer);
    charbuffer_append(&l->buffer, webserver.contextPath ? webserver.contextPath : "/");
    charbuffer_append(&l->buffer, (char *) REST_BASE);
    charbuffer_append(&l->buffer, (char *) n);
    charbuffer_append(&l->buffer, (char *) suffix);

    int len;
    char *url = charbuffer_tostring(&l->buffer, &len);

    webserver_add_response_handler(url);
    return url;
}

/*
static void addHistory(struct rest_logger *l, const char *n, struct MHD_Response **response) {
    charbuffer_reset(&l->buffer);
    charbuffer_append(&l->buffer, webserver.contextPath ? webserver.contextPath : "/");
    charbuffer_append(&l->buffer, (char *) HIST_BASE);
    charbuffer_append(&l->buffer, (char *) n);
    charbuffer_append(&l->buffer, (char *) ".json");

    int len;
    char *url = charbuffer_tostring(&l->buffer, &len);

    webserver_add_response_handler(url, response);
}
 */

static struct rest_service *getService(struct rest_logger *l, char *n) {
    struct rest_service *s = hashmapGet(l->services, n);
    if (!s) {
        s = (struct rest_service *) malloc(sizeof (struct rest_service));
        memset(s, 0, sizeof (struct rest_service));

        s->json = addRest(l, n, ".json");
        s->text = addRest(l, n, ".txt");
        s->xml = addRest(l, n, ".xml");

        /*
                // JSON History of last 24 hours data
                addHistory(l, n, &s->histJson);
                time_t age = 86400;
                history_init(&s->history, age);
         */

        hashmapPut(l->services, n, s);
    }
    return s;
}

/**
 * Init hook.
 * 
 * All we do is pre-create rest services for all enabled sensors so the endpoints exist in the webserver.
 * 
 * If additional sensors come online later, then they'll get created as required.
 * 
 * @param logger
 */
static void init(struct logger *logger) {
    struct rest_logger *l = (struct rest_logger *) logger;

    struct sensor *s = sensors->sensors;
    while (s) {
        getService(l, (char *) s->name);
        s = s->next;
    }
}

static void update(struct logger *logger, struct log_entry *e) {
    struct rest_logger *l = (struct rest_logger *) logger;

    struct rest_service *s = getService(l, e->node.name);

    // Text rest service is just the text field
    charbuffer_reset(&l->buffer);
    charbuffer_append(&l->buffer, e->text);
    replaceResponseCharBuffer(s->text, &l->buffer, "text/plain");

    // xml rest service
    log_entry_to_xml(&l->buffer, e);
    replaceResponseCharBuffer(s->xml, &l->buffer, "text/xml");

    // json rest service
    log_entry_to_json(&l->buffer, e);
    replaceResponseCharBuffer(s->json, &l->buffer, "text/json");

    /*
        // Also append json of time (java epoch) and raw value
        int len;
        char tmp[32];
        snprintf(tmp,32,"[%ld000,%d]", e->time, e->value);
        struct HistoryNode *hn = (struct HistoryNode *) malloc(sizeof (struct HistoryNode));
        memset(hn, 0, sizeof (struct HistoryNode));
        hn->time = e->time;
        hn->node.name = strdup(tmp);
        history_add(&s->history, hn);

        // Recreate the history json as a list of all current entries
        history_expire(&s->history);
        charbuffer_reset(&l->buffer);
        charbuffer_add(&l->buffer, '[');
        len = 0;
        struct Node *n = s->history.list.l_head;
        while (list_isNode(n)) {
            if (len)
                charbuffer_add(&l->buffer, ',');
            charbuffer_append(&l->buffer, n->name);
            len = 1;
            n = n->n_succ;
        }
        charbuffer_add(&l->buffer, ']');
        replaceResponseCharBuffer(&s->histJson, &l->buffer, "text/json");
     */
}

void register_rest_logger() {
    struct rest_logger *logger = (struct rest_logger *) malloc(sizeof (struct rest_logger));
    memset(logger, 0, sizeof (struct rest_logger));

    logger->logger.node.name = strdup("rest");
    logger->logger.init = init;
    logger->logger.update = update;

    charbuffer_reset(&logger->buffer);

    // Hashmap of sensor name's
    logger->services = hashmapCreate(10, hashmapStringHash, hashmapStringEquals);

    CONFIG_SECTION *sect = config_getSection("webserver");
    if (sect) {
        // By default if webserver is present then rest is enabled unless disabled in config
        logger->logger.enabled = 1;
        config_getBooleanParameter(sect, "rest", &logger->logger.enabled);
    }

    logger_register(&logger->logger);
}
