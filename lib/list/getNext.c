
#include <stdlib.h>
#include "lib/list.h"

struct Node *list_getNext(struct Node *n) {
    if (list_isNode(n) && list_isNode(n->n_succ))
        return n->n_succ;
    else
        return NULL;
}
