
char *findEndOfLine(char *p) {
    while (*p && *p != '\n' && *p != '\r')
        p++;

    return p;
}
