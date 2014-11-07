
#include <stdlib.h>
#include <stdio.h>
#include <gd.h>

#include "lib/string.h"

gdImagePtr imagefilter_readFile( char *n ) {
    if(!n)
        return NULL;
    
    FILE *f = fopen(n,"r");
    if(!f)
        return NULL;
    
    gdImagePtr img = NULL;
    
    if( strendswith(n,".gd"))
        img = gdImageCreateFromGd(f);
    
    if( strendswith(n,".gd2"))
        img = gdImageCreateFromGd2(f);
    
    if( strendswith(n,".gif"))
        img = gdImageCreateFromGif(f);
    
    if( strendswith(n,".jpg"))
        img = gdImageCreateFromJpeg(f);
    
    if( strendswith(n,".png"))
        img = gdImageCreateFromPng(f);
    
    if( strendswith(n,".bmp"))
        img = gdImageCreateFromWBMP(f);
    
    if( strendswith(n,".xbm"))
        img = gdImageCreateFromXbm(f);
    
    fclose(f);
    return img;
}
