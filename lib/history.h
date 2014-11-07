/* 
 * File:   history.h
 * Author: Peter T Mount
 *
 * Created on April 12, 2014, 1:30 PM
 */

#ifndef HISTORY_H
#define	HISTORY_H

#include <time.h>
#include "lib/list.h"

/**
 * A Node within a History.
 */
struct HistoryNode {
    // Node within List
    struct Node node;
    // The timestamp of this entry, used in expiry
    time_t time;
};

/**
 * A List who's entries are kept only for a specified amount of time.
 * 
 * This list is ordered as a FiFO queue, so new entries go to the lists tail
 * but entries are only expired when they are at the queue's head and have
 * expired.
 * 
 * free is an optional hook to free any resources in this node.
 * 
 * WARNING: if free is NULL or not NULL, node_free() will be called to free the
 * actual node - the hook is there only to free anything other than what's in
 * HistoryNode or HistoryNode->node.name
 */
struct History {
    // The historical data
    struct List list;
    // The amount of time to keep data
    time_t max_age;
    // Optional hook to be called when a HistoryNode is freed
    void (*free)(struct HistoryNode *n);
};

extern void history_add(struct History *h, struct HistoryNode *n);
extern void history_clear(struct History *h);
extern void history_expire(struct History *h);
extern void history_free(struct History *h, struct HistoryNode *n);
extern void history_init(struct History *h, time_t max_age);

#endif	/* HISTORY_H */

