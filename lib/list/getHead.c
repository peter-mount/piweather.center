
#include <stdlib.h>
#include "lib/list.h"

struct Node *list_getHead(struct List *l) {
    return l->l_head->n_succ == NULL ? NULL : l->l_head;
}
