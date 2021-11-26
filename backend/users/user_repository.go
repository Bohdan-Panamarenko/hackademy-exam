package users

import (
	"errors"
	"sync"
)

type InMemoryUserStorage struct {
	lock    sync.RWMutex
	storage map[string]*UserToDoStorage
}

func NewInMemoryUserStorage() *InMemoryUserStorage {
	return &InMemoryUserStorage{
		lock:    sync.RWMutex{},
		storage: make(map[string]*UserToDoStorage),
	}
}

func (s *InMemoryUserStorage) Add(key string, user *UserToDoStorage) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.storage[key] != nil {
		return errors.New("Key '" + key + "' already exists")
	}

	s.storage[key] = user
	return nil
}

func (s *InMemoryUserStorage) Update(key string, user *UserToDoStorage) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.storage[key] == nil {
		return errors.New("Key '" + key + "' doesn't exist")
	}

	s.storage[key] = user
	return nil
}

func (s *InMemoryUserStorage) Get(key string) (user *UserToDoStorage, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	user, exists := s.storage[key]
	if exists {
		return user, nil
	}
	return (nil), errors.New("Key '" + key + "' doesn't exist")
}

func (s *InMemoryUserStorage) Delete(key string) (user *UserToDoStorage, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	user, exists := s.storage[key]
	if exists {
		delete(s.storage, key)
		return user, nil
	}
	return (nil), errors.New("Key '" + key + "' doesn't exist")
}
