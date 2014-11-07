#include <stdlib.h>
#include <stdio.h>
#include <gd.h>
#include "filter.h"

#define OPERATION "Merge"

gdImagePtr performFilter(gdImagePtr img1, gdImagePtr img2) {
    return imagefilter_merge(img1, img2);
}

// This must be last
#include "pw_app.h"

