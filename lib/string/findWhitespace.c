
char *findWhitespace(char *p) {
    while (*p && *p != ' ' && *p != '\t')
        p++;

    return p;
}
