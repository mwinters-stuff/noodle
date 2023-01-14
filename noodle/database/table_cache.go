package database

import (
	"github.com/mwinters-stuff/noodle/server/models"
	"golang.design/x/go2generics/maps"
)

type Cacheable interface {
	models.ApplicationTab |
		models.Application |
		models.ApplicationTemplate |
		models.GroupApplications |
		models.Group |
		models.Tab |
		models.UserApplications |
		models.UserGroup |
		models.User
}

//go:generate go run github.com/vektra/mockery/v2 --with-expecter --name TableCache
type TableCache[V Cacheable] interface {
	Add(index int64, value V)
	DeleteIndex(index int64)
	DeleteValue(value V)
	Update(index int64, value V)
	GetID(index int64) (bool, V)

	ForEach(f func(index int64, value V) bool)
	Find(f func(index int64, value V) bool) (bool, *V)
	FindAll(f func(index int64, value V) bool) (bool, []V)
}

type TableCacheImpl[V Cacheable] struct {
	cache map[int64]V
}

func (i TableCacheImpl[V]) Add(index int64, value V) {
	i.cache[index] = value
}

func (i TableCacheImpl[V]) DeleteIndex(index int64) {
	delete(i.cache, index)
}

func (i TableCacheImpl[V]) DeleteValue(valueToDelete V) {
	maps.Filter(i.cache, func(index int64, value V) bool {
		return valueToDelete != value
	})
}

func (i TableCacheImpl[V]) Update(index int64, value V) {
	i.cache[index] = value
}

func (i TableCacheImpl[V]) GetID(index int64) (bool, V) {
	v, e := i.cache[index]
	return e, v
}

func (i TableCacheImpl[V]) ForEach(f func(index int64, value V) bool) {
	for index, v := range i.cache {
		if !f(index, v) {
			return
		}
	}
}

func (i TableCacheImpl[V]) Find(f func(index int64, value V) bool) (bool, *V) {
	for index, v := range i.cache {
		if f(index, v) {
			return true, &v
		}
	}
	return false, nil
}

func (i TableCacheImpl[V]) FindAll(f func(index int64, value V) bool) (bool, []V) {
	var result = []V{}

	for index, v := range i.cache {
		if f(index, v) {
			result = append(result, v)
		}
	}
	return len(result) > 0, result

}

func NewTableCache[T Cacheable]() TableCache[T] {
	return TableCacheImpl[T]{
		cache: make(map[int64]T),
	}
}
