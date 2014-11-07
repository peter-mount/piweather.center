
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

gdImagePtr imagefilter_sun(gdImagePtr srcImage, int lim) {
    if(!lim)
        lim=240;
    
    int x,y,c,r,g,b;
    
    gdImagePtr img = gdImageCreateTrueColor(srcImage->sx, srcImage->sy);

    int blackColour = gdImageColorAllocate(img,0,0,0);
    
    for (y = 0; y < srcImage->sy; y++)
        for (x = 0; x < srcImage->sx; x++) {
            c = gdImageGetPixel(srcImage, x, y);
            r = gdImageRed(srcImage, c);
            g = gdImageGreen(srcImage, c);
            b = gdImageBlue(srcImage, c);

            if( g>=lim)
                c = blackColour;
            else
                c = gdImageColorAllocate(img,r,g,b);
            
            gdImageSetPixel(img,x,y,c);
        }
    
    return img;
}