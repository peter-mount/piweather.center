
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <errno.h>

/**
 * mkdirs implementation as there's no c version
 * 
 * http://stackoverflow.com/questions/2336242/recursive-mkdir-system-call-on-unix
 * 
 * @param path
 * @param nmode
 * @return 0 if successful, 1 on failure
 */
int mkdirs(char *path, mode_t nmode) {
    int oumask;
    struct stat sb;

    if (stat(path, &sb) == 0) {
        if (S_ISDIR(sb.st_mode) == 0) {
            fprintf(stderr, "%s: file exists but is not a directory", path);
            return 1;
        }

        if (chmod(path, nmode)) {
            fprintf(stderr, "%s: %s", path, strerror(errno));
            return 1;
        }

        return 0;
    }

    oumask = umask(0);
    char *npath = strdup(path);

    char *p = npath;

    // Skip leading slashes.
    while (*p == '/')
        p++;

    while (p = strchr(p, '/')) {
        *p = '\0';
        if (stat(npath, &sb) != 0) {
            if (mkdir(npath, nmode)) {
                fprintf(stderr, "cannot create directory %s: %s", npath, strerror(errno));
                umask(oumask);
                free(npath);
                return 1;
            }
        } else if (S_ISDIR(sb.st_mode) == 0) {
            fprintf(stderr, "%s: file exists but is not a directory", npath);
            umask(oumask);
            free(npath);
            return 1;
        }

        *p++ = '/'; /* restore slash */
        while (*p == '/')
            p++;
    }

    /* Create the final directory component. */
    if (stat(npath, &sb) && mkdir(npath, nmode)) {
        fprintf(stderr, "cannot create directory `%s': %s", npath, strerror(errno));
        umask(oumask);
        free(npath);
        return 1;
    }

    umask(oumask);
    free(npath);
    return 0;
}
