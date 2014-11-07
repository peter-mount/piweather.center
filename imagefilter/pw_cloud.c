#include <stdlib.h>
#include <stdio.h>
#include <gd.h>
#include "filter.h"

#define OPERATION "Cloud"

gdImagePtr performFilter(gdImagePtr img1, gdImagePtr img2) {
    return imagefilter_apply_mask(img1, img2);
}

// This must be last
#include "pw_app.h"

