package main

import (
	"container/list"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sort"
)

const IdLength = 20

type NodeID [IdLength]byte

func NewNodeID(data string) (ret NodeID) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}
	return
}

func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}
	return
}

func (node NodeID) String() string {
	return hex.EncodeToString(node[0:IdLength])
}

func (node NodeID) Equals(other NodeID) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other.(NodeID)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

func (node NodeID) Xor(other NodeID) (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = node[i] ^ other[i]
	}
	return
}

func (node NodeID) PrefixLen() (ret int) {
	for i := 0; i < IdLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return IdLength*8 - 1
}

const BucketSize = 20

type Contact struct {
	id NodeID
}

type RoutingTable struct {
	node    *Contact
	buckets [IdLength * 8]*list.List
}

func NewRoutingTable(node *Contact) (ret RoutingTable) {
	for i := 0; i < IdLength*8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.node = node
	return
}

func (table *RoutingTable) Update(contact *Contact) {
	prefix_length := contact.id.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[prefix_length]
	element := iterable.Find(bucket, func(x interface{}) bool {
		return x.(*Contact).id.Equals(table.node.id)
	})
	if element == nil {
		if bucket.Len() <= BucketSize {
			bucket.PushFront(contact)
		}
		// TODO: Handle insertion when the list is full by evicting old elements if
		// they don't respond to a ping.
	} else {
		bucket.MoveToFront(element.(*list.Element))
	}
}

type ContactRecord struct {
	node    *Contact
	sortKey NodeID
}

func (rec *ContactRecord) Less(other interface{}) bool {
	return rec.sortKey.Less(other.(*ContactRecord).sortKey)
}

func copyToVector(start, end *list.Element, vec *vector.Vector, target NodeID) {
	for elt := start; elt != end; elt = elt.Next() {
		contact := elt.Value.(*Contact)
		vec.Push(&ContactRecord{contact, contact.id.Xor(target)})
	}
}

func (table *RoutingTable) FindClosest(target NodeID, count int) (ret *vector.Vector) {
	ret = new(vector.Vector).Resize(0, count)

	bucket_num := target.Xor(table.node.id).PrefixLen()
	bucket := table.buckets[bucket_num]
	copyToVector(bucket.Front(), nil, ret, target)

	for i := 1; (bucket_num-i >= 0 || bucket_num+i < IdLength*8) && ret.Len() < count; i++ {
		if bucket_num-i >= 0 {
			bucket = table.buckets[bucket_num-i]
			copyToVector(bucket.Front(), nil, ret, target)
		}
		if bucket_num+i < IdLength*8 {
			bucket = table.buckets[bucket_num+i]
			copyToVector(bucket.Front(), nil, ret, target)
		}
	}

	sort.Sort(ret)
	if ret.Len() > count {
		ret.Cut(count, ret.Len())
	}
	return
}

type Kademlia struct {
	routes    RoutingTable
	NetworkId string
}

func NewKademlia(self *Contact, networkId string) (ret *Kademlia) {
	ret = new(Kademlia)
	ret.routes = NewRoutingTable(self)
	ret.NetworkId = networkId
	return
}

type RPCHeader struct {
	Sender    *Contact
	NetworkId string
}

func (k *Kademlia) HandleRPC(request, response *RPCHeader) os.Error {
	if request.NetworkId != k.NetworkId {
		return os.NewError(fmt.Sprintf("Expected network %s, got %s", k.NetworkId, request.NetworkId))
	}
	if request.Sender != nil {
		k.routes.Update(request.Sender)
	}
	response.Sender = &k.routes.node
	return nil
}

type KademliaCore struct {
	kad *Kademlia
}

type PingRequest struct {
	RPCHeader
}

type PingResponse struct {
	RPCHeader
}

func (kc *KademliaCore) Ping(args *PingRequest, response *PingResponse) (err os.Error) {
	if err = kc.kad.HandleRPC(&args.RPCHeader, &response.RPCHeader); err == nil {
		log.Stderr("Ping from %s\n", args.RPCHeader)
	}
	return
}

func (k *Kademlia) Serve() (err os.Error) {
	rpc.Register(&KademliaCore{k})

	rpc.HandleHTTP()
	if i, err := net.Listen("udp", k.routes.node.address); err == nil {
		go http.Serve(1, nil)
	}
	return
}
