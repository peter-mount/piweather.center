#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
#include "lib/charbuffer.h"

/**
 * Appends to a charbuffer being used to generate json field
 * 
 * @param b charbuffer
 * @param n field name
 * @param fmt field value or format
 * @param ... args when fmt is a format
 */
void charbuffer_append_json(struct charbuffer *b, char *n, char *fmt, ...) {
    va_list argp;

    // If not first entry then , separator
    if (b->pos > 1)
        charbuffer_put(b, ",", 1);
    
    charbuffer_put(b, "\"", 1);
    charbuffer_append(b, n);
    charbuffer_append(b, "\":\"");

    char tmp[128];
    va_start(argp, fmt);
    vsnprintf(tmp, 128, fmt, argp);
    charbuffer_append(b, tmp);
    va_end(argp);

    charbuffer_put(b, "\"", 1);
}

