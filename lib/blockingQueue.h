/* 
 * File:   queue.h
 * Author: peter
 *
 * Created on March 4, 2014, 6:39 PM
 */

#ifndef BLOCKINGQUEUE_H
#define	BLOCKINGQUEUE_H

#include "lib/list.h"

struct blocking_queue {
    struct List list;
    pthread_mutex_t mutex;
    pthread_cond_t condition;
};

/**
 * Initialise a queue
 * @param queue 
 */
extern void blockingqueue_init(struct blocking_queue *queue);
/**
 * Get the next element from the queue, blocking until an element is available
 * @param queue
 * @return 
 */
extern struct Node *blockingqueue_get(struct blocking_queue *queue);
/**
 * Add a new element to the queue, notifying any listeners
 * @param queue
 * @param element
 */
extern void blockingqueue_add(struct blocking_queue *queue, struct Node *node);

#endif	/* BLOCKINGQUEUE_H */

