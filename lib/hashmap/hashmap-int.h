
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
 * 
 * File:   hashmap-int.h
 * Author: Peter T Mount
 *
 * Created on April 7, 2014, 10:28 AM
 */

#ifndef HASHMAP_INT_H
#define	HASHMAP_INT_H

#include "lib/hashmap.h"
#include <assert.h>
#include <errno.h>
#include <pthread.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <sys/types.h>

typedef struct Entry Entry;

struct Entry {
    void* key;
    int hash;
    void* value;
    Entry* next;
};

struct Hashmap {
    Entry** buckets;
    size_t bucketCount;
    int (*hash)(void* key);
    bool (*equals)(void* keyA, void* keyB);
    pthread_mutex_t lock;
    size_t size;
};

extern size_t hashmapCalculateIndex(size_t bucketCount, int hash);
extern Entry* hashmapCreateEntry(void* key, int hash, void* value);
extern void hashmapExpandIfNecessary(Hashmap* map);
extern bool hashmapEqualKeys(void* keyA, int hashA, void* keyB, int hashB, bool (*equals)(void*, void*));
extern int hashmapHashKey(Hashmap* map, void* key);

#endif	/* HASHMAP_INT_H */

