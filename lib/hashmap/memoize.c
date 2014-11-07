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

#include "hashmap-int.h"

void* hashmapMemoize(Hashmap* map, void* key, void* (*initialValue)(void* key, void* context), void* context) {
    int hash = hashmapHashKey(map, key);
    size_t index = hashmapCalculateIndex(map->bucketCount, hash);

    Entry** p = &(map->buckets[index]);
    while (true) {
        Entry* current = *p;

        // Add a new entry.
        if (current == NULL) {
            *p = hashmapCreateEntry(key, hash, NULL);
            if (*p == NULL) {
                errno = ENOMEM;
                return NULL;
            }
            void* value = initialValue(key, context);
            (*p)->value = value;
            map->size++;
            hashmapExpandIfNecessary(map);
            return value;
        }

        // Return existing value.
        if (hashmapEqualKeys(current->key, current->hash, key, hash, map->equals)) {
            return current->value;
        }

        // Move to next entry.
        p = &current->next;
    }
}
