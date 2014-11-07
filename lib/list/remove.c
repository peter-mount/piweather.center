
#include <stdlib.h>
#include "lib/list.h"

/**
 * Removes a node from a List
 * @param n Node to remove
 * @return the node that was after this one before removal
 */
struct Node *list_remove(struct Node *n) {
    struct Node *next = n->n_succ;
    n->n_succ->n_pred = n->n_pred;
    n->n_pred->n_succ = n->n_succ;
    return next;
}
