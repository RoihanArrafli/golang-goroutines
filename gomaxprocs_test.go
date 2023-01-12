package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		func(){
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	
	totalcpu := runtime.NumCPU()
	fmt.Println("total cpu : ",totalcpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread : ",totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine : ", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		func(){
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	
	totalcpu := runtime.NumCPU()
	fmt.Println("total cpu : ",totalcpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread : ",totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine : ", totalGoroutine)

	group.Wait()
}