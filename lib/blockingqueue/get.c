/**
 * A concurrent linked queue
 */

#include <pthread.h>
#include "lib/blockingQueue.h"
#include "lib/list.h"

struct Node *blockingqueue_get(struct blocking_queue *queue) {
    pthread_mutex_lock(&queue->mutex);

    while (list_isEmpty(&queue->list))
        pthread_cond_wait(&queue->condition, &queue->mutex);

    // Remove the element
    struct Node *n = list_removeHead( &queue->list);

    pthread_mutex_unlock(&queue->mutex);

    return n;
}
