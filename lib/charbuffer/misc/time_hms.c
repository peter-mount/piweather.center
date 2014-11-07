#include <stdlib.h>
#include <math.h>
#include "lib/charbuffer.h"

/**
 * Append time to a charbuffer as HH:MM
 * @param b buffer
 * @param m time to render in minutes
 */
void charbuffer_time_hms(struct charbuffer *b, double t) {
    div_t d = div((int) round(t * 3600.0), 3600);
    charbuffer_append_int(b, d.quot, 2);
    charbuffer_add(b, ':');

    d = div(d.rem, 60);
    charbuffer_append_int(b, d.quot, 2);
    charbuffer_add(b, ':');
    charbuffer_append_int(b, d.rem, 2);
}
