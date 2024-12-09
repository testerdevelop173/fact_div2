package main

import (
	"fmt"
	"math/big"
	"os"
	"sync"
)

func factorial(n *big.Int, ch chan *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n.Cmp(big.NewInt(0)) == 0 {
		ch <- big.NewInt(1)
		return
	}
	partialResult := big.NewInt(0)
	wg.Add(1)
	go factorial(new(big.Int).Sub(n, big.NewInt(1)), ch, wg)
	partialResult.Mul(n, <-ch)
	ch <- partialResult
}

func main() {
	var wg sync.WaitGroup

	var n int64 // число факториала
	var m int64 // число гоурутин

	//n := int64(33) число факториала

	fmt.Println("Введите число факториала")
	fmt.Fscan(os.Stdin, &n)
	fmt.Println("Введите количество потоков")
	fmt.Fscan(os.Stdin, &m)

	result := big.NewInt(1)

	wg.Add(1)
	ch := make(chan *big.Int, 1)
	go factorial(big.NewInt(n), ch, &wg)

	wg.Wait()
	result.Set(<-ch)

	fmt.Println("Factorial calculated using goroutines:", result)
}
