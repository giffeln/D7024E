package main

import "container/list"

const bucketSize = 20

type routingTable struct {
	kadId   kademliaId
	buckets [idLength * 8]*list.List
}

func newRT(node kademliaId) (ret routingTable) {
	for i := 0; i < idLength*8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.kadId = node
	return
}
