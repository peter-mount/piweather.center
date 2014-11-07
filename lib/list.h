/* 
 * File:   list.h
 * Author: Peter T Mount
 *
 * Created on March 26, 2014, 1:16 PM
 */

#ifndef LIST_H
#define	LIST_H

#include <stdlib.h>

#ifdef	__cplusplus
extern "C" {
#endif

    /**
     * A node within a list.
     */
    struct Node {
        struct Node *n_succ;
        struct Node *n_pred;
        char *name;
        int8_t pri;
        int8_t pad;
    };

    /**
     * A minimal Node. This can be used for all of the functions defined here
     * EXCEPT for node_free().
     */
    struct MinNode {
        struct Node *n_succ;
        struct Node *n_pred;
    };

    struct List {
        struct Node *l_head;
        struct Node *l_tail;
        struct Node *l_tailpred;
    };

    extern void list_init(struct List *list);
    extern void list_addHead(struct List *l, struct Node *n);
    extern void list_addTail(struct List *l, struct Node *n);
    extern struct Node *list_findNode(struct List *l, const char *name);
    extern struct Node *list_getHead(struct List *l);
    extern struct Node *list_getNext(struct Node *n);
    extern struct Node *list_getPred(struct Node *n);
    extern struct Node *list_getTail(struct List *l);
    extern void list_insert(struct List *l, struct Node *n, struct Node *after);
    extern int list_isEmpty(struct List *l);
    extern int list_isHead(struct Node *n);
    extern int list_isNode(struct Node *n);
    extern int list_isTail(struct Node *n);
    extern struct Node *list_remove(struct Node *n);
    extern struct Node *list_removeHead(struct List *l);
    extern struct Node *list_removeTail(struct List *l);
    extern int list_size(struct List *list);

    extern void node_alloc(char *name);
    extern void node_init(struct Node *node);
    extern void node_free(struct Node *n);
#ifdef	__cplusplus
}
#endif

#endif	/* LIST_H */

