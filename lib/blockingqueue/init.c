/**
 * A concurrent linked queue
 */

#include <pthread.h>
#include "lib/blockingQueue.h"
#include "lib/list.h"

void blockingqueue_init(struct blocking_queue *queue) {
    pthread_mutex_init(&queue->mutex, NULL);
    pthread_cond_init(&queue->condition, NULL);
    list_init(&queue->list);
}
