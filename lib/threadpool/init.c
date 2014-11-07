
#include <stdlib.h>
#include <pthread.h>
#include "lib/blockingQueue.h"
#include "lib/threadpool.h"

void threadpool_init(struct thread_pool *pool, int size) {
    pool->pool_size = size;
    blockingqueue_init(&pool->queue);
}