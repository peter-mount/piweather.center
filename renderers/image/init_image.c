#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "renderers/image/image.h"

void init_image(CONFIG_SECTION *sect, struct image *image, const char *prefix) {
    char tmp[256];

    image->prefix = prefix;
    
    snprintf(tmp, sizeof (tmp), "%s.enabled", prefix);
    config_getBooleanParameter(sect, tmp, &image->enabled);

    snprintf(tmp, sizeof (tmp), "%s.width", prefix);
    config_getIntParameter(sect, tmp, &image->width);

    snprintf(tmp, sizeof (tmp), "%s.height", prefix);
    config_getIntParameter(sect, tmp, &image->height);

    snprintf(tmp, sizeof (tmp), "%s.ident", prefix);
    config_getCharParameter(sect, tmp, &image->ident);

    snprintf(tmp, sizeof (tmp), "%s.path", prefix);
    config_getCharParameter(sect, tmp, &image->path);
}
