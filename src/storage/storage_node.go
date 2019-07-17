package storage

import (
	"errors"
	"github.com/rs/xid"
	"sync"
)

var (
	storageNodes []StorageNode
	mtx sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initializeNodes)
}

func initializeNodes() {
	storageNodes = []StorageNode{}
}

type StorageNode struct {
	ID string `json:"id"`
	Mip string `json:"mip"`
	Description string `json:"description"`
}

func get() []StorageNode {
	return storageNodes
}

func add(mip string, description string) string{
	node := newStorageNode(mip, description)
	mtx.Lock()
	storageNodes = append(storageNodes, node)
	mtx.Unlock()
	return node.ID
}

func delete(id string) error {
	location, err := findStorageLocation(id)
	if err != nil {
		return err
	}
	removeElement(location)
	return nil
}

func deleteByMip(mip string) error {
	location, err := findMip(mip)
	if err != nil {
		return err
	}
	removeElement(location)
	return nil
}

func newStorageNode(mip string, description string) StorageNode {
	return StorageNode{
		ID: xid.New().String(),
		Mip: mip,
		Description: description,
	}
}

func findStorageLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, n := range storageNodes {
		if n.ID == id {
			return i, nil
		}
	}
	return 0, errors.New("could not find storage node")
}

func findMip(mip string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, n := range storageNodes {
		if n.Mip == mip {
			return i, nil
		}
	}
	return 0, errors.New("could not find storage node")
}

func removeElement(i int) {
	mtx.Lock()
	storageNodes = append(storageNodes[:i], storageNodes[i+1:]...)
	mtx.Unlock()
}



