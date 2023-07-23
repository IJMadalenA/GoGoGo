package main

import (
	"fmt"
)

func main() {
	// Array.
	// ways to declare an array:
	// var <name> [<size>] <type>
	// var <name> = [<size>]<type>{<values>}
	// <name> := [<size>]<type>{<values>}
	// <name> := [...]<type>{<values>}

	// An array is a collection of elements of the same type placed in contiguous memory locations that can be individually referenced by adding an index to a unique identifier.
	// The first element in an array is always at index 0, and the last element is at index length-1, where length is the total number of elements in the array.
	// The length of an array is part of its type. The types [10]int and [20]int are distinct, so arrays cannot be resized.
	// Arrays are useful when planning the detailed layout of memory for variables that must be allocated contiguously, such as buffers or image pixels or elements of a matrix.
	// Arrays are essentially just pointers to the first element in the array, and the length of the array. This is why arrays are passed by value, and not by reference.

	// The main characteristic of arrays is that they are fixed length, which means that once they are created, they cannot grow or shrink.
	// Another characteristic of arrays is that they are homogenous, which means that all the elements in an array must be of the same type.
	// Arrays are value types, which means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable.
	// If changes are made to the new variable, it will not affect the original array, and vice versa.

	// ways to declare an array:
	var array1 [5]int
	var array2 = [5]int{1, 2, 3, 4, 5}
	array3 := [5]int{1, 2, 3, 4, 5}
	array4 := [...]int{1, 2, 3, 4, 5}

	fmt.Println("\nWays to declare an array:")
	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
	fmt.Println("array3:", array3)
	fmt.Println("array4:", array4)

	// len function.
	// The len function returns the length of an array, which is the number of elements it contains.
	fmt.Println("\nlen function:")
	fmt.Println("len(array1):", len(array1))
	fmt.Println("len(array2):", len(array2))
	fmt.Println("len(array3):", len(array3))
	fmt.Println("len(array4):", len(array4))

	// cap function.
	// The cap function returns the capacity of an array, which is the number of elements it can hold.
	fmt.Println("\ncap function:")
	fmt.Println("cap(array1):", cap(array1))
	fmt.Println("cap(array2):", cap(array2))
	fmt.Println("cap(array3):", cap(array3))
	fmt.Println("cap(array4):", cap(array4))

	// Slices.
	// A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change.
	// A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
	// A slice is a reference to an underlying array. If you assign one slice to another, both refer to the same array.
	// If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array.
	// You can think of a slice as a struct with a pointer, length, and capacity. To access the length of a slice, use the built-in len function. To access the capacity of a slice, use the built-in cap function.
	// The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	//Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

	// ways to declare a slice:
	var slice1 []int
	var slice2 = []int{1, 2, 3, 4, 5}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{1, 2, 3, 4, 5}

	fmt.Println("\nWays to declare a slice:")
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)

	// Accessing elements of a slice.
	// The elements of a slice are accessed using the subscript notation, [i], where i is the index of the element to be accessed.
	// The first element in a slice is always at index 0, and the last element is at index length-1, where length is the total number of elements in the slice.
	// The length of a slice is part of its type. The types []int and []float64 are distinct, so slices cannot be resized.
	fmt.Println("\nAccessing elements of a slice:")
	fmt.Println(slice2[0])
	fmt.Println(slice2[:3])
	fmt.Println(slice2[2:])
	fmt.Println(slice2[1:3])

	// len function.
	// The len function returns the length of a slice, which is the number of elements it contains.
	fmt.Println("\nlen function for slices:")
	fmt.Println("len(slice1):", len(slice1))
	fmt.Println("len(slice2):", len(slice2))
	fmt.Println("len(slice3):", len(slice3))
	fmt.Println("len(slice4):", len(slice4))

	// cap function.
	// The cap function returns the capacity of a slice, which is the number of elements it can hold.
	fmt.Println("\ncap function for slices:")
	fmt.Println("cap(slice1):", cap(slice1))
	fmt.Println("cap(slice2):", cap(slice2))
	fmt.Println("cap(slice3):", cap(slice3))
	fmt.Println("cap(slice4):", cap(slice4))

	// append function.
	// The append function appends elements to the end of a slice, and grows the slice if a greater capacity is needed.
	slice4 = append(slice4, 6, 7, 8, 9, 10)
	fmt.Println("\nappend function for slices:")
	fmt.Println("slice4:", slice4)
	fmt.Println("len(slice4):", len(slice4))
	fmt.Println("cap(slice4):", cap(slice4))

	// copy function.
	// The copy function copies elements from a source slice into a destination slice.
	fmt.Println("\ncopy function for slices:")
	slice5 := make([]int, 5)
	copy(slice5, slice4)
	fmt.Println("slice5:", slice5)
	fmt.Println("len(slice5):", len(slice5))
	fmt.Println("cap(slice5):", cap(slice5))

	// range function.
	// The range function iterates over a slice or map.
	fmt.Println("\nrange function for slices:")
	for index, value := range slice5 {
		fmt.Println("index:", index, "value:", value)
	}

	// The range function ignores the index or value by assigning it to the blank identifier.
	fmt.Println("\nrange function for slices (ignoring the index):")
	for _, value := range slice5 {
		fmt.Println("value:", value)
	}
	fmt.Println("\nrange function for slices (ignoring the value):")
	for index, _ := range slice5 {
		fmt.Println("index:", index)
	}

	// The range function can also be used to iterate over a string.
	fmt.Println("\nrange function for strings:")
	for index, value := range "Hello World." {
		fmt.Println("index:", index, "value:", value)
	}

	// The range function can also be used to iterate over a map.
	fmt.Println("\nrange function for maps:")
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range map1 {
		fmt.Println("key:", key, "value:", value)
	}

	// The range function can also be used to iterate over a channel.
	fmt.Println("\nrange function for channels:")
	channel1 := make(chan int)
	go func() {
		channel1 <- 1
		channel1 <- 2
		channel1 <- 3
		close(channel1)
	}()
	for value := range channel1 {
		fmt.Println("value:", value)
	}

	// Multidimensional arrays and slices.
	// Multidimensional arrays and slices are arrays and slices of arrays and slices, are useful when working with matrices and tables.
	// Are declared in the same way as one-dimensional arrays and slices, except that each element is itself an array or slice.
	// Multidimensional arrays and slices are accessed using multiple indices, the first index selects an element in the outermost array or slice, and the second index selects an element in the innermost array or slice.

	// ways to declare a multidimensional array:
	var array5 [5][5]int
	var array6 = [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	array7 := [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	array8 := [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	fmt.Println("\nWays to declare a multidimensional array:")
	fmt.Println("array5:", array5)
	fmt.Println("array6:", array6)
	fmt.Println("array7:", array7)
	fmt.Println("array8:", array8)

	// ways to declare a multidimensional slice:
	var slice6 [][]int
	var slice7 = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	slice8 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	slice9 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	fmt.Println("\nWays to declare a multidimensional slice:")
	fmt.Println("slice6:", slice6)
	fmt.Println("slice7:", slice7)
	fmt.Println("slice8:", slice8)
	fmt.Println("slice9:", slice9)
}
