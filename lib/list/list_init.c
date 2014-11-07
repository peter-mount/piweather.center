
#include <stdlib.h>
#include "lib/list.h"

/**
 * Initialise a List
 * @param list List to initialise
 */
void list_init(struct List *list) {
    list->l_head = (struct Node *) &list->l_tail;
    list->l_tail = NULL;
    list->l_tailpred = (struct Node *) &list->l_head;
}
