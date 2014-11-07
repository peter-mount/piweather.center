/* 
 * File:   charbuffer.h
 * Author: peter
 *
 * Created on February 11, 2014, 4:49 PM
 */

#ifndef CHARBUFFER_H
#define	CHARBUFFER_H

#include <pthread.h>
#include <stdio.h>
#include <stdarg.h>

#include "global_config.h"

#define CHARBUFFER_OK           0
#define CHARBUFFER_ERROR        1

#define CHARBUFFER_INITIAL_SIZE 1024

struct charbuffer {
    // The current position in the buffer where new data will be appended
    int pos;
    // The actual size of the buffer in memory, not it's contents
    int size;
    // The actual memory being used by the buffer
    char *buffer;
    // mutex used to ensure the buffer is concurrent
    pthread_mutex_t mutex;
};

extern int charbuffer_add(struct charbuffer *b, char c);
extern int charbuffer_append(struct charbuffer *b, char *src);
extern void charbuffer_appendbuffer(struct charbuffer *dest, struct charbuffer *src);
extern void charbuffer_free(struct charbuffer *b);
extern int charbuffer_init(struct charbuffer *b);
extern int charbuffer_printf(struct charbuffer *b, char *fmt, ...);
extern int charbuffer_put(struct charbuffer *b, char *src, int len);
extern void charbuffer_read(struct charbuffer *b, FILE *in);
extern void charbuffer_reset(struct charbuffer *b);
extern int charbuffer_size(struct charbuffer *b);
extern void *charbuffer_toarray(struct charbuffer *b, int *len);
extern char *charbuffer_tostring(struct charbuffer *b, int *len);

extern void charbuffer_append_int(struct charbuffer *b, int v, int width);
extern void charbuffer_time_hm(struct charbuffer *b, double t);
extern void charbuffer_time_hms(struct charbuffer *b, double t);
extern int charbuffer_append_padleft(struct charbuffer *b, char *src, int width);
extern int charbuffer_append_padright(struct charbuffer *b, char *src, int width);
extern int charbuffer_append_center(struct charbuffer *b, char *src, int width);

// JSON
extern void charbuffer_reset_json(struct charbuffer *b);
extern void charbuffer_append_json(struct charbuffer *b, char *n, char *fmt, ...);
extern void charbuffer_end_json(struct charbuffer *b);

// XML
extern void charbuffer_reset_xml(struct charbuffer *b, char *tag);
extern void charbuffer_append_xml(struct charbuffer *b, char *n, char *fmt, ...);
extern void charbuffer_end_xml(struct charbuffer *b, char *tag);

// Curl
extern int charbuffer_append_urlencode(struct charbuffer *b, char *src);
extern int charbuffer_append_form_field(struct charbuffer *b, char *name, char *value);
extern int charbuffer_append_form_fieldf(struct charbuffer *b, char *name, char *fmt, ...);
#ifdef HAVE_CURL
extern size_t charbuffer_curl_write(void *b, size_t s, size_t n, void *stream);
extern size_t charbuffer_curl_write_sink(void *b, size_t s, size_t n, void *stream);
#endif

#endif	/* CHARBUFFER_H */

