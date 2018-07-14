package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	printDate(t)

	t = time.Date(t.Year(), t.Month(), 13, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
	printDate(t)

	for i := 0; i < 3*12; i++ {
		t = t.AddDate(0, 1, 0)
		if t.Weekday() == 5 {
			printDate(t)
		}
	}
}

func printDate(t time.Time) {
	fmt.Print(t)
	fmt.Println(t.Weekday())
}
