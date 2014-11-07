
#include <stdlib.h>
#include "lib/list.h"

struct Node *list_getTail(struct List *l) {
    return l->l_head->n_succ == NULL ? NULL : l->l_tailpred;
}
