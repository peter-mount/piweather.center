#include <stdlib.h>
#include <stdio.h>
#include <sys/stat.h>
#include <string.h>
#include <errno.h>
#include <gd.h>
#include "lib/config.h"
#include "lib/file.h"
#include "renderers/image/image.h"

extern int verbose;

void write_image(gdImagePtr img, time_t *date, struct image *image) {
    char tmp[256];
    struct tm timeinfo;
    gmtime_r(date, &timeinfo);

    // Make sure the directory exists
    snprintf(tmp, sizeof (tmp),
            "%s/%04d/%02d/%02d",
            image->path,
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1,
            timeinfo.tm_mday);
    mkdirs(tmp, S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH);

    // Try to create the file
    snprintf(tmp, sizeof (tmp),
            "%s/%04d/%02d/%02d/%02d-%02d-%02d.jpg",
            image->path,
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1,
            timeinfo.tm_mday,
            timeinfo.tm_hour,
            timeinfo.tm_min,
            timeinfo.tm_sec);

    FILE *file = fopen(tmp, "w");
    if (!file) {
        char e[256];
        strerror_r(errno, e, sizeof (e));
        fprintf(stderr, "Unable to write %s: %s", tmp, e);
    } else {
        if (verbose)
            fprintf(stderr, "Writing %s\n", tmp);
        gdImageJpeg(img, file, 90);
        fclose(file);
    }
}
