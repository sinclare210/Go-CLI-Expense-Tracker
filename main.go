package main

import (
	"fmt"
	"time"
)

type Expense struct{
	id int
	amount float64
	category string
	description string
	date time.Time
}

func main(){
	var  expenses []Expense
	var amount float64
	var description string
	var category string
	idCounter := 1

	keepGoing := "y"
	for keepGoing == "y"{
	fmt.Println("Hello, What are you expenses for today?")
	fmt.Println("Amount:")
	fmt.Scanln(&amount)
	fmt.Println("category")
	fmt.Scanln(&category)
	fmt.Println("description")
	fmt.Scanln(&description)

	expense := Expense{
		id: idCounter,
		amount: amount,
		category: category,
		description: description,
		date: time.Now(),

	}
	idCounter++
	expenses = append(expenses, expense)
	fmt.Println("Expense Added")
	fmt.Println("Do you want to add another expenes? (y/n)")
	fmt.Scanln(&keepGoing)

	

}

for _, e := range expenses {
	fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
			e.id, e.amount, e.category, e.description, e.date.Format("2006-01-02 15:04:05"))
	}
}

	
