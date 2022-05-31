package rbmap

import (
	"YesSearch/basic"
	"strings"
)

type Info struct  {
	Url string
	Caption string
}

type comparator func(key1, key2 string) int64

// Map method
// design to be concurrent safe
// should support int key?
type Map interface {
	Put(key string, value interface{})                            // put key pairs
	Delete(key string)                                            // delete a key
	Get(key string) (value interface{}, exist bool)               // get value from key
	GetInt(key string) (value int, exist bool, err error)         // get value auto change to Int
	GetInt64(key string) (value int64, exist bool, err error)     // get value auto change to Int64
	GetString(key string) (value string, exist bool, err error)   // get value auto change to string
	GetFloat64(key string) (value float64, exist bool, err error) // get value auto change to string
	GetBytes(key string) (value []byte, exist bool, err error)    // get value auto change to []byte
	Contains(key string) (exist bool)                             // map contains key?
	Len() int64                                                   // map key pairs num
	KeyList() []string                                            // map key out to list from top to bottom which is layer order
	KeySortedList() []string                                      // map key out to list sorted
	Iterator() MapIterator                                        // map iterator, iterator from top to bottom which is layer order
	MaxKey() (key string, value interface{}, exist bool)          // find max key pairs
	MinKey() (key string, value interface{}, exist bool)          // find min key pairs
	SetComparator(comparator) Map                                 // set compare func to control key compare
	Check() bool                                                  // just help
	Height() int64                                                // just help
	OutPut(buff chan<- Info)
}

// MapIterator Iterator concurrent not safe
// you should deal by yourself
type MapIterator interface {
	HasNext() bool
	Next() (key string, value interface{})
}


// NewMap default map is rbt implement
func NewMap() Map {
	t := new(rbTree)
	t.c = comparatorDefault
	return t
}

// New default map is rbt implement
func New() Map {
	t := new(rbTree)
	t.c = comparatorDefault
	return t
}

// NewRBMap default map is rbt implement
func NewRBMap() Map {
	t := new(rbTree)
	t.c = comparatorDefault
	return t
}

// compare two key,
//key 以一个float32的相似度和一个剩余url进行排序，
//value为caption
func comparatorDefault(key1, key2 string) int64 {
	key1cmp := basic.ByteToFloat32(basic.String2Bytes(key1)[0:4])
	key2cmp := basic.ByteToFloat32(basic.String2Bytes(key2)[0:4])
	if key1cmp > key2cmp{
		return 1
	} else if key1cmp < key2cmp {
		return -1
	}
	return int64(strings.Compare(key1, key2))
}
