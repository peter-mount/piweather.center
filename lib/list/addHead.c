
#include <stdlib.h>
#include "lib/list.h"

void list_addHead(struct List *l, struct Node *n) {
    n->n_succ = l->l_head;
    n->n_pred = (struct Node *) &l->l_head;
    l->l_head->n_pred = n;
    l->l_head = n;
}
