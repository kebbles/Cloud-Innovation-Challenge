package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Result represents a customer
type Result struct {
	Name string
	Age  int
}

type information struct {
	Gender          string
	Age             string
	MortgageBalance string
	MonthlyPayment  string
	YearsRemaining  int
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	p := Result{Name: "Hasan Ayyad", Age: 27}
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, p)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	gender := r.PostFormValue("gender")
	age := r.PostFormValue("age")
	mortgageBalance := r.PostFormValue("mortgageBalance")
	monthlyPayment := r.PostFormValue("monthlyPayment")
	mortgageBalanceInt, _ := strconv.Atoi(mortgageBalance)
	monthlyPaymentInt, _ := strconv.Atoi(monthlyPayment)
	yearsRemaining := (mortgageBalanceInt / monthlyPaymentInt) / 12

	person := information{
		Gender:          gender,
		Age:             age,
		MortgageBalance: mortgageBalance,
		MonthlyPayment:  monthlyPayment,
		YearsRemaining:  yearsRemaining,
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, person)
}

func main() {
	// Handling static assets
	files := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// Routing requests
	http.HandleFunc("/form/", formHandler)
	http.HandleFunc("/results/", resultsHandler)

	// Starting server
	http.ListenAndServe(":8080", nil)
}
