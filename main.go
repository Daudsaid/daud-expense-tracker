package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
	Date     string  `json:"date"` // "YYYY-MM-DD"
}

var expenses []Expense

const dataFile = "expenses.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	loadExpenses()

	for {
		fmt.Println("============== Expense Tracker ==============")
		fmt.Println("1) Add expense")
		fmt.Println("2) List all expenses")
		fmt.Println("3) Show total spent")
		fmt.Println("4) Show total per category")
		fmt.Println("5) List expenses by category")
		fmt.Println("6) Save & Exit")
		fmt.Println("7) Delete an expense")
		fmt.Println("8) Edit an expense")
		fmt.Print("Choose an option (1-8): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addExpense(reader)
		case "2":
			listExpenses()
		case "3":
			showTotal()
		case "4":
			showTotalPerCategory()
		case "5":
			listByCategory(reader)
		case "6":
			saveExpenses()
			fmt.Println("Data saved. Goodbye!")
			return
		case "7":
			deleteExpense(reader)
		case "8":
			editExpense(reader)
		default:
			fmt.Println("Invalid option, try again.")
		}

		fmt.Println()
	}
}

func addExpense(reader *bufio.Reader) {
	fmt.Print("Amount (e.g. 12.50): ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount, please use numbers only (e.g. 12.50).")
		return
	}

	fmt.Print("Category (e.g. Food, Transport, Rent): ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Note (optional, press Enter to skip): ")
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)

	fmt.Print("Date (YYYY-MM-DD, press Enter for today): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		dateStr = time.Now().Format("2006-01-02")
	}

	expense := Expense{
		Amount:   amount,
		Category: category,
		Note:     note,
		Date:     dateStr,
	}

	expenses = append(expenses, expense)
	fmt.Println("✅ Expense added!")
}

func listExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("------------- Your Expenses -------------")
	for i, e := range expenses {
		fmt.Printf("%d) %s - £%.2f [%s]", i+1, e.Date, e.Amount, e.Category)
		if e.Note != "" {
			fmt.Printf(" - %s", e.Note)
		}
		fmt.Println()
	}
}

func showTotal() {
	if len(expenses) == 0 {
		fmt.Println("No expenses yet, total is £0.00")
		return
	}

	var total float64
	for _, e := range expenses {
		total += e.Amount
	}

	fmt.Printf("💰 Total spent: £%.2f\n", total)
}

// showTotalPerCategory sums amounts for each category and prints them.
func showTotalPerCategory() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	totals := make(map[string]float64)

	for _, e := range expenses {
		totals[e.Category] += e.Amount
	}

	fmt.Println("----- Total per Category -----")
	for category, total := range totals {
		fmt.Printf("%s: £%.2f\n", category, total)
	}
}

// listByCategory prints expenses for a specific category.
func listByCategory(reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Print("Enter category to filter by: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Printf("Expenses in category: %s\n", category)

	found := false
	for i, e := range expenses {
		if strings.EqualFold(e.Category, category) {
			found = true
			fmt.Printf("%d) %s - £%.2f", i+1, e.Date, e.Amount)
			if e.Note != "" {
				fmt.Printf(" - %s", e.Note)
			}
			fmt.Println()
		}
	}

	if !found {
		fmt.Println("No expenses found in this category.")
	}
}

// deleteExpense removes an expense by its displayed number.
func deleteExpense(reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("No expenses to delete.")
		return
	}

	listExpenses()

	fmt.Print("Enter the number of the expense to delete: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)

	index, err := strconv.Atoi(numStr)
	if err != nil || index < 1 || index > len(expenses) {
		fmt.Println("Invalid number.")
		return
	}

	// Convert to zero-based index
	index--

	// Remove the item at position index
	expenses = append(expenses[:index], expenses[index+1:]...)

	fmt.Println("🗑 Expense deleted.")
}

// editExpense allows editing fields of an existing expense.
func editExpense(reader *bufio.Reader) {
	if len(expenses) == 0 {
		fmt.Println("No expenses to edit.")
		return
	}

	listExpenses()

	fmt.Print("Enter the number of the expense to edit: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)

	index, err := strconv.Atoi(numStr)
	if err != nil || index < 1 || index > len(expenses) {
		fmt.Println("Invalid number.")
		return
	}

	index-- // zero-based

	e := expenses[index]
	fmt.Printf("Editing expense %d: %s - £%.2f [%s]", index+1, e.Date, e.Amount, e.Category)
	if e.Note != "" {
		fmt.Printf(" - %s", e.Note)
	}
	fmt.Println()

	// For each field: show current value, allow Enter to keep it.

	fmt.Printf("New amount (current: £%.2f, Enter to keep): ", e.Amount)
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	if amountStr != "" {
		newAmount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Invalid amount, keeping old value.")
		} else {
			e.Amount = newAmount
		}
	}

	fmt.Printf("New category (current: %s, Enter to keep): ", e.Category)
	categoryStr, _ := reader.ReadString('\n')
	categoryStr = strings.TrimSpace(categoryStr)
	if categoryStr != "" {
		e.Category = categoryStr
	}

	fmt.Printf("New note (current: %s, Enter to keep): ", e.Note)
	noteStr, _ := reader.ReadString('\n')
	noteStr = strings.TrimSpace(noteStr)
	if noteStr != "" {
		e.Note = noteStr
	}

	fmt.Printf("New date (current: %s, Enter to keep): ", e.Date)
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)
	if dateStr != "" {
		e.Date = dateStr
	}

	// Save back to slice
	expenses[index] = e
	fmt.Println("✏️ Expense updated.")
}

func saveExpenses() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(expenses); err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}

func loadExpenses() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening data file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&expenses); err != nil {
		fmt.Println("Error reading JSON:", err)
	}
}
