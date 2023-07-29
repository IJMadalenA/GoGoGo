package main

func main() {
	// continue statement.
	// The continue statement is used to skip the current iteration of the for loop, and the next iteration will be executed.
	println("continue statement.")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		println(i)
	}

	// break statement.
	// the break statement is used to terminate the for loop, switch statement, etc.
	println("break statement.")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}
}
