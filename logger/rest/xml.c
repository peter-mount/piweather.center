/**
 * Handles JSON output in a rest service
 */

#include <stdlib.h>
#include "lib/charbuffer.h"
#include "logger/logger.h"

static char *TAG = "sensor";

/**
 * Utility to convert a log_entry into json.
 * 
 * Although primarily for rest it's also used by rabbitmq
 * 
 * @param b charbuffer
 * @param e log entry
 */
void log_entry_to_xml(struct charbuffer *b, struct log_entry *e) {
    charbuffer_reset_xml(b, TAG);
    charbuffer_append_xml(b, "hostid", loggers.hostid);
    charbuffer_append_xml(b, "name", e->node.name);
    charbuffer_append_xml(b, "text", e->text);
    charbuffer_append_xml(b, "value", "%d", e->value);
    charbuffer_append_xml(b, "timestamp", "%d", e->time);
    charbuffer_append_xml(b, "updated", e->updated ? "true" : "false");
    charbuffer_end_xml(b, TAG);
}
