#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

void fatalError(char *fmt, ...) {
    va_list argp;

    va_start(argp, fmt);
    vfprintf(stderr, fmt, argp);
    va_end(argp);

    exit(1);
}
