package main

import "fmt"
import "time"

func main() {
	// 現在日取得
	t := time.Now()
	printDate(t)

	if t.Day() != 13 {
		// 一度、当月の13日に変換する
		t = time.Date(t.Year(), t.Month(), 13, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
		printDate(t)
	}

	// 指定年数分の日付検出
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
