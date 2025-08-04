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

func main (){
	var expenses []Expense

	var amount float64
	var category string
	var description string
	
	fmt.Println("Hello, Input your expenses for today?")
	fmt.Println("Enter Amount:")
	fmt.Scanln(&amount)
	fmt.Println("What Category:")
	fmt.Scanln(&category)
	fmt.Println("Desscription:")
	fmt.Scanln(&description)
	
	expense := Expense{
		id: 1,
		amount: amount,
		category: category,
		description: description,
		date: time.Now(),
	}

	expenses = append(expenses, expense)
	fmt.Println("\n✅ Expense Added!")
	fmt.Println("Current Expenses:")
	for _, e := range expenses {
		fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
	e.id,e.amount,e.category,e.description,e.date.Format("2006-01-02 15:04:05"))
	}



	

	


}