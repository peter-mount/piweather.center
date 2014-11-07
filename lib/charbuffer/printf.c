/**
 * Handles an extensible reusable char buffer
 */

#define _GNU_SOURCE
#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include "lib/charbuffer.h"

int charbuffer_printf(struct charbuffer *b, char *fmt, ...) {
    va_list argp;
    int ret;
    char *c;
    va_start(argp, fmt);
    int s = vasprintf(&c, fmt, argp);
    if (s > 0) {
        ret = charbuffer_put(b, c, s);
        free(c);
    }
    va_end(argp);
    return ret;
}

