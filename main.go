package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

type MyArray struct {
	Data     []int
	Size     int
	SizeUsed int
}

func NewDynamicArray(initialCapacity int) *MyArray {
	MyArray := MyArray{
		Data:     make([]int, 0, initialCapacity),
		Size:     initialCapacity,
		SizeUsed: 0,
	}
	return &MyArray
}

func (arr *MyArray) Add(elemnt int) {
	arr.Data = append(arr.Data, elemnt)
	arr.SizeUsed++
}

func (arr *MyArray) Remove(index int) {
	if index < 0 || index >= arr.SizeUsed {
		fmt.Println("Index out of range")
	}
	arr.Data = append(arr.Data[:index], arr.Data[index+1:]...)

	// for i := index; i < arr.SizeUsed-1; i++ {
	// 	arr.Data[i] = arr.Data[i+1]
	// }

	arr.SizeUsed--

}

func (arr *MyArray) InsertElement(element int, index int) error {
	if index < 0 || index >= arr.SizeUsed {
		return errors.New("Index out of range")
	}

	oldelemnt := arr.Data[index]

	//copy(arr.Data[index+1:], arr.Data[index:])
	arr.Data[index] = element
	arr.Data = append(arr.Data, oldelemnt)
	arr.SizeUsed++
	return nil
}
func (arr *MyArray) InsertElementInSorted(element int, index int) error {
	if index < 0 || index > arr.SizeUsed {

		return errors.New("Index out of range")
	}

	//oldelemnt := arr.Data[index]
	arr.Data = append(arr.Data, 0)

	fmt.Println("index", index, "size", arr.SizeUsed)
	for i := arr.SizeUsed; index < i; i-- {
		arr.Data[i] = arr.Data[i-1]
	}
	arr.Data[index] = element
	// arr.Data = append(arr.Data, element)
	// copy(arr.Data[index+1:], arr.Data[index:])
	// arr.Data[index] = element
	arr.SizeUsed++
	return nil
}

func (arr *MyArray) Transverse() {
	fmt.Println("t size:", arr.SizeUsed)
	for i := 0; i < arr.SizeUsed; i++ {
		fmt.Printf("index: %d array value %d", i, arr.Data[i])
		fmt.Println()
	}

}
func (arr *MyArray) LinearSearch(value int) (int, error) {
	for i := 0; i < arr.SizeUsed; i++ {
		if arr.Data[i] == value {
			return i, nil
		}
	}
	return -1, errors.New("Element not found")
}

func (arr *MyArray) LinearSearchData(value int) (int, error) {

	return -1, errors.New("ELement have error")
}

func Switchcase() {
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}
	switch runtime.GOOS {
	case "darwin1":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux os is installed")
	case "windows":
		fmt.Println("microsoft window is installed")
		// default:
		// 	fmt.Println("microsoft window is installed")
		// 	fmt.Println(runtime.GOOS)
	}
}
func main() {

	//arr := NewDynamicArray(30)
	// arr.Data = []int{2, 4, 6, 8, 23, 87}
	// arr.SizeUsed = 6
	// // Add elements to the array
	// arr.Add(9)
	// arr.Add(7)
	// fmt.Println("Array after adding elements:", arr.Data)
	//arr.Remove(5)
	// arr.Remove(3)
	//arr.Remove(2)
	//arr.InsertElement(5, 2)
	// arr.Transverse()
	// _ = arr.InsertElementInSorted(6, 2)
	// arr.Transverse()
	// _ = arr.InsertElementInSorted(4, 5)
	// arr.Transverse()
	// val, err := arr.LinearSearch(23)
	// if err != nil {
	// 	//fmt.Println("Element not found")
	// }
	// fmt.Println("Element found at index:", val)
	// fmt.Println("Print element:", arr.Data)
	//fmt.Printf("Array after remove elements: %v\n", arr.InsertElement(6, 1))
	//fmt.Println("Array after insert elements:", arr.Data)

	//Switchcase()
	// x := 23
	// val := test_closer(x)

	// fmt.Println(val())
	// fmt.Println(val())
	// y := 12
	// val1 := test_closer(y)
	// fmt.Println(val1())
	// fmt.Println(val1())

	//test_pointer()
	//test_array()
	//test_map()
	//test_slice()
	//test_concurency()
	//test_channel()
	//test_channel_concurrency()
	//test_channel_directions()
	test_select_with_channel()
	timeout_example_channel()
}
