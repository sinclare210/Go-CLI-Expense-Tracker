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
	var option int

	fmt.Println("Do you want to add or delete expense(1 for add, 2 for delete, 3 for view all expenses)")
	fmt.Scanln(&option)

	keepGoing := "y"

	switch option {
	case 1:
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
case 2:
	var deleteInt int
	fmt.Println("The id of the expense you want to delete?")
	fmt.Scanln(&deleteInt)

	for i,e := range expenses {
		if e.id == deleteInt{
			expenses = append(expenses[:i], expenses[i+1:]...)
			fmt.Printf("Expense with ID %d deleted.\n", deleteInt)
			break
		}
	}

	case 3:
		for _, e := range expenses {
	fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
			e.id, e.amount, e.category, e.description, e.date.Format("2006-01-02 15:04:05"))
	}
		
	}

	
for _, e := range expenses {
	fmt.Printf("ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
			e.id, e.amount, e.category, e.description, e.date.Format("2006-01-02 15:04:05"))
	}



}

	
