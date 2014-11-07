#include <stdlib.h>
#include <math.h>
#include "lib/charbuffer.h"

/**
 * Append time to a charbuffer as HH:MM
 * @param b buffer
 * @param m time to render in minutes
 */
void charbuffer_time_hm(struct charbuffer *b, double t) {
    div_t d = div((int) round(t * 60.0), 60);
    charbuffer_append_int(b, d.quot, 2);
    charbuffer_add(b, ':');
    charbuffer_append_int(b, d.rem, 2);
}

