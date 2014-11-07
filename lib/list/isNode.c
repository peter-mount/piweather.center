
#include <stdlib.h>
#include "lib/list.h"

/**
 * Is node really a node?
 * 
 * This can happen if you are traversing the list. When you end up back on the list then this returns 0
 * 
 * @param n
 * @return 
 */
int list_isNode(struct Node *n) {
    return (n && n->n_succ && n->n_pred) ? 1 : 0;
}
