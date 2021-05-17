package main

import (
	"Go5-6/counters"
	"fmt"
)
func main0()  {
	counter := counters.AlterCounter(10)
	counter0 := counters.New(10)
	fmt.Printf("Counter:%d\n",counter)
	fmt.Printf("Counter0:%d\n",counter0)
}
