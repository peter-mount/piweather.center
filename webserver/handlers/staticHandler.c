#include <microhttpd.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
#include <unistd.h>
#include <sys/stat.h>
#include "lib/config.h"
#include "lib/string.h"
#include "webserver/webserver.h"

/* 32k page size */
#define PAGE_SIZE 327667
static const char *BASE = "/var/www";

// Don't allow any file begining with . to prevent people trying to scan filesystem
static const char *INVALID = "/.";

/**
 * Callback from microhttpd to read a block from the file
 * 
 * @param cls The FILE to read
 * @param pos The position in the file
 * @param buf Buffer to read into
 * @param max The size of buf
 * @return The number of bytes read into buf
 */
static ssize_t file_reader(void *cls, uint64_t pos, char *buf, size_t max) {
    FILE *file = cls;

    (void) fseek(file, pos, SEEK_SET);
    return fread(buf, 1, max, file);
}

/**
 * Callback from microhttpd when it's finished with the file. Here we close it.
 * 
 * @param cls FILE to close
 */
static void free_callback(void *cls) {
    FILE *file = cls;
    fclose(file);
}

/**
 * Handles static content under /var/www. This handles checks for attempts by the client to gain access outside of
 * that directory, then hands the file to microhttpd to stream it back.
 * 
 * @param connection
 * @return 
 */
int staticHandler(struct MHD_Connection * connection, const char *url) {
    char *p;
    int l;
    struct stat buf;
    struct MHD_Response *response;
    FILE *file;

    // Validate url - i.e. must start with /
    if (url[0] != '/')
        return MHD_NO;

    // Validate url - cannot contain /. (also mean /.. as well)
    // This prevents people from scanning outside of /var/www using ../../ style url's
    p = findString((char *) url, (char *) INVALID);
    if (p)
        return MHD_NO;

    // Form the file path
    l = strlen(url) + strlen(BASE) + 1;
    p = (char *) malloc(l);
    if (!p)
        return MHD_NO;
    strcpy(p, BASE);
    strcat(p, url);

    // Do we have access?
    if (stat(p, &buf)) {
        free(p);
        return MHD_NO;
    }

    // Open the file, although the stat above should pass we could still fail here
    file = fopen(p, "rb");
    if (!file) {
        free(p);
        return MHD_NO;
    }

    // Hand the file to microhttpd to stream back to the client
    response = MHD_create_response_from_callback(buf.st_size, PAGE_SIZE, &file_reader, file, &free_callback);
    if (response == NULL) {
        free(p);
        fclose(file);
        return MHD_NO;
    }

    int ret = MHD_queue_response(connection, MHD_HTTP_OK, response);
    MHD_destroy_response(response);

    // Cleanup
    free(p);
    return ret;
}
