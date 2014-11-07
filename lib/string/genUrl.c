
#include <stdlib.h>
#include <string.h>

/**
 * Util to generate a simple url from two strings
 * @param contextPath ContextPath or prefix
 * @param suffix Suffix to append
 * @return url
 */
char* genurl(const char *contextPath, const char *suffix) {
    int len = strlen(contextPath) + strlen(suffix);
    char *url = (char *) malloc(len + 1);
    strcpy(url, contextPath);
    strcat(url, suffix);

    return url;
}
