package WorkerPool

import (
	"fmt"
	"math/rand"
	"time"
)

func doSomeWork(result chan<- int, num chan<- int, numWorker int) {
	rand.Seed(time.Now().UTC().UnixNano())
	sec := rand.Intn(3) + 1
	time.Sleep(time.Duration(sec) * time.Second)
	result <- numWorker * numWorker
	num <- numWorker
}

func WorkerPool(numWorkers int) {
	result := make(chan int)
	num := make(chan int)
	for numWorker := 1; numWorker <= numWorkers; numWorker++ {
		go doSomeWork(result, num, numWorker)
	}
FOREMAN:
	for {
		a := <-result
		b := <-num
		numWorkers--
		if a == b*b {
			fmt.Println("Работник №", b, "выполнил работу")
		} else {
			fmt.Println("Работник №", b, "не выполнил работу")
		}
		if numWorkers == 0 {
			break FOREMAN
		}
	}
}
