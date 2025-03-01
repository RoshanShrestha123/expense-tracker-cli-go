package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	Id          int
	Date        time.Time
	Description string
	Amount      float32
}

type Argument struct {
	Id          int
	Month       string
	Description string
	Amount      float32
	Action      string
}

func main() {

	args := os.Args
	userInput := Argument{}

	commands := args[1:]

	for i, value := range commands {
		if i == 0 {
			userInput.Action = value
			continue
		}

		switch value {
		case "--description":
			data := commands[i+1 : i+2]

			if len(data) == 0 {
				log.Fatal("invalid input")
			}
			userInput.Description = data[0]

		case "--month":
			data := commands[i+1 : i+2]

			if len(data) == 0 {
				log.Fatal("invalid input")
			}
			userInput.Month = data[0]

		case "--amount":
			data := commands[i+1 : i+2]

			if len(data) == 0 {
				log.Fatal("invalid input")
			}
			parsedAmount, _ := strconv.ParseFloat(data[0], 64)
			userInput.Amount = float32(parsedAmount)

		case "--id":
			data := commands[i+1 : i+2]

			if len(data) == 0 {
				log.Fatal("invalid input")
			}
			id, _ := strconv.ParseInt(data[0], 10, 64)
			userInput.Id = int(id)
		}
	}

	expenses := []Expense{}

	switch userInput.Action {
	case "add":
		expense := Expense{
			Id:          len(expenses) + 1,
			Date:        time.Now(),
			Description: userInput.Description,
			Amount:      (userInput.Amount),
		}

		expenses = append(expenses, expense)

		for _, value := range expenses {
			fmt.Printf("%d %s %f at %d\n", value.Id, value.Description, value.Amount, value.Date.Month())
		}

	case "list":
		for _, value := range expenses {
			fmt.Printf("%d %s %f at %d\n", value.Id, value.Description, value.Amount, value.Date.Month())
		}

	}

}
