package main

import (
	"fmt"
	"sync"
	"time"
)

func processingCustomerByCashier(id int, customerin <-chan int, customerout chan<- int, timePerCustomer int, wg *sync.WaitGroup) {
	for j := range customerin {
		defer wg.Done()
		fmt.Println("-->  Cashier ", id, ": Customer ", j, " Started")
		time.Sleep(time.Second)
		fmt.Println("-->  Cashier ", id, ": Customer ", j, " Completed")
	}
}

func main() {

	var numCashiers int
	var numCustomers int
	var timePerCustomer int

	customerin := make(chan int, numCustomers)
	customerout := make(chan int, numCustomers)
	var wg sync.WaitGroup

	wg.Add(numCashiers)
	fmt.Println("-->  Bank Simulation Started")

	for w := 1; w <= numCashiers; w++ {
		go processingCustomerByCashier(w, customerin, customerout, timePerCustomer, &wg)
	}
	wg.Wait()
}
