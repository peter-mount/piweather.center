/**
 * Handle Node's in a List
 */

#include <stdlib.h>
#include <string.h>
#include "lib/list.h"

/**
 * Convenience method for simple lists.
 * 
 * Allocates a new Node.
 * 
 * Note: If this node is freed, then node_free() must be used. However name will
 * also be freed, so it should be allocated via malloc() or strdup().
 */
void node_alloc(char *name) {
    struct Node *node = (struct Node *) malloc(sizeof (struct Node));
    memset(node, 0, sizeof (struct Node));
    if (node) {
        node->n_pred = NULL;
        node->n_succ = NULL;
        node->name = name;
    }
}

