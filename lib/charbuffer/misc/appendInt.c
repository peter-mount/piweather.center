#include <stdlib.h>
#include <sys/types.h>
#include <stdbool.h>
#include "lib/charbuffer.h"

#define INT_DIGITS 19		/* enough for 64 bit integer */

void charbuffer_append_int(struct charbuffer *b, int v, int width) {
    char buf[INT_DIGITS];
    int i = 0;
    int m = width ? width : INT_DIGITS;
    div_t d;
    bool neg = v < 0;
    d.quot = neg ? -v : v;
    do {
        d = div(d.quot, 10);
        buf[i++] = '0' + d.rem;
    } while (i < m && d.quot);

    while (i < width && i < m)
        buf[i++] = '0';

    if (neg && i < m)
        buf[i++] = '-';

    while (i)
        charbuffer_add(b, buf[--i]);
}
