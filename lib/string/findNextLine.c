
char *findNextLine(char *p) {
    while (*p && *p != '\n')
        p++;
    if (*p == '\n')
        p++;

    return p;
}
