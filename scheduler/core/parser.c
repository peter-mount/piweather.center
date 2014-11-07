#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include "lib/string.h"
#include "scheduler/scheduler.h"

static int getNumber(char *p, int *a, int max) {
    return (sscanf(p, "%d", a) == 0 || *a < 0 || *a > max) ? 1 : 0;
}

static int findRange(char *p, int *a, int *b, int max) {
    char *s = find(p, '-');

    if (*s != '-' || sscanf(p, "%d-%d", a, b) != 2 || *a < 0 || *a > max || *b < 0 || *b > max || *a>*b)
        return 1;

    return 0;
}

static int findMultiplier(char *p, int *b, int max) {
    char *s = find(p, '/');

    if (*s != '/')
        return 1;

    return getNumber(s + 1, b, max);
}

static int parseMin(char *def, SCHEDULE_ENTRY *e, int offset) {
    char *p = def, c;
    int i, a, b, m;
    if (*p == '*') {
        for (i = 0; i < 60; i++)
            scheduler_setBit(e, offset + i);
        return 0;
    }

    // Comma for multiple fields?
    char *q = find(p, ',');
    if (*q) {
        do {
            c = *q;
            *q = '\0';
            if (parseMin(p, e, offset))
                return i;
            *q = c;
            p = q + 1;
            q = find(p, ',');
        } while (*q);
    }

    // Look for /
    if (!findMultiplier(p, &m, 60)) {
        if (findRange(p, &a, &b, 60)) {
            // Just a/m i.e. repeat a every m
            if (getNumber(p, &a, 60))
                return 1;
            for (i = a; i < 60; i += m)
                scheduler_setBit(e, offset + i);
        } else {
            // a-b/m, i.e. repeat a-b every m
            for (; a < 60; a += m, b += m) {
                for (i = a; i <= b && i < 60; i++)
                    scheduler_setBit(e, offset + i);
            }
        }
        return 0;
    }

    // Look for -
    if (!findRange(p, &a, &b, 60)) {
        for (i = a; i < b; i++)
            scheduler_setBit(e, offset + i);
        return 0;
    }

    // Just a number
    if (getNumber(p, &a, 60))
        return 1;

    scheduler_setBit(e, offset + a);
    return 0;
}

static int parseHour(char *def, char *minDef, SCHEDULE_ENTRY *e) {
    char *p = def, c;
    int i, a, b, m;
    if (*p == '*') {
        for (i = 0; i < 24; i++)
            if (parseMin(minDef, e, i * 60))
                return 1;
        return 0;
    }

    // Comma for multiple fields?
    char *q = find(p, ',');
    if (*q) {
        do {
            c = *q;
            *q = '\0';
            if (parseHour(p, minDef, e))
                return 1;
            *q = c;
            p = q + 1;
            q = find(p, ',');
        } while (*q);
    }

    // Look for /
    if (!findMultiplier(p, &m, 24)) {
        if (findRange(p, &a, &b, 24)) {
            // Just a/m i.e. repeat a every m
            if (getNumber(p, &a, 24))
                return 1;
            for (i = a; i < 24; i += m)
                if (parseMin(minDef, e, i * 60))
                    return 1;
        } else {
            // a-b/m, i.e. repeat a-b every m
            for (; a < 24; a += m, b += m) {
                for (i = a; i <= b && i < 24; i++)
                    if (parseMin(minDef, e, i * 60))
                        return 1;
            }
        }
        return 0;
    }

    // Look for -
    if (!findRange(p, &a, &b, 24))
        for (i = a; i < b; i++)
            if (parseMin(minDef, e, i * 60))
                return 0;

    // Just a number
    if (getNumber(p, &a, 24))
        return 1;

    return parseMin(minDef, e, a * 60);
}

/**
 * Parses a definition to populate a schedule
 * @param e SCHEDULE_ENTRY
 * @param def definition
 * @return 1 on failure
 */
int scheduler_parse(SCHEDULE_ENTRY *e, char *def) {
    // We need to manipulate the string. If def is a hardcoded string in
    // code then we would segfault if we don't do this
    char *d = strdup(def);
    
    // first word is the minute definition
    char *ms = findNonWhitespace(d);
    char *me = findWhitespace(ms);
    if (!*me) {
        fprintf(stderr, "Invalid minute definition %s\n", def);
        return 1;
    }

    // second word is the hour definition
    char *hs = findNonWhitespace(me);
    char *he = findWhitespace(hs);

    // Terminate the works
    *me = '\0';
    *he = '\0';

    // Parse minute definition
    int r = parseHour(hs, ms, e);
    
    // Free our work string
    free(d);
    
    if (r)
        fprintf(stderr, "Invalid definition %s\n", def);

    return r;
}
