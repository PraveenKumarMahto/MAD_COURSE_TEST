package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func processingCustomerByCashier(cashierId int, customerin <-chan int, timePerCustomer int, wg *sync.WaitGroup) {
	for custId := range customerin {

		custInTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%s --> Cashier %d: Customer %d Started\n", custInTime, cashierId, custId)

		time.Sleep(time.Second)

		custOutTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%s --> Cashier %d: Customer %d Completed\n", custOutTime, cashierId, custId)
	}
	wg.Done()
}

func main() {

	numCashiers := flag.Int("numCashiers", 0, "No of cashier")
	numCustomers := flag.Int("numCustomers", 0, "No of customer")
	timePerCustomer := flag.Int("timePerCustomer", 0, "Time per customer")

	flag.Parse()

	customerin := make(chan int, *numCustomers)
	var wg sync.WaitGroup

	wg.Add(*numCashiers)

	simulationStartTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s --> Bank Simulation Started\n", simulationStartTime)

	for cashierid := 1; cashierid <= *numCashiers; cashierid++ {
		go processingCustomerByCashier(cashierid, customerin, *timePerCustomer, &wg)
	}

	for customerid := 1; customerid <= *numCustomers; customerid++ {
		customerin <- customerid
	}
	close(customerin)

	wg.Wait()

	simulationEndTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s --> Bank Simulated Ended\n", simulationEndTime)
}
