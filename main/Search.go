package main

import (
	"YesSearch/leveldbOp"
	"bufio"
	"fmt"
	"os"
	"sync"
)
var wg sync.WaitGroup
func main1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		wg.Add(2)
		go cacheSearch(str)
		go mainSearch(str)
		wg.Wait()
		fmt.Println("this ok")
	}
}

func cacheSearch(str string) {
	defer wg.Done()
	leveldbOp.ReturnResCache(str)
}

func mainSearch(str string) {
	defer wg.Done()
	leveldbOp.ReturnResMain(str)
}