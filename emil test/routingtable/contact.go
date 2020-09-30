package routingtable

import(
	"sort"
	"../node"
)

type contact struct {
	nodeId *kademliaId
	address  string
	distance *kademliaId
}

