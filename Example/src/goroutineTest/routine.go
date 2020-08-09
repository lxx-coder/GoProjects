package goroutineTest

import (
	"fmt"
	"sync"
	"time"
)

func Routine(){
	theMine := [5]string{"ore1","ore2","ore3","ore4","ore5"}
	oreChan := make(chan string,3)

	//Finder
	go func(mine [5]string) {
		for _,item := range mine {
			oreChan <- item
			fmt.Println("Miner sent ", item)
		}
	}(theMine)

	// ore breaker
	go func() {
		for foundOre := range oreChan {
			//foundOre := <-oreChan //receive
			fmt.Println("Miner: Received ", foundOre, " from finder")
		}
	}()
	<-time.After(time.Second * 5)
}

func calc(a int, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Println("%d + %d = %d", a,b,c)
	defer n.Done()
}

func Calc() {
	var go_sync sync.WaitGroup
	for i := 0; i < 10; i++ {
		go_sync.Add(1)
		go calc(i, i+1, &go_sync)
	}
	go_sync.Wait()
}


