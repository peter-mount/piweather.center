

#include <stdlib.h>
#include "lib/list.h"
#include "lib/history.h"

/**
 * Frees a History Node.
 * 
 * Note: if History->free is not null it is called with this node to
 * free any resources.
 * 
 * node_free() is then called regardless of free being present to free
 * the HistoryNode and HistoryNode->node.name if set.
 * 
 * @param h History
 * @param n HistoryNode
 */
void history_free(struct History *h, struct HistoryNode *n) {
    // Optional hook
    if (h->free)
        h->free(n);

    // Free HistoryNode and the node's name if present.
    node_free(&n->node);
}
