
#include <stdlib.h>
#include "lib/list.h"

struct Node *list_removeHead(struct List *l) {
    struct Node *n = NULL;

    if (!list_isEmpty(l)) {
        n = l->l_head;
        n->n_succ->n_pred = n->n_pred;
        n->n_pred->n_succ = n->n_succ;
    }

    return n;
}
