package main

import "fmt"

type Employee struct {
	name   string
	salary int
	sales  int
	bonus  int
}

const BONUS_PERCENTAGE = 10

func getBonusPercentage(salary int) int {
	percentage := (salary * BONUS_PERCENTAGE) / 100
	return percentage
}

func findEmployeeBonus(salary, noOfSales int) int {
	bonusPercentage := getBonusPercentage(salary)
	bonus := bonusPercentage * noOfSales
	return bonus
}

func main() {
	var john = Employee{"John", 5000, 5, 0}
	john.bonus = findEmployeeBonus(john.salary, john.sales)
	fmt.Println(john.bonus)
}
