package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSum2(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{8, 9, 9},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("GetMax was incorrect, got: %d, want: %d.", max, item.c)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{10, 55},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.b {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", fib, item.b)
		}
	}
}

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Person: Person{
							Name: "Test",
							Age:  30,
						},
						id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Test",
						Age:  30,
						DNI:  "12345678",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Employee: Employee{
					Person: Person{
						Name: "Test",
						Age:  30,
					},
					id:       1,
					Position: "CEO",
				},
				Person: Person{
					Name: "Test",
					Age:  30,
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		employee, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Error(err)
		}

		if employee.Age != test.expectedEmployee.Age {
			t.Errorf("Expected employee age to be %d, got %d", test.expectedEmployee.Age, employee.Age)
		}

		if employee.Name != test.expectedEmployee.Name {
			t.Errorf("Expected employee name to be %s, got %s", test.expectedEmployee.Name, employee.Name)
		}

		if employee.DNI != test.expectedEmployee.DNI {
			t.Errorf("Expected employee DNI to be %s, got %s", test.expectedEmployee.DNI, employee.DNI)
		}

		if employee.id != test.expectedEmployee.id {
			t.Errorf("Expected employee id to be %d, got %d", test.expectedEmployee.id, employee.id)
		}
	}
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
