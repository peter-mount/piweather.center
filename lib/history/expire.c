
#include <stdlib.h>
#include <time.h>
#include "lib/list.h"
#include "lib/history.h"

/**
 * Expires any HistoryNode's who's time is older than max_age seconds ago.
 * 
 * Note: When an entry is re
 * @param h History to expire
 */
void history_expire(struct History *h) {
    // The time before which a HistoryNode is to be expired
    time_t retirement;
    time(&retirement);
    retirement = retirement - h->max_age;
    
    while (list_isNode(h->list.l_head) && ((struct HistoryNode *) h->list.l_head)->time < retirement)
        history_free(h, (struct HistoryNode *) list_removeHead(&h->list));
}
