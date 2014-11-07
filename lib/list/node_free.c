/**
 * Handle Node's in a List
 */

#include <stdlib.h>
#include "lib/list.h"

/**
 * Free's a node.
 * 
 * Note: If this is at the start of another structure then that structure will be freed.
 * So you must free any other resources first.
 * 
 * If the node has a name then the name is also freed.
 * 
 * This also does nothing if n is not a node, i.e. NULL or the list Header
 * 
 * @param n Node to free
 */
void node_free(struct Node *n) {
    if (n->name)
        free(n->name);
    free(n);
}

