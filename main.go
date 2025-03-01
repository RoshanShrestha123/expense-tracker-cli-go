package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
}

type Argument struct {
	Id          int
	Month       string
	Description string
	Amount      float32
	Action      string
}

func main() {

	fmt.Println(time.Now())
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

	res, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Unable to read the data from file")
	}

	var expenses []Expense

	err = json.Unmarshal(res, &expenses)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(expenses)

	switch userInput.Action {
	case "add":

		expense := Expense{
			Id:          len(expenses) + 1,
			Date:        time.Now(),
			Description: userInput.Description,
			Amount:      (userInput.Amount),
		}

		expenses = append(expenses, expense)

		data, err := json.Marshal(expenses)
		if err != nil {
			fmt.Println(err)
		}

		os.WriteFile("data.json", data, 0600)

	case "list":

		for _, value := range expenses {
			fmt.Printf("%d %s %f at %d/%d/%d\n", value.Id, value.Description, value.Amount, value.Date.Month(), value.Date.Day(), value.Date.Year())
		}

	}

}
