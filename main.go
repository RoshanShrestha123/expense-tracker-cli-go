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

func AutoIncrementId(data *[]Expense) int {
	max := 0

	for _, value := range *data {
		if value.Id > max {
			max = value.Id
		}
	}

	return max + 1
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

	res, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Unable to read the data from file")
	}

	var expenses []Expense

	err = json.Unmarshal(res, &expenses)
	if err != nil {
		log.Fatal(err)
	}

	switch userInput.Action {
	case "add":

		expense := Expense{
			Id:          AutoIncrementId(&expenses),
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

	case "delete":
		var index int
		if userInput.Id == 0 {
			log.Fatal("Please provide the id to be removed")
		}

		for i, value := range expenses {
			if value.Id == userInput.Id {
				index = i
				break
			}
		}

		expenses = append(expenses[:index], expenses[index+1:]...)

		data, err := json.Marshal(expenses)
		if err != nil {
			fmt.Println(err)
		}

		os.WriteFile("data.json", data, 0600)

	case "update":
		var index int
		if userInput.Id == 0 {
			log.Fatal("Please provide the id ")
		}

		for i, value := range expenses {
			if value.Id == userInput.Id {
				index = i
				break
			}
		}

		toBeUpdate := expenses[index]

		toBeUpdate.Description = userInput.Description
		toBeUpdate.Amount = userInput.Amount

		expenses[index] = toBeUpdate

		data, err := json.Marshal(expenses)
		if err != nil {
			fmt.Println(err)
		}

		os.WriteFile("data.json", data, 0600)
	case "summary":
		var total float32

		for _, value := range expenses {

			if userInput.Month == "" {
				total += value.Amount

			}

			userInputMonth, _ := strconv.ParseInt(userInput.Month, 10, 64)
			if int(value.Date.Month()) == int(userInputMonth) {
				total += value.Amount
			}
		}

		fmt.Printf("Total Expenses is %f\n", total)

	}

}
