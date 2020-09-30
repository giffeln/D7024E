package routingtable

import (
	"container/list"
	"../node"
	"time"
)

const bucketsize = 32

type bucket struct {
	*list.List
	AccessTime time.Time
}

type table struct {
	buckets     [node.IDLength]*bucket
	me          Contact
	refreshTime time.Duration
}

type distance [node.IDByteLength]byte

func calcDistance (a, b node.ID) ({
	return a.ID ^ b.ID 
}