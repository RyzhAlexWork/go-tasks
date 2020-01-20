package workerPool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type work struct {
	i      int
	result int
}

func workerPool(numWorkers int) {
	//Создаётся канал для всех работников
	gener := func() <-chan int {

		out := make(chan int)

		go func() {
			for i := 1; i <= numWorkers; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}
	//Каждый работник из прошлого канала начинает выполнять работу
	cons := func(in <-chan int) <-chan work {

		out := make(chan work)

		go func() {
			for i := range in {
				rand.Seed(time.Now().UTC().UnixNano())
				sec := rand.Intn(3) + 1
				time.Sleep(time.Duration(sec) * time.Second)
				out <- work{i, i * i}
			}
			close(out)
		}()
		return out
	}
	//Сбор всех результатов работников в один канал
	final := func(in []<-chan work) <-chan work {
		out := make(chan work)

		var wg sync.WaitGroup

		output := func(c <-chan work) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}

		wg.Add(len(in))

		for i := 0; i < len(in); i++ {
			go output(in[i])
		}

		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}

	out1 := gener()
	outCh := make([]<-chan work, numWorkers)

	for i := 0; i < numWorkers; i++ {
		outCh[i] = cons(out1)
	}

	out3 := final(outCh)

	//Прораб проверяет работу
	for val := range out3 {
		if val.i*val.i != val.result {
			fmt.Println("Работа выполнена неверно! Номер рабочего:", val.i, ", его результат:", val.result)
		}
	}
}
