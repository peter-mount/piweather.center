
#include <stdlib.h>
#include "lib/list.h"

int list_isHead(struct Node *n) {
    return list_isNode(n) && !list_isNode(n->n_pred);
}
