/*
 * Copyright (C) 2007 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * hash and equals functions for string key's
 */

#include <stdbool.h>
#include <stdint.h>

bool hashmapStringEquals(void *key_a, void *key_b)
{
    return !strcmp((const char *)key_a, (const char *)key_b);
}

/* use djb hash unless we find it inadequate */
int hashmapStringHash(void *str)
{
    uint32_t hash = 5381;
    char *p;

    for (p = str; p && *p; p++)
        hash = ((hash << 5) + hash) + *p;
    return (int)hash;
}
