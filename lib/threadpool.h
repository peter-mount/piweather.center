/* 
 * File:   threadpool.h
 * Author: Peter T Mount
 *
 * Created on March 26, 2014, 3:05 PM
 */

#ifndef THREADPOOL_H
#define	THREADPOOL_H

#include "lib/blockingQueue.h"
#include "lib/list.h"

struct thread_pool {
    // The thread pool
    int pool_size;
    pthread_t *threads;
    // Queue used by the threads
    struct blocking_queue queue;
};

struct thread_worker {
    struct thread_pool *pool;
    void *(*worker) (struct Node *, void *);
    void *data;
};

extern void threadpool_init(struct thread_pool *pool, int size);
extern void threadpool_start(struct thread_pool *pool, void *(*poolWorker) (struct Node *, void *), void *data);
extern void threadpool_submit(struct thread_pool *pool, struct Node *job);

#endif	/* THREADPOOL_H */

