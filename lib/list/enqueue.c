
#include <stdlib.h>
#include "lib/list.h"

void list_enqueue(struct List *l, struct Node *n) {
    // Empty list so just add it
    if (list_isEmpty(l)) {
        list_addHead(l, n);
        return;
    }

    // Find the first node with lower priority to ours.
    // So, if we have nodes of the same priority we skip it as we
    // enqueue after it
    struct Node *c = l->l_head;
    while (list_isNode(c) && c->pri >= n->pri)
        c = c->n_succ;

    if (list_isNode(c)) {
        // If it's the head then just add to the head
        if (list_isHead(c))
            list_addHead(l, n);
        else {
            // insert before the found node
            n->n_pred = c->n_pred;
            n->n_succ = c;
            c->n_pred = n;
            n->n_pred->n_succ = n;
        }
    } else
        // Add to the tail
        list_addTail(l, n);
}
