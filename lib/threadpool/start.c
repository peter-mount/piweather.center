
#include <stdlib.h>
#include <pthread.h>
#include <string.h>
#include "lib/blockingQueue.h"
#include "lib/threadpool.h"

static void *worker(void *arg) {
    struct thread_worker *worker = (struct thread_worker *) arg;
    while (1) {
        struct Node *n = blockingqueue_get(&worker->pool->queue);
        if (n) {
            worker->worker(n, worker->data);
        }
    }
}

/**
 * Start's the thread pool with the provided worker thread
 * 
 * @param pool thread_pool
 * @param worker worker
 * @param data argument passed to worker
 */
void threadpool_start(struct thread_pool *pool, void *(*poolWorker) (struct Node *, void *), void *data) {
    struct thread_worker *w = (struct thread_worker *) malloc(sizeof (struct thread_worker));
    memset(w, 0, sizeof (struct thread_worker));
    w->pool = pool;
    w->worker = poolWorker;
    w->data = data;

    int i = 0;
    pool->threads = (pthread_t *) calloc(pool->pool_size, sizeof (pthread_t *));
    for (i = 0; i < pool->pool_size; i++)
        pthread_create(&(pool->threads[i]), NULL, worker, (void *) w);
}
