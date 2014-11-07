
#include <stdlib.h>
#include <stdio.h>
#include <gd.h>

#include "lib/string.h"

void imagefilter_writeFile(gdImagePtr img, char *n) {
    if (!n || !img)
        return;

    FILE *f = fopen(n, "w");
    if (!f)
        return;

    if (strendswith(n, ".gd"))
        gdImageGd(img, f);

    /*
        if( strendswith(n,".gd2"))
            gdImageGd2(image,f);
     */

    if (strendswith(n, ".gif"))
        gdImageGif(img, f);

    if (strendswith(n, ".jpg"))
        gdImageJpeg(img, f, 90);

    if (strendswith(n, ".png"))
        gdImagePng(img, f);

    /*
        if( strendswith(n,".bmp"))
            img = gdImageWBMP(image,f);
     */

    fclose(f);
}
