
#include <stdlib.h>
#include "lib/list.h"

void list_addTail(struct List *l, struct Node *n) {
    n->n_succ = (struct Node *) &l->l_tail;
    n->n_pred = l->l_tailpred;
    l->l_tailpred->n_succ = n;
    l->l_tailpred = n;
}
