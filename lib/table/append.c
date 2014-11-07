#include <stdlib.h>
#include <string.h>
#include "lib/charbuffer.h"
#include "lib/list.h"
#include "lib/table.h"

void table_append(struct charbuffer *b, TABLE *t) {
    struct Node *rn, *cn;
    TABLE_ROW *r;
    TABLE_CELL *c;
    int s, i;

    rn = t->rows.l_head;
    while (list_isNode(rn)) {
        r = (TABLE_ROW *) rn;

        i = 0;
        cn = r->cells.l_head;
        while (list_isNode(cn)) {
            if (i)
                charbuffer_add(b, ' ');

            c = (TABLE_CELL *) cn;

            // The cell width accounting for spanning
            int w = -1; // -1 to account for adding a space afterwards
            for (s = 0; s < c->span; s++)
                w += t->width[i++]+1;

            if (c->align == TABLE_ALIGN_LEFT)
                charbuffer_append_padright(b, c->node.name, w);
            else if (c->align == TABLE_ALIGN_RIGHT)
                charbuffer_append_padleft(b, c->node.name, w);
            else
                charbuffer_append_center(b, c->node.name, w);

            cn = cn->n_succ;
        }

        charbuffer_add(b, '\n');
        rn = rn->n_succ;
    }
}
