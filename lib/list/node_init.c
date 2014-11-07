/**
 * Handle Node's in a List
 */

#include <stdlib.h>
#include "lib/list.h"

void node_init(struct Node *node) {
    node->n_pred = NULL;
    node->n_succ = NULL;
    node->name = NULL;
}

