
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <curl/curl.h>
#include "lib/charbuffer.h"

/**
 * Used to capture the body of the response.
 * 
 * To use this method with libcurl:
 * 
 *    struct charbuffer *buffer;
 *    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, charbuffer_curl_write);
 *    curl_easy_setopt(curl, CURLOPT_WRITEDATA, buffer);
 *    charbuffer_reset(buffer);
 * 
 * Note: If you do not reset before each request then you'll append to any existing content.
 * 
 * @param b Pointer to received buffer
 * @param s size of block
 * @param n number of blocks in buffer
 * @param stream pointer to charbuffer
 * @return s*n always
 */
size_t charbuffer_curl_write(void *b, size_t s, size_t n, void *stream) {
    struct charbuffer *l = (struct charbuffer *) stream;
    charbuffer_put(l, (char *) b, s * n);
    return s*n;
}

/**
 * By default curl will write the response body to the console, so using this
 * prevents that by ignoring the response
 * 
 * @param b Pointer to received buffer
 * @param s size of block
 * @param n number of blocks in buffer
 * @param stream pointer to charbuffer
 * @return s*n always
 */
size_t charbuffer_curl_write_sink(void *b, size_t s, size_t n, void *stream) {
    return s*n;
}
