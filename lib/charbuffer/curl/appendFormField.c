
#define _GNU_SOURCE
#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include "lib/charbuffer.h"

int charbuffer_append_form_field(struct charbuffer *b, char *name, char *value) {
    int ret = CHARBUFFER_OK;

    if (b->pos)
        ret = charbuffer_add(b, '&');

    if (!ret)
        ret = charbuffer_append(b, name);

    if (!ret)
        ret = charbuffer_add(b, '=');

    if (!ret)
        ret = charbuffer_append_urlencode(b, value);

    return ret;
}
