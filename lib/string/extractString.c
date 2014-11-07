
#include <string.h>

char *extractString(char *start, char **end) {
    // Terminate the string, move forward then dup
    **end = '\0';
    *end = *end + 1;

    return strdup(start);
}
