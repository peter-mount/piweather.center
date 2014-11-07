#include <stdlib.h>
#include <string.h>
#include "lib/list.h"
#include "lib/table.h"

void table_format(TABLE *t) {
    struct Node *rn, *cn;
    TABLE_ROW *r;
    TABLE_CELL *c;
    int s, i;

    // Work out width in cells
    rn = t->rows.l_head;
    while (list_isNode(rn)) {
        r = (TABLE_ROW *) rn;
        rn = rn->n_succ;

        s = 0;
        cn = r->cells.l_head;
        while (list_isNode(cn)) {
            c = (TABLE_CELL *) cn;
            s += c->span;
            cn = cn->n_succ;
        }
        if (s > t->max_cells)
            t->max_cells = s;
    }
    if (t->width)
        free(t->width);
    t->width = (int *) calloc(t->max_cells, sizeof (int));

    // Now work out the cell widths
    rn = t->rows.l_head;
    while (list_isNode(rn)) {
        r = (TABLE_ROW *) rn;
        rn = rn->n_succ;

        i = 0;
        cn = r->cells.l_head;
        while (list_isNode(cn)) {
            c = (TABLE_CELL *) cn;
            if (c->span == 1 && c->width > t->width[i])
                t->width[i] = c->width;
            i += c->span;
            cn = cn->n_succ;
        }
    }
}