package routingtable

import(
	"sort"
)

type contact struct {
	nodeId node.ID
	address  string
	distance distance
}

