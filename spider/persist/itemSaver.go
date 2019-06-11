package persist

import "log"

func ItemSaver() chan interface{}{
	out:=make(chan  interface{})
	go func() {
		itemCount:=0
		log.Printf("item saver got #%d,%v",itemCount,<-out)
		itemCount++
	}()
	return out
}