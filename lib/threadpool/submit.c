
#include <stdlib.h>
#include <pthread.h>
#include "lib/blockingQueue.h"
#include "lib/threadpool.h"

void threadpool_submit(struct thread_pool *pool, struct Node *job) {
    blockingqueue_add(&pool->queue, job);
}
