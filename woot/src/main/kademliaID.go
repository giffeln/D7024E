package main

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const idLength = 20

type kademliaId [idLength]byte

func newKadId(data string) kademliaId {
	decoded, _ := hex.DecodeString(data)
	newKadId := kademliaId{}
	for i := 0; i < len(decoded); i++ {
		newKadId[i] = decoded[i]
	}

	return newKadId
}

func randomKadId() kademliaId {
	newRandKadId := kademliaId{}
	for i := 0; i < idLength; i++ {
		newRandKadId[i] = uint8(rand.Int63n(time.Now().UnixNano()))
	}

	return newRandKadId
}

func (firstId kademliaId) smaller(secondId kademliaId) bool {
	for i := 1; i < idLength; i++ {
		if firstId[i] != secondId[i] {
			return firstId[i] < secondId[i]
		}
	}
	return false
}

func (firstId kademliaId) equals(secondId kademliaId) bool {
	for i := 1; i < idLength; i++ {
		if firstId[i] != secondId[i] {
			return false
		}
	}
	return true
}

func (host kademliaId) calcDist(target kademliaId) kademliaId {
	distance := kademliaId{}
	for i := 0; i < idLength; i++ {
		distance[i] = host[i] ^ target[i]
	}
	return distance
}

func (kademliaId kademliaId) toString() string {
	return hex.EncodeToString(kademliaId[:])
}
