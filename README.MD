# Expense Tracker CLI

A simple command-line application to track expenses using a JSON file as a database. You can add, list, delete, update, and summarize expenses with various command-line arguments.

## Features

- Add expenses with a description and amount
- List all expenses
- Delete expenses by ID
- Update expense details
- Summarize expenses for a given month

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/RoshanShrestha123/expense-tracker-cli-go
   ```

2. Change into the project directory:

   ```sh
   cd expense-tracker-cli-go
   ```

3. Build the project:

   ```sh
   go build -o expense-tracker
   ```

## Usage

Run the executable with the desired command:

### 1. Add an Expense

```sh
./expense-tracker add --description "Lunch" --amount 15.50
```

### 2. List Expenses

```sh
./expense-tracker list
```

### 3. Delete an Expense

```sh
./expense-tracker delete --id 1
```

### 4. Update an Expense

```sh
./expense-tracker update --id 1 --description "Dinner" --amount 20.00
```

### 5. Summarize Expenses

- Get total expenses for all months:

  ```sh
  ./expense-tracker summary
  ```

- Get total expenses for a specific month (e.g., January):

  ```sh
  ./expense-tracker summary --month 1
  ```

## Data Storage

The expenses are stored in a `data.json` file in the same directory as the application.

## Error Handling

- If an invalid command is used, the program will exit with an error message.
- If required arguments are missing, the program will prompt an error.

## Assignment URL
https://roadmap.sh/projects/expense-tracker
