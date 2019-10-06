package goroutineTest

import (
	"fmt"
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
