package finance

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "expenses.txt"

type Expense struct {
	Title  string
	Amount float64
}

func ReadAll() ([]Expense, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []Expense{}, err
	}
	lines := strings.Split(string(data), "\n")

	res := []Expense{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		expenses := strings.Split(line, ";")
		sum, _ := strconv.ParseFloat(expenses[1], 64)
		res = append(res, Expense{Title: expenses[0], Amount: sum})

	}
	return res, nil
}

func PrintAll() {
	data, err := ReadAll()

	if err != nil {
		fmt.Println("Ошибка")
	}

	for i, val := range data {
		fmt.Printf("%d) %s -- %.2f\n", i, val.Title, val.Amount)
	}
}

func Add(title string, amount float64) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}
	str := title + ";" + fmt.Sprintf("%.2f", amount)

	var newStr string
	
	read := strings.TrimSpace(string(data))
	if read == "" {
		newStr = str
	} else {
		newStr = read + "\n" + str
	}
	

	return os.WriteFile(filename, []byte(newStr), 0644)
}

func Delete(index int) error {
	items, err := ReadAll()

	if err != nil {
		return err
	}

	if index < 0 || index >= len(items) {
		return fmt.Errorf("Некоректный индекс %d", index)
	}
	 
  items = append(items[:index], items[index+1:]...)
  lines := ""
 
  for _, val := range items {
	lines += fmt.Sprintf("%s;%.2f\n", val.Title, val.Amount)
  }
 
	return os.WriteFile(filename, []byte(lines), 0644)
}

func Total() (float64, error) {
	data, err := ReadAll()

	if err != nil {
		return 0, fmt.Errorf("Ошибка")
	}
	total := 0.0

	for _, val := range data {
		total += val.Amount
	}
	return total, nil
}
