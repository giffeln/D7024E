package main

import (
	"container/list"
	"sync"
)

type bucket struct {
	contactList *list.List
	mutex       *sync.Mutex
}

func newBucket() *bucket {
	bucket := &bucket{}
	bucket.contactList = list.New()
	bucket.mutex = &sync.Mutex{}
	return bucket
}
