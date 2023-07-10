package main

func main() {
	// if statement.
	// if <expression> {
	// 	<code>
	// }
	if true {
		println("true")
	}

	// if-else statement.
	// if <expression> {
	// 	<code>
	// } else {
	// 	<code>
	// }
	if false {
		println("true")
	} else {
		println("false")
	}

	// if-else if-else statement.
	// if <expression> {
	// 	<code>
	// } else if <expression> {
	// 	<code>
	// } else {
	// 	<code>
	// }
	if false {
		println("true")
	} else if true {
		println("false")
	} else {
		println("false")
	}

	// if statement with short statement.
	// if <short statement>; <expression> {
	// 	<code>
	// }
	if a := 1; a == 1 {
		println("true")
	}

	// if-else statement with short statement.
	// if <short statement>; <expression> {
	// 	<code>
	// } else {
	// 	<code>
	// }
	if a := 1; a == 1 {
		println("true")
	} else {
		println("false")
	}

	// if-else if-else statement with short statement.
	// if <short statement>; <expression> {
	// 	<code>
	// } else if <expression> {
	// 	<code>
	// } else {
	// 	<code>
	// }
	if a := 1; a == 1 {
		println("true")
	} else if a == 2 {
		println("false")
	} else {
		println("false")
	}
}
