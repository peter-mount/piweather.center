
#include <stdlib.h>
#include <time.h>
#include "lib/list.h"
#include "lib/history.h"

/**
 * Clears a History of all entries
 * 
 * @param h History to expire
 */
void history_clear(struct History *h) {
    while (!list_isEmpty(&h->list))
        history_free(h, (struct HistoryNode *) list_removeHead(&h->list));
}
