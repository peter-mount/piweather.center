#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
#include "lib/charbuffer.h"

static char *START = "{";
#define START_SIZE 1

/**
 * Resets a charbuffer ready for generating json
 * 
 * @param b charbuffer
 */
void charbuffer_reset_json(struct charbuffer *b) {
    charbuffer_reset(b);
    charbuffer_put(b, START, START_SIZE);
}

