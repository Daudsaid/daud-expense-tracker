📒 Expense Tracker (Go CLI)

A simple, beginner-friendly command-line Expense Tracker written in Go.
You can add expenses, list them, filter by category, view totals, and save everything to a JSON file.

⸻

⭐ Features
	•	➕ Add expenses with:
	•	Amount
	•	Category
	•	Note
	•	Date (auto-fills today if left empty)
	•	📄 List all expenses
	•	🧮 Show total spent
	•	📊 Show total spent per category
	•	🔍 Filter expenses by category
	•	💾 Automatically saves all data to expenses.json
	•	📂 Data loads automatically when the program starts

🚀 How to Run

Clone the project:
git clone https://github.com/Daudsaid/daud-expense-tracker
cd daud-expense-tracker

Run the program:
go run main.go
That’s it!

📁 File Structure

daud-expense_tracker/
│
├── main.go
├── expenses.json
└── README.md

📦 JSON Data Format

All expenses are saved in expenses.json like this:

[
  {
    "amount": 12.5,
    "category": "Food",
    "note": "Pret coffee",
    "date": "2025-11-16"
  }
]

🛠 Technologies Used
	•	Go (Golang)
	•	Standard library only — no external packages
	•	JSON file persistence


  🧑‍💻 Future Improvements (Optional)

You can add these later if you want:
	•	Export to CSV
	•	Support multi-currency
	•	Add monthly breakdown
	•	Add delete/edit expense
	•	Build a TUI (Terminal UI) using BubbleTea
	•	Make a Go web API version using Gin

⸻

👤 Author

Daud Abdi
GitHub: https://github.com/Daudsaid
www.linkedin.com/in/daudabdi0506






