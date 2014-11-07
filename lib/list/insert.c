
#include <stdlib.h>
#include "lib/list.h"

void list_insert(struct List *l, struct Node *n, struct Node *after) {
    if (after == NULL) {
        list_addHead(l, n);
        return;
    }

    n->n_pred = after;
    n->n_succ = after->n_succ;
    after->n_succ = n;
    n->n_succ->n_pred = n;
}
