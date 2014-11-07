/* 
 * File:   pw_app.h
 * Author: Peter T Mount
 *
 * Created on April 22, 2014, 5:55 PM
 */

#ifndef PW_APP_H
#define	PW_APP_H

/*
 * Common code for the imagefilter utility commands - DO NOT USE IN NORMAL CODE - unless you are doing something
 * similar with 2 or more images ;-)
 */

#include <stdlib.h>
#include <stdio.h>
#include <gd.h>
#include "filter.h"

extern gdImagePtr performFilter(gdImagePtr img1, gdImagePtr img2);

/*
 * A simple application which performs a single filter
 */

static void showHelp(const char *argv0) {
    fprintf(stderr, "%s two or more images together.\n\nSyntax: %s image1 ... imagen destImage\n", OPERATION, argv0);
    exit(1);
}

int main(const int argc, const char** argv) {
    gdImagePtr image1 = NULL, image2, image3 = NULL;
    int i, lastFileId = argc - 1;

    if (argc < 3)
        showHelp(argv[0]);

    for (i = 1; i < lastFileId; i++) {

        image2 = imagefilter_readFile((char *) argv[i]);

        if (image1) {
            // Perform image & make it the current one
            image3 = performFilter(image1, image2);
            gdImageDestroy(image1);
            gdImageDestroy(image2);
            image1 = image3;
        } else
            // First image
            image1 = image2;
    }

    if (image1) {
        imagefilter_writeFile(image1, (char *) argv[lastFileId]);
    } else {
        fprintf(stderr, "No output image generated\n");
        exit(1);
    }
}

#endif	/* PW_APP_H */

