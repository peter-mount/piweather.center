/*
 * Generates a simple ephemeris for the current date on the console
 * 
 * Initially it's just testing the code
 */


#include <stdlib.h>
#include <stdio.h>
#include <sys/types.h>
#include <math.h>
#include "time.h"
#include "astro/time.h"
#include "astro/sun.h"
#include "lib/charbuffer.h"
#include "lib/table.h"
#include "lib/config.h"
#include "location.h"

static void _table_add_riseset(TABLE_ROW *r, RISE_SET rs) {
    if (rs.type == RISE_SET_NEVER_RISES)
        table_addCellCenter(r, "Never Rises")->span = 2;
    else if (rs.type == RISE_SET_NEVER_SETS)
        table_addCellCenter(r, "Never Sets")->span = 2;
    else {
        table_add_hm(r, rs.rise);
        table_add_hm(r, rs.set);
    }
    table_add_hm(r, rs.length);
}

static void table_add_riseset(TABLE *t, char *s, RISE_SET today, RISE_SET tomorrow) {
    TABLE_ROW *r = table_newRow(t);
    table_addCellRight(r, s);
    _table_add_riseset(r, today);
    _table_add_riseset(r, tomorrow);
}

static void table_add_tm_date(TABLE_ROW *r, struct tm *tm) {
    table_addCellRight(r, "%d-%d-%d", tm->tm_year + 1900, tm->tm_mon + 1, tm->tm_mday);
}

static void table_add_tm_time(TABLE_ROW *r, struct tm *tm) {
    table_addCellRight(r, "%02d:%02d:%02d", tm->tm_hour, tm->tm_min, tm->tm_sec);
    table_addCell(r, (char *) tm->tm_zone);
}

int main(int argc, char **argv) {
    struct charbuffer b;
    charbuffer_init(&b);

    config_parse_dir("/etc/weather");
    astro_init();

    time_t now;
    time(&now);

    double JD = astro_julday_time(&now);
    double JD0 = astro_julian_0h(JD);

    double GST0 = astro_siderial_greenwich_0h(JD);
    double GST = astro_siderial_greenwich(JD);
    double lha = observatory.longitude / 15.0;
    double LST0 = GST0 + lha;
    double LST = GST + lha;

    SOLAR_EPHEMERIS today, tomorrow;
    astro_sunriseset(JD0, &observatory, &today);
    astro_sunriseset(JD0 + 1.0, &observatory, &tomorrow);

    struct tm tm_local, tm_jd, tm_jd0;
    localtime_r(&now, &tm_local);
    astro_calday(&tm_jd, JD);
    astro_calday(&tm_jd0, JD0);

    TABLE *t;
    TABLE_ROW *r;

    // Date table
    t = table_create();
    r = table_newRow(t);
    table_addCellCenter(r, "Station/Observatory Ephemeris")->span = 5;

    r = table_newRow(t);
    r = table_newRow(t);
    table_addCellRight(r, "Station Name");
    table_addCell(r, observatory.name)->span = 4;

    r = table_newRow(t);
    table_addCellRight(r, "Latitude");
    table_add_dms(r, observatory.latitude, 'N', 'S')->span = 2;
    table_addCellRight(r, "%.6f", observatory.latitude);

    r = table_newRow(t);
    table_addCellRight(r, "Longitude");
    table_add_dms(r, observatory.longitude, 'E', 'W')->span = 2;
    table_addCellRight(r, "%.6f", observatory.longitude);

    r = table_newRow(t);
    table_addCellRight(r, "Altitude");
    table_addCellRight(r, "%dm", (int) observatory.altitude)->span = 2;

    r = table_newRow(t);
    r = table_newRow(t);
    table_blankCell(r)->span = 2;
    table_addCellCenter(r, "Date");
    table_addCellCenter(r, "Time")->span = 2;
    table_addCellCenter(r, "DoW");
    table_addCellCenter(r, "DoY");

    r = table_newRow(t);
    table_addCellRight(r, "Local Time");
    table_blankCell(r);
    table_add_tm_date(r, &tm_local);
    table_add_tm_time(r, &tm_local);
    table_addCellRight(r, "%d", tm_local.tm_wday);
    table_addCellRight(r, "%d", tm_local.tm_yday + 1);

    r = table_newRow(t);
    table_addCellRight(r, "Julian Day");
    table_blankCell(r);
    table_addCellRight(r, "%.6f", JD);
    table_add_tm_time(r, &tm_jd);

    r = table_newRow(t);
    table_blankCell(r);
    table_addCellRight(r, "0h");
    table_addCellRight(r, "%.6f", JD0);
    table_add_tm_time(r, &tm_jd0);

    r = table_newRow(t);
    table_format(t);
    table_append(&b, t);
    table_destroy(t);

    // Time table
    t = table_create();
    r = table_newRow(t);
    table_addCellCenter(r, "Siderial Time")->span = 7;

    r = table_newRow(t);
    table_blankCell(r);
    table_addCellCenter(r, "Now");
    table_addCellCenter(r, "0h");

    r = table_newRow(t);
    table_addCellRight(r, "Greenwich");
    table_add_hms(r, GST);
    table_add_hms(r, GST0);

    r = table_newRow(t);
    table_addCellRight(r, "Local");
    table_add_hms(r, LST);
    table_add_hms(r, LST0);

    r = table_newRow(t);
    table_format(t);
    table_append(&b, t);
    table_destroy(t);

    // Sun data table
    t = table_create();

    r = table_newRow(t);
    table_addCellCenter(r, "Solar Ephemeris")->span = 7;
    r = table_newRow(t);

    r = table_newRow(t);
    table_blankCell(r);
    table_addCellCenter(r, "Today")->span = 3;
    table_addCellCenter(r, "Tomorrow")->span = 3;

    r = table_newRow(t);
    table_blankCell(r);
    table_addCellCenter(r, "Rise");
    table_addCellCenter(r, "Set");
    table_addCellCenter(r, "Len");
    table_addCellCenter(r, "Rise");
    table_addCellCenter(r, "Set");
    table_addCellCenter(r, "Len");

    table_add_riseset(t, "Daytime", today.day, tomorrow.day);
    table_add_riseset(t, "Civil Twilight", today.civil, tomorrow.civil);
    table_add_riseset(t, "Nautical Twilight", today.nautical, tomorrow.nautical);
    table_add_riseset(t, "Astronomical Twilight", today.astronomical, tomorrow.astronomical);

    table_format(t);
    table_append(&b, t);
    table_destroy(t);

    int len;
    char *p = charbuffer_tostring(&b, &len);
    printf(p);
    free(p);
    charbuffer_free(&b);
}
