package users

import (
	"errors"
	"fmt"
	"sync"
)

type UserToDoStorage struct {
	User
	lists   map[uint]*ToDoList // Map listId -> List
	lock    sync.RWMutex
	factory ToDoListFactory
}

func NewUserToDoStorage() *UserToDoStorage {
	return &UserToDoStorage{
		lists: make(map[uint]*ToDoList),
	}
}

func NewUserToDoStorageWithUser(user User) *UserToDoStorage {
	storage := NewUserToDoStorage()
	storage.User = user
	return storage
}

func (utds *UserToDoStorage) NewToDoList(name string) uint {
	newToDoList, id := utds.factory.NewToDoList(name)
	utds.Add(newToDoList.id, newToDoList)
	return id
}

func (utds *UserToDoStorage) Add(key uint, list *ToDoList) error {
	utds.lock.Lock()
	defer utds.lock.Unlock()

	if utds.lists[key] != nil {
		return errors.New("Key '" + fmt.Sprint(key) + "' already exists")
	}
	utds.lists[key] = list
	return nil
}

func (utds *UserToDoStorage) Update(key uint, list *ToDoList) error {
	utds.lock.Lock()
	defer utds.lock.Unlock()

	if utds.lists[key] == nil {
		return errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
	}

	utds.lists[key] = list
	return nil
}

func (utds *UserToDoStorage) Get(key uint) (list *ToDoList, err error) {
	utds.lock.RLock()
	defer utds.lock.RUnlock()

	list, exists := utds.lists[key]
	if exists {
		return list, nil
	}
	return (nil), errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
}

func (utds *UserToDoStorage) Delete(key uint) (list *ToDoList, err error) {
	utds.lock.Lock()
	defer utds.lock.Unlock()

	list, exists := utds.lists[key]
	if exists {
		delete(utds.lists, key)
		return list, nil
	}
	return (nil), errors.New("Key '" + fmt.Sprint(key) + "' doesn't exist")
}
