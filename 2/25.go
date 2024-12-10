package main

import (
	"fmt"
	"time"
)

// CustomSleepUsingLoop выполняет задержку через активное ожидание
func CustomSleepUsingLoop(duration time.Duration) {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {
		// Пустой цикл
	}
}

func main() {
	fmt.Println("Start")
	CustomSleepUsingLoop(2 * time.Second) // Задержка на 2 секунды
	fmt.Println("End")
}
