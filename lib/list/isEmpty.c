
#include <stdlib.h>
#include "lib/list.h"

int list_isEmpty( struct List *l )
{
    return l->l_head->n_succ==NULL ? 1 : 0;
}
