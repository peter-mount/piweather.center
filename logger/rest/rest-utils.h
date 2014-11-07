/* 
 * File:   rest.h
 * Author: Peter T Mount
 *
 * Created on April 6, 2014, 9:24 AM
 */

#ifndef REST_H
#define	REST_H

#include "lib/charbuffer.h"
#include "logger/logger.h"

extern void log_entry_to_json(struct charbuffer *b, struct log_entry *e);
extern void log_entry_to_xml(struct charbuffer *b, struct log_entry *e);

#endif	/* REST_H */

