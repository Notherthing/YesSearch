package leveldbOp

import (
	"YesSearch/rbmap"
	"YesSearch/tokenizer"
	"YesSearch/translate"
	"fmt"
	"sync"

)

var channelSize int64
var wg sync.WaitGroup

func init() {
	channelSize = 100
}

func ReturnResCache(caption string) {
	if !Ishan(caption) {
		caption = translate.Query(caption)
	}
	fmt.Println(caption)
	mp := rbmap.New() //rbtree的根节点
	words := tokenizer.Cut_word(caption)

	buff := make(chan rbmap.Info, channelSize)
	for _, word := range words {
		if !Ishan(word) {
			continue
		}
		//wg.Add(1)
		ReadWordUrlCache(word, mp, caption)
		//go readEach(word, mp, caption)
	}
	wg.Add(1)
	go producer(mp, buff) //将map中数据中序输入缓冲区
	consumer(buff)
	wg.Wait()
}

func consumer(buff <-chan rbmap.Info) {
	for  tmp := range buff {
		//url直接打成数据包传输，caption
		fmt.Println(tmp.Url,tmp.Caption)
	}
}

func producer(mp rbmap.Map, buff chan<-rbmap.Info) {
	defer wg.Done()
	mp.OutPut(buff)
	close(buff) //关闭通道
}
/*
func readEach(word string, mp rbmap.Map, caption string) {
	defer wg.Done()
	ReadWordUrl(word, mp, caption)
}
*/

func ReturnResMain(caption string) {
	if !Ishan(caption) {
		caption = translate.Query(caption)
	}
	fmt.Println(caption)
	mp := rbmap.New() //rbtree的根节点
	words := tokenizer.Cut_word(caption)

	buff := make(chan rbmap.Info, channelSize)
	for _, word := range words {
		if !Ishan(word) {
			continue
		}
		//wg.Add(1)
		ReadWordUrlMain(word, mp, caption)
		//go readEach(word, mp, caption)
	}
	wg.Add(1)
	go producer(mp, buff) //将map中数据中序输入缓冲区
	consumer(buff)
	wg.Wait()
}