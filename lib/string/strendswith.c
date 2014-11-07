#include <stdlib.h>
#include <string.h>

/**
 * Test to see if string s ends with string p
 * @param s String to test
 * @param p String to locate
 * @return 1 if s ends with p, 0 otherwise
 */
int strendswith(char *s, char *p) {
    // Both same then a match
    if (s == p)
        return 1;

    // Either null then no match
    if (!s || !p)
        return 0;

    int sl = strlen(s);
    int pl = strlen(p);

    // Either is 0 then match only if both are 0
    if (sl == 0 || pl == 0)
        return sl == pl;

    // no match if s is shorter than p
    if (sl < pl)
        return 0;

    // Find start point in source string then compare
    int i = sl - pl;
    return i < 0 ? 0 : strncmp(s + i, p, pl) == 0;
}
