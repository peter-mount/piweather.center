#include <stdlib.h>
#include <stdio.h>
#include <gd.h>
#include "filter.h"

int main(const int argc, const char** argv) {

    if (argc < 3 || argc > 4) {
        fprintf(stderr, "sun source [lim] dest\n");
    }

    gdImagePtr src = imagefilter_readFile((char *) argv[1]);
    if (!src) {
        fprintf(stderr, "Cannot read %s\n", argv[1]);
        exit(1);
    }

    int lim = 0;
    if (argc == 4)
        if ((sscanf(argv[2], "%d", &lim) != 1 || lim > 255)) {
            fprintf(stderr, "Invalid lim \"%s\"\n", argv[2]);
            exit(1);
        }

    gdImagePtr dst = imagefilter_sun(src, lim);

    if (dst) {
        imagefilter_writeFile(dst, (char *) argv[argc - 1]);
    } else {
        fprintf(stderr, "No output image generated\n");
        exit(1);
    }
}
