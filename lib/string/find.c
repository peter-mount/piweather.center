
char *find(char *p, char c) {
    while (*p && *p != c && *p != ' ' && *p != '\t')
        p++;
    return p;
}
