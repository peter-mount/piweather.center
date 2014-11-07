#include <stdlib.h>
#include <string.h>
#include "lib/list.h"
#include "lib/table.h"

static void free_row(TABLE_ROW *r) {
    struct Node *n = r->cells.l_head;
    while (list_isNode(n)) {
        TABLE_CELL *c = (TABLE_CELL *) n;
        n = n->n_succ;
        node_free(&c->node);
    }
    node_free(&r->node);
}

void table_destroy(TABLE *t) {
    struct Node *n = t->rows.l_head;
    while (list_isNode(n)) {
        TABLE_ROW *r = (TABLE_ROW *) n;
        n = n->n_succ;
        free_row(r);
    }
    free(t);
}
