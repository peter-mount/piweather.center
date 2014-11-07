/* 
 * File:   string.h
 * Author: Peter T Mount
 *
 * Created on March 26, 2014, 4:58 PM
 */

#ifndef STRING_H
#define	STRING_H

extern char* genurl(const char *contextPath, const char *suffix);
extern void fatalError(char *fmt, ...);
extern void fatalIfNull(void *v, char *fmt, ...);

extern char *extractString(char *start, char **end);
extern char *find(char *p, char c);
extern char *findEndOfLine(char *p);
extern char *findNextLine(char *p);
extern char *findNonWhitespace(char *p);
extern char *findString(char *p, char *s);
extern char *findWhitespace(char *p);

extern int strendswith(char *s, char *p);

#endif	/* STRING_H */

