
#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

#include <ctype.h>

static char hex[] = "0123456789abcdef";

/**
 * Appends a string to a charbuffer urlencoding it in the process
 */
int charbuffer_append_urlencode(struct charbuffer *b, char *src) {
    int ret = CHARBUFFER_OK;
    char tmp[3];
    char *p = src;
    while (*p && ret == CHARBUFFER_OK) {
        if (isalnum(*p) || *p == '-' || *p == '_' || *p == '.' || *p == '~')
            ret = charbuffer_add(b, *p);
        else if (*p == ' ')
            ret = charbuffer_add(b, '+');
        else {
            tmp[0] = '%';
            tmp[1] = hex[*p >> 4];
            tmp[2] = hex[*p & 15];
            ret = charbuffer_put(b, tmp, 3);
        }
        p++;
    }
    return ret;
}
