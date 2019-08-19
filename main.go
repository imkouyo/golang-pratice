package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateArray(amount int) []int{
	var list = make([]int, amount)
	for i := 0; i < amount; i ++{
		list[i] = rand.Intn(100)
	}
	return list
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(mySort.BubbleSort(GenerateArray(10)))
	//fmt.Println(mySort.SelectionSort(GenerateArray(10)))
	//firstX := rand.Intn(8)
	//firstY := rand.Intn(8)
	//myChest.KnightGoCheckerboard()
	fmt.Println(true && false)
}

