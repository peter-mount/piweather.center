
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

gdImagePtr imagefilter_createThumbnail(gdImagePtr srcImg, int w, int h) {
    gdImagePtr image = gdImageCreateTrueColor(w, h);

    if (w != srcImg->sx && h != srcImg->sy)
        gdImageCopyResized(image, srcImg,
            0, 0,
            0, 0,
            w, h,
            srcImg->sx, srcImg->sy);
    else
        gdImageCopy(image, srcImg, 0, 0, 0, 0, w, h);
    
    return image;
}