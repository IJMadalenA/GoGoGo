package main

import "fmt"

func main() {
	first_map := make(map[string]int)
	first_map["one"] = 1
	first_map["two"] = 2
	first_map["three"] = 3

	second_map := map[string]int{"one": 1, "two": 2, "three": 3}
	second_map["four"] = 4
	second_map["five"] = 5
	second_map["six"] = 6

	third_map := map[string]int{}
	third_map["one"] = 1
	third_map["two"] = 2
	third_map["three"] = 3

	fmt.Println("first_map:", first_map)
	fmt.Println("second_map:", second_map)
	fmt.Println("third_map:", third_map)

	// The delete function can be used to delete a key-value pair from a map.
	delete(second_map, "one")
	fmt.Println("\nsecond_map after deleting \"one\":", second_map)

	// The len function can be used to get the number of key-value pairs in a map.
	fmt.Println("\nlen(first_map):", len(first_map))
	fmt.Println("len(second_map):", len(second_map))
	fmt.Println("len(third_map):", len(third_map))

	// The range function can be used to iterate over a map.
	fmt.Println("\nrange function for maps:")
	for key, value := range first_map {
		fmt.Println("key:", key, "value:", value)
	}

	// The range function ignores the key or value by assigning it to the blank identifier.
	fmt.Println("\nrange function for maps (ignoring the key):")
	for _, value := range first_map {
		fmt.Println("value:", value)
	}
	fmt.Println("\nrange function for maps (ignoring the value):")
	for key, _ := range first_map {
		fmt.Println("key:", key)
	}

	// Get the value of a key.
	fmt.Println("\nfirst_map[\"one\"]:", first_map["one"])

	// Get the value of a key that does not exist.
	fmt.Println("\nfirst_map[\"four\"]:", first_map["four"])
	// If the key does not exist, the value returned is the zero value of the value type.
	// In this case, the value type is int, so the zero value is 0. This is not the best way to check if a key exists in a map.
	// The best way to check if a key exists in a map is to use the two-value assignment form of the map index expression.
	// The first value is the value stored in the map, and the second value is a boolean that reports whether the key was found in the map.
	value, ok := first_map["four"]
	fmt.Println("value:", value, "ok:", ok)
	// The ok variable is a boolean that reports whether the key was found in the map.
	// If the key is found, ok is true, and the value stored in the map is assigned to the value variable.
	// If the key is not found, ok is false, and the value stored in the map is the zero value of the value type.
	// In this case, the value type is int, so the zero value is 0.
	// The ok variable can be used to check if the key exists in the map.
	if ok {
		fmt.Println("first_map[\"four\"] exists.")
	} else {
		fmt.Println("first_map[\"four\"] does not exist.")
	}

	// The two-value assignment form of the map index expression can be used to check if a key exists in a map.
	_, ok = first_map["four"]
	if ok {
		fmt.Println("first_map[\"four\"] exists.")
	} else {
		fmt.Println("first_map[\"four\"] does not exist.")
	}

	_, ok = first_map["one"]
	if ok {
		fmt.Println("first_map[\"one\"] exists.")
	} else {
		fmt.Println("first_map[\"one\"] does not exist.")
	}
}
