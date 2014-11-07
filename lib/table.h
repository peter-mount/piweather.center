/* 
 * File:   table.h
 * Author: Peter T Mount
 *
 * Created on April 29, 2014, 5:53 PM
 */

#ifndef TABLE_H
#define	TABLE_H

#include "lib/charbuffer.h"
#include "lib/list.h"

typedef enum {
    TABLE_ALIGN_LEFT,
    TABLE_ALIGN_CENTER,
    TABLE_ALIGN_RIGHT
} TABLE_ALIGN_T;

typedef struct {
    struct Node node;
    TABLE_ALIGN_T align;
    int span;
    int width;
} TABLE_CELL;

typedef struct {
    struct Node node;
    struct List cells;
} TABLE_ROW;

typedef struct {
    struct List rows;
    // Maximum number of cells per row
    int max_cells;
    // Array of cell widths
    int *width;
} TABLE;

extern void table_append(struct charbuffer *b, TABLE *t);
extern TABLE *table_create();
extern void table_destroy(TABLE *t);
extern TABLE_ROW *table_newRow(TABLE *t);
extern TABLE_CELL *table_blankCell(TABLE_ROW *r);
extern TABLE_CELL *table_addCell(TABLE_ROW *r, char *fmt, ...);
extern TABLE_CELL *table_addCellCenter(TABLE_ROW *r, char *fmt, ...);
extern TABLE_CELL *table_addCellRight(TABLE_ROW *r, char *fmt, ...);
extern TABLE_CELL *table_add_hm(TABLE_ROW *r, double t);
extern TABLE_CELL *table_add_hms(TABLE_ROW *r, double t);
extern TABLE_CELL *table_add_dms(TABLE_ROW *r, double t, char p, char n);

#endif	/* TABLE_H */

