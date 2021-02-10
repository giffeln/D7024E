package main

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const idLength = 20

type kademliaId [idLength]byte

func newKadId(data string) (ret kademliaId) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < len(decoded); i++ {
		ret[i] = decoded[i]
	}

	return
}

func randomKadId() (ret kademliaId) {
	for i := 0; i < idLength; i++ {
		ret[i] = uint8(rand.Int63n(time.Now().UnixNano()))
	}

	return
}

func (firstId kademliaId) smaller(secondId kademliaId) bool {
	for i := 0; i < idLength; i++ {
		if firstId[i] != secondId[i] {
			return firstId[i] < secondId[i]
		}
	}
	return false
}

func (firstId kademliaId) equals(secondId kademliaId) bool {
	for i := 0; i < idLength; i++ {
		if firstId[i] != secondId[i] {
			return false
		}
	}
	return true
}

func (host kademliaId) calcDist(target kademliaId) (ret kademliaId) {
	for i := 0; i < idLength; i++ {
		ret[i] = host[i] ^ target[i]
	}
	return
}

func (node kademliaId) prefixLen() (ret int) {
	for i := 0; i < idLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return idLength*8 - 1
}

func (kademliaId kademliaId) toString() string {
	return hex.EncodeToString(kademliaId[:])
}
