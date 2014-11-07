/* 
 * File:   bytebuffer.h
 * Author: peter
 *
 * Created on February 7, 2014, 4:52 PM
 */

#ifndef BYTEBUFFER_H
#define	BYTEBUFFER_H

#ifdef	__cplusplus
extern "C" {
#endif

#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>

#define BYTEBUFFER_OK           0
#define BYTEBUFFER_ERROR        1

#define BYTEBUFFER_INITIAL_SIZE 10240

    struct bytebuffer {
        // The current position in the buffer where new data will be appended
        int pos;
        // The actual size of the buffer in memory, not it's contents
        int size;
        // The actual memory being used by the buffer
        void *buffer;
        // mutex used to ensure the buffer is concurrent
        pthread_mutex_t mutex;
    };

    extern int bytebuffer_init(struct bytebuffer *b);
    extern void bytebuffer_reset(struct bytebuffer *b);
    extern int bytebuffer_put(struct bytebuffer *b, void *src, int len);
    extern int bytebuffer_size(struct bytebuffer *b);
    extern void *bytebuffer_toarray(struct bytebuffer *b, int *len);
    extern void bytebuffer_write(struct bytebuffer *b, FILE *out);
    extern void bytebuffer_read(struct bytebuffer *b, FILE *in);
    extern void bytebuffer_free(struct bytebuffer *b);

#ifdef	__cplusplus
}
#endif

#endif	/* BYTEBUFFER_H */

