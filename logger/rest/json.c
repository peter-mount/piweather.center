/**
 * Handles JSON output in a rest service
 */

#include <stdlib.h>
#include "lib/charbuffer.h"
#include "logger/logger.h"

/**
 * Utility to convert a log_entry into json.
 * 
 * Although primarily for rest it's also used by rabbitmq
 * 
 * @param b charbuffer
 * @param e log entry
 */
void log_entry_to_json(struct charbuffer *b, struct log_entry *e) {
    charbuffer_reset_json(b);
    charbuffer_append_json(b, "hostid", loggers.hostid);
    charbuffer_append_json(b, "name", e->node.name);
    charbuffer_append_json(b, "text", e->text);
    charbuffer_append_json(b, "value", "%d", e->value);
    charbuffer_append_json(b, "timestamp", "%d000", e->time);
    charbuffer_append_json(b, "updated", e->updated ? "true" : "false");
    charbuffer_end_json(b);
}
