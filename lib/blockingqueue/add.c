/**
 * A concurrent linked queue
 */

#include <pthread.h>
#include "lib/blockingQueue.h"
#include "lib/list.h"

void blockingqueue_add(struct blocking_queue *queue, struct Node *node) {
    pthread_mutex_lock(&queue->mutex);

    list_addTail(&queue->list, node);

    pthread_mutex_unlock(&queue->mutex);

    // Signal waiting threads
    pthread_cond_signal(&queue->condition);
}
