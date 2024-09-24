package main

import (
	"fmt"
	"sync"
	"time"
)

// func Switchcase() {
// 	switch time.Now().Weekday() {
// 	case time.Sunday, time.Saturday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")

// 	}
// }

func test_closer(x int) func() int {
	x = x
	return func() int {
		x++
		return x
	}

}

func test_pointer() {
	i, j := 10, 30
	p := &i
	fmt.Println("value of i is ", i)
	fmt.Println("value of j is ", j)

	fmt.Println("address of i is ", p)
	i = 100
	fmt.Println("value of new i is ", i)
	*p = 200
	fmt.Println("value of 200 i is ", i)
	fmt.Println("value of *p is ", *p)
	q := &j
	j = *q / 10

	fmt.Println("value of j is ", j)

}

func test_array() {
	var a [10]int
	a[5] = 34
	fmt.Println("value of array a is ", a)
	fmt.Println("lenght of array a is ", len(a))
	b := [14]int{1, 34, 6545, 646, 977, 8324}
	fmt.Println("value of array b is ", b)
	fmt.Println("lenght of array b is ", len(b))

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	var multiarray [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			multiarray[i][j] = i + j
		}

		fmt.Println("")
	}
	fmt.Println(multiarray)
}

func test_map() {

	var mapdata = make(map[string]int)
	mapdata["a"] = 12
	mapdata["b"] = 30
	fmt.Println("map value is : ", mapdata)
	delete(mapdata, "a")

	val, del := mapdata["b"]
	fmt.Println("deleeted value of map exist :", del)
	fmt.Println("map deleted value :", val)
}

func test_slice() {
	var s = make([]string, 4)
	s[0] = "PRADEEP"
	s[1] = "test"
	s[2] = " mskw"
	s[3] = "dinesh"
	fmt.Println("slice value s is :", s)

	b := make([]string, 7)
	copy(b, s)

	b[4] = "surssh"
	fmt.Println("vale of slice b is :", b)
	p := s[:3]
	q := s[2:]
	fmt.Println("value of q slice is :", p)
	fmt.Println("value of q slice is :", q)
	b = append(b, "janu", "manu", "darling", "dilruba")
	fmt.Println("value of appnet slsice", b[9])

}

func test_concurency() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		count("go count")
		wg.Done()
	}()
	go func() {
		count("open")
		wg.Done()
	}()

	wg.Wait()

}

func count(thing string) {

	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 300)
	}
}
