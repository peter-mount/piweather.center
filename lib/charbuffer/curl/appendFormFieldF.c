
#define _GNU_SOURCE
#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include "lib/charbuffer.h"

int charbuffer_append_form_fieldf(struct charbuffer *b, char *name, char *fmt, ...) {
    va_list argp;
    int ret = CHARBUFFER_OK;

    if (b->pos)
        ret = charbuffer_add(b, '&');

    if (!ret)
        ret = charbuffer_append(b, name);

    if (!ret)
        ret = charbuffer_add(b, '=');

    if (!ret) {
        char *c;
        va_start(argp, fmt);
        int s = vasprintf(&c, fmt, argp);
        if (s > 0) {
            ret = charbuffer_put(b, c, s);
            free(c);
        }
        va_end(argp);
    }

    return ret;
}

