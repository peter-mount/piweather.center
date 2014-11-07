
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

void fatalIfNull(void *v, char *fmt, ...) {
    if (!v) {
        va_list argp;

        va_start(argp, fmt);
        vfprintf(stderr, fmt, argp);
        va_end(argp);

        exit(1);
    }
}