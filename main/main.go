package main

import "fmt"

func deleteElement(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func deleteElementEfficient(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func deleteElementGeneric[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func deleteElementWithShrink[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	if cap(slice) > 2*len(slice) {
		newSlice := make([]T, len(slice))
		copy(newSlice, slice)
		slice = newSlice
	}
	return slice
}
func main() {
	// 初始切片
	slice := []int{10, 20, 30, 40, 50}
	index := 2

	// 常规删除操作
	deletedSlice := deleteElement(slice, index)
	fmt.Println("After deleteElement:", deletedSlice)

	// 高效删除操作
	deletedSliceEfficient := deleteElementEfficient(slice, index)
	fmt.Println("After deleteElementEfficient:", deletedSliceEfficient)

	// 泛型删除操作
	stringSlice := []string{"A", "B", "C", "D", "E"}
	deletedStringSlice := deleteElementGeneric(stringSlice, index)
	fmt.Println("After deleteElementGeneric with string slice:", deletedStringSlice)

	// 泛型删除并缩容操作
	deletedSliceWithShrink := deleteElementWithShrink(slice, index)
	fmt.Println("After deleteElementWithShrink:", deletedSliceWithShrink)
}
