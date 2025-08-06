package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

const dataFile = "expenses.json"

func loadExpenses() ([]Expense, error) {
	var expenses []Expense

	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return expenses, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func saveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling expenses to JSON:", err)
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Expenses saved successfully.")
	return nil
}


func main() {
	expenses, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	idCounter := 1
	if len(expenses) > 0 {
		idCounter = expenses[len(expenses)-1].ID + 1
	}

	var option int
	fmt.Println("Choose an option:\n1 - Add expense\n2 - Delete expense\n3 - View all expenses")
	fmt.Scanln(&option)

	switch option {
	case 1:
		keepGoing := "y"
		for keepGoing == "y" {
			var amount float64
			var category, description string

			fmt.Print("Amount: ")
			fmt.Scanln(&amount)

			fmt.Print("Category: ")
			fmt.Scanln(&category)

			fmt.Print("Description: ")
			fmt.Scanln(&description)

			expense := Expense{
				ID:          idCounter,
				Amount:      amount,
				Category:    category,
				Description: description,
				Date:        time.Now(),
			}
			idCounter++

			expenses = append(expenses, expense)
			fmt.Println("Expense Added!")

			fmt.Print("Add another expense? (y/n): ")
			fmt.Scanln(&keepGoing)
		}

	case 2:
		var deleteID int
		fmt.Print("Enter the ID of the expense to delete: ")
		fmt.Scanln(&deleteID)

		found := false
		for i, e := range expenses {
			if e.ID == deleteID {
				expenses = append(expenses[:i], expenses[i+1:]...)
				fmt.Printf("Expense with ID %d deleted.\n", deleteID)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Expense ID not found.")
		}

	case 3:
		if len(expenses) == 0 {
			fmt.Println("No expenses recorded.")
		} else {
			fmt.Println("All Expenses:")
			for _, e := range expenses {
				fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
					e.ID, e.Amount, e.Category, e.Description, e.Date.Format("2006-01-02 15:04:05"))
			}
		}

	default:
		fmt.Println("Invalid option.")
	}

	err = saveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
	} else {
		fmt.Println("Expenses saved successfully.")
	}
}
