#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include "scheduler.h"
#include "lib/config.h"
#include "astro/location.h"

static int filter(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_above_horizon(&schedule.today.civil, m);
}

static void *test(void *arg) {
    time_t now;
    struct tm timeinfo;
    time(&now);
    localtime_r(&now, &timeinfo);

    printf("%04d:%02d:%02d %02d:%02d:%02d Ping!\n",
            timeinfo.tm_year + 1900,
            timeinfo.tm_mon + 1,
            timeinfo.tm_mday,
            timeinfo.tm_hour,
            timeinfo.tm_min,
            timeinfo.tm_sec
            );
}

int main(int argc, char **argv) {
    int i, h, m, o;

    config_parse_dir("/etc/weather");
    astro_init();
    scheduler_init();

    test(NULL);

    for (i = 1; i < argc; i++) {
        printf("\nTest: %s\n", argv[i]);
        SCHEDULE_ENTRY *e = scheduler_new(test, NULL);
        scheduler_parse(e, argv[i]);
        //e->filter = filter;

        printf("   ");
        for (m = 0; m < 60; m++, o++)
            printf(" %02d", m);
        printf("\n");

        for (h = 0; h < 24; h++) {
            printf("%02d ", h);
            o = h * 60;
            for (m = 0; m < 60; m++, o++)
                printf(" %2d", scheduler_trigger(e, o));
            //printf(" %2d", scheduler_getBit(e, o));
            printf("\n");
        }

        scheduler_add(e);
    }

    printf("Testing scheduler...\n");
    sleep(300);
}
