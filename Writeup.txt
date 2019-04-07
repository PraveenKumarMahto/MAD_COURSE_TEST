
we can implement this by worker pool logic of golang

As we have arguments 
--numCashiers=3 --numCustomers=100 --timePerCustomer=3

we use channels and goroutines.

structure:-

struct data {
    customerid int
    cashierid int
    starttime uint64
    endtime uint64 
    date uint64
    counterno int
    //etc some metadata
}

function required :-

func processingCustomerByCashier(cashierid int,customerid int) error{}
func savetoDB(object of data)
func addCustomerInQueue()
func completedCustomerByCashier(cashierid int,customerid int) error{}

0.> --numCashiers, --numCustomers and  --timePerCustomer,  we can get args.

1.> need to create numCashiers gorutines for processingCustomerByCashier()

2.> need to create buffered channel of size numCustomers

3.> also in function processingCustomerByCashier need to add sleep of timePerCustomer

4.> one funtion that will add customer to the queue.

5.> also we can create one chan at the time of (2) for completingProcess or results that will notify the completedCustomerByCashier funtion till all goroutines complieted their work or we can keep count using global variable using mutex

