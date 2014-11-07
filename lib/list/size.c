
#include <stdlib.h>
#include "lib/list.h"

/**
 * Returns the size of the list.
 * 
 * It is better to use list_isEmpty if you want to know if the list is empty
 * 
 * @param list List
 * @return number of Node's in the list
 */
int list_size(struct List *list) {
    int size = 0;
    struct Node *n = list->l_head;
    while (n->n_succ) {
        size++;
        n = n->n_succ;
    }
    return size;
}
