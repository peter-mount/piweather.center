

#include <stdlib.h>
#include "lib/config.h"
#include "lib/list.h"

/**
 * Finds the NamedNode in a list by name.
 * 
 * @param l List
 * @param name Name to find
 * @return NamedNode or NULL if not found
 */
struct Node *list_findNode(struct List *l, const char *name) {

    struct Node *n = list_getHead(l);
    while (list_isNode(n)) {
        if (strcmp(n->name, name) == 0)
            return n;
        n = n->n_succ;
    }

    return NULL;
}
