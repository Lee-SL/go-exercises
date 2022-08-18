package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filteredRecords []Record
	for _, r := range in {
		if predicate(r) {
			filteredRecords = append(filteredRecords, r)
		}
	}

	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day <= p.To && r.Day >= p.From {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecords := Filter(in, ByDaysPeriod(p))

	totalExpense := 0.0
	for _, r := range filteredRecords {
		totalExpense += r.Amount
	}

	return totalExpense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filterByCategory := Filter(in, ByCategory(c))
	if len(filterByCategory) <= 0 {
		return 0.0, errors.New("No records by this category")
	}

	filterByDay := Filter(in, ByDaysPeriod(p))
	totalExpense := 0.0
	for _, r := range filterByDay {
		if r.Category == c {
			totalExpense += r.Amount
		}
	}

	return totalExpense, nil
}
