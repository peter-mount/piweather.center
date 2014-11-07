/*
 * A R/B filter used to identify clouds from sky
 */
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

/**
 * A filter which generates an image showing cloud cover.
 * 
 * rblim is the limit (0..1) above which a pixel is deemed to be cloud.
 * 
 * If not NULL, pixTotal & pixCloud will hold totals of pixels not black and cloud
 * respectively.
 * 
 * i.e. to get percentage cloud cover:
 * 
 *   double cloudPC=0.0;
 *   if(pixTotal>0)
 *     cloudPC = 100.0*pixCloud/pixTotal;
 * 
 * @param srcImage source image
 * @param rblim rb limit for cloud, DEFAULT_RBLIM is a decent value
 * @param cloud cloud colour in image
 * @param sky sky colour in image
 * @param black black colour in image
 * @param pixTotal int to store total pixels not black, NULL for none
 * @param pixCloud int to store total pixels that are cloud, NULL for none
 * @return Filtered image
 */
gdImagePtr imagefilter_rb_ratio(gdImagePtr srcImage,
        double rblim,
        uint32_t cloud, uint32_t sky, uint32_t black,
        int *pixTotal,
        int *pixCloud) {
    int x, y;
    
    // Count of cloud/sky pixels
    int cloudy = 0, total = 0;
    int c,r,g,b,np;
    double rb;

    gdImagePtr clouds = gdImageCreateTrueColor(srcImage->sx, srcImage->sy);

    int cloudColour = gdImageColorAllocate(clouds,
            (cloud & 0xff0000) >> 16,
            (cloud & 0xff00) >> 8,
            (cloud & 0xff));

    int skyColour = gdImageColorAllocate(clouds,
            (sky & 0xff0000) >> 16,
            (sky & 0xff00) >> 8,
            (sky & 0xff));

    int blackColour = gdImageColorAllocate(clouds,
            (black & 0xff0000) >> 16,
            (black & 0xff00) >> 8,
            (black & 0xff));
    
    for (y = 0; y < srcImage->sy; y++)
        for (x = 0; x < srcImage->sx; x++) {
            c = gdImageGetPixel(srcImage, x, y);
            r = gdImageRed(srcImage, c);
            g = gdImageGreen(srcImage, c);
            b = gdImageBlue(srcImage, c);

            np = blackColour;

            if (b) {
                // We have some blue so we are either sky or cloud
                rb = (double) r / (double) b;
                if (rb >= rblim) {
                    np = cloudColour;
                    cloudy++;
                } else if (rb >= 0.1)
                    np = skyColour;
            }

            if (np != blackColour)
                total++;

            gdImageSetPixel(clouds, x, y, np);
        }

    if (pixTotal)
        *pixTotal = total;

    if (pixCloud)
        *pixCloud = cloudy;

    return clouds;
}