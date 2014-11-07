/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

int charbuffer_append_padright(struct charbuffer *b, char *src, int width) {
    int l = strlen(src);
    if (l > width)
        l = width;
    
    int r = charbuffer_put(b, src, l);
    
    int i = width - l;
    while (i>0) {
        charbuffer_add(b, ' ');
        i--;
    }
    
    return r;
}
