/**
 * Set the thread priority
 */

#include <sched.h>
#include <string.h>
/**
 * Attempt to set the priority of the running program/thread
 * @param priority new priority
 * @return 
 */
int thread_setPriority(const int priority) {
    struct sched_param sched;
    memset(&sched, 0, sizeof (sched));

    int maxPriority = sched_get_priority_max(SCHED_RR);
    if (priority > maxPriority)
        sched.sched_priority = maxPriority;
    else
        sched.sched_priority = priority;

    return sched_setscheduler(0, SCHED_RR, &sched);
}