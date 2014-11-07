#include <stdlib.h>
#include <string.h>
#include <stdarg.h>
#include <math.h>
#include "lib/list.h"
#include "lib/table.h"

static TABLE_CELL *newCell(TABLE_ROW *r, char *content, TABLE_ALIGN_T align) {
    TABLE_CELL *c = (TABLE_CELL *) malloc(sizeof (TABLE_CELL));
    memset(c, 0, sizeof (TABLE_CELL));
    c->node.name = strdup(content);
    c->align = align;
    c->span = 1;
    c->width = strlen(content);
    list_addTail(&r->cells, &c->node);
    return c;
}

#define MAX_CELL_WIDTH 128

TABLE_CELL *table_blankCell(TABLE_ROW *r) {
    return newCell(r, "", TABLE_ALIGN_LEFT);
}

TABLE_CELL *table_addCell(TABLE_ROW *r, char *fmt, ...) {
    va_list argp;
    char tmp[MAX_CELL_WIDTH];
    va_start(argp, fmt);
    vsnprintf(tmp, MAX_CELL_WIDTH, fmt, argp);
    va_end(argp);
    return newCell(r, tmp, TABLE_ALIGN_LEFT);
}

TABLE_CELL *table_addCellCenter(TABLE_ROW *r, char *fmt, ...) {
    va_list argp;
    char tmp[MAX_CELL_WIDTH];
    va_start(argp, fmt);
    vsnprintf(tmp, MAX_CELL_WIDTH, fmt, argp);
    va_end(argp);
    return newCell(r, tmp, TABLE_ALIGN_CENTER);
}

TABLE_CELL *table_addCellRight(TABLE_ROW *r, char *fmt, ...) {
    va_list argp;
    char tmp[MAX_CELL_WIDTH];
    va_start(argp, fmt);
    vsnprintf(tmp, MAX_CELL_WIDTH, fmt, argp);
    va_end(argp);
    return newCell(r, tmp, TABLE_ALIGN_RIGHT);
}

TABLE_CELL *table_add_hm(TABLE_ROW *r, double t) {
    div_t d = div((int) round(t * 60.0), 60);
    return table_addCellRight(r, "%02d:%02d", d.quot, d.rem);
}

TABLE_CELL *table_add_hms(TABLE_ROW *r, double t) {
    div_t d = div((int) round(t * 3600.0), 3600);
    int h = d.quot;
    d = div(d.rem, 60);
    return table_addCellRight(r, "%02d:%02d:%02d", h, d.quot, d.rem);
}

TABLE_CELL *table_add_dms(TABLE_ROW *r, double t, char p, char n) {
    double v;
    int s;
    if (t < 0.0) {
        v = -t;
        s = n;
    } else {
        v = t;
        s = p;
    }
    div_t d = div((int) round(v * 3600.0), 3600);
    int h = d.quot;
    d = div(d.rem, 60);
    return table_addCellRight(r, "%3dd %02dm %02ds %c", h, d.quot, d.rem, s);
}
