#include <stdlib.h>
#include <string.h>

/**
 * Find s in p
 * @param p String to search
 * @param s String to locate
 * @return position of p or NULL if not present
 */
char *findString(char *p, char *s) {
    // NULL & zero length check
    if (!p || !s || !*p || !*s)
        return NULL;
    
    int pl = strlen(p);
    int sl = strlen(s);

    // Whilst we have enough chars left compare current position
    while (*p && pl >= sl) {
        if (strncasecmp(p, s, sl) == 0)
            return p;
        p++;
        pl--;
    }

    return NULL;
}