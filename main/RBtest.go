package main

import (
	"YesSearch/rbmap"
	"github.com/hyperjumptech/beda"
	"YesSearch/basic"
	"fmt"
)




func main1() {
	//checkMap := make(map[string]struct{})

	// 1. new a map default is rb tree
	// 1. 新建一个 Map，默认为标准红黑树实现
	m := rbmap.New()
	//m = gomap.NewAVLMap()    // avl tree better version
	//m = gomap.NewAVLRecursionMap() // avl tree bad version

	//m.SetComparator(comparatorInt) // set inner comparator

	s1:="我真帅"
	s2:=[]byte("我是傻子")
	key := basic.BytesCombine(
		basic.Float32ToByte(beda.NewStringDiff(s1, string(s2)).JaroWinklerDistance(0.1)),
		s2)
	//fmt.Println(key[0:4])
	fmt.Println((beda.NewStringDiff(s1, string(s2)).JaroWinklerDistance(0.1)))
	fmt.Println(basic.ByteToFloat32(key[0:4]))
	m.Put(string(key),"123")

	fmt.Println("map len is ", m.Len())

	// 3. can iterator
	// 3. 迭代器使用
	iterator := m.Iterator()
	for iterator.HasNext() {
		k, v := iterator.Next()
		fmt.Printf("Iterator key:%s,value %v\n", k, v)
	}

	/*
	// 4. get key
	// 4. 获取键中的值
	key := "9"
	value, exist := m.Get(key)
	if exist {
		fmt.Printf("%s exist, value is:%s\n", key, value)
	} else {
		fmt.Printf("%s not exist\n", key)
	}

	// 5. get int will err
	// 5. 获取键中的值，并且指定类型，因为类型是字符串，所以转成整形会报错
	_, _, err := m.GetInt(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 6. check is a rb tree
	// 6. 内部辅助，检查是不是一颗正常的红黑树
	if m.Check() {
		fmt.Println("is a rb tree,len:", m.Len())
	}

	// 7. delete '9' then find '9'
	// 7. 删除键 '9' 然后再找 '9'
	m.Delete(key)
	value, exist = m.Get(key)
	if exist {
		fmt.Printf("%s exist, value is:%s\n", key, value)
	} else {
		fmt.Printf("%s not exist\n", key)
	}

	// 8. key list
	// 8. 获取键列表
	fmt.Printf("keyList:%#v,len:%d\n", m.KeyList(), m.Len())
	fmt.Printf("keySortList:%#v,len:%d\n", m.KeySortedList(), m.Len())

	// 9. delete many
	// 9. 删除很多键值对
	for key := range checkMap {
		v, ok := m.Get(key)
		fmt.Printf("find %s:%v=%v, delete key:%s\n", key, v, ok, key)
		m.Delete(key)
		if !m.Check() {
			fmt.Println("is not a rb tree,len:", m.Len())
		}
	}
*/
	// 10. key list
	// 10. 获取键列表
	fmt.Printf("keyList:%#v,len:%d\n", m.KeyList(), m.Len())
	fmt.Printf("keySortList:%#v,len:%d\n", m.KeySortedList(), m.Len())

	// 11. check is a rb tree
	// 11. 再次检查是否是一颗正常的红黑树
	if m.Check() {
		fmt.Println("is a rb tree,len:", m.Len())
	}
}
