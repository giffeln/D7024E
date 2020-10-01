package main

import (
	"fmt"
	"sort"
)

type contact struct {
	nodeId   kademliaId
	address  string
	distance kademliaId
}

type candidates struct {
	contacts []contact
}

func makeContact(id kademliaId, address string) contact {
	return contact{id, address, nil}
}

func (contact contact) setDistance(target kademliaId) {
	contact.distance = contact.nodeId.calcDist(target)
}

func (firstContact contact) smaller(secondContact contact) bool {
	return firstContact.distance.smaller(secondContact.distance)
}

func (contact contact) toString() string {
	return fmt.Sprintf(contact("%s", "%s"), contact.nodeId, contact.address)
}

func (candidates candidates) add(addCandidates []contact) {
	candidates.contacts = append(candidates.contacts, addCandidates...)
}

func (candidates candidates) getCandidates(amount int) []contact {
	return candidates.contacts[:amount]
}

func (candidates candidates) sort() {
	sort.Sort(candidates)
}

func (candidates candidates) len() int {
	return len(candidates.contacts)
}

func (candidates candidates) swap(a, b int) {
	candidates.contacts[a], candidates.contacts[b] = candidates.contacts[b], candidates.contacts[a]
}

func (candidates candidates) less(a, b int) {
	return candidates.contacts[a].smaller(&candidates.contacts[b])
}
