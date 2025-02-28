package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int
	Date        time.Time
	Description string
	Amount      float64
}

func main() {

	expenses := []Expense{
		{Id: 1, Date: time.Now(), Description: "some test", Amount: 10.3},
	}

	var expense Expense
	fmt.Println("Enter expense amount")
	fmt.Scan(&expense.Amount)

	fmt.Println("Enter expense description")
	fmt.Scan(&expense.Description)

	expense.Id = len(expenses) + 1
	expense.Date = time.Now()

	expenses = append(expenses, expense)

	fmt.Println("======================")
	for _, value := range expenses {
		fmt.Printf("-  %d %s %s %f \n", value.Id, value.Date, value.Description, value.Amount)
	}

}
