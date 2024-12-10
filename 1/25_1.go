package main

import (
	"fmt"
	"time"
)

// CustomSleep блокирует выполнение программы на заданное количество миллисекунд
func CustomSleep(duration time.Duration) {
	select {
	case <-time.After(duration):
	}
}

func main() {
	fmt.Println("Start")
	CustomSleep(2 * time.Second) // Задержка на 2 секунды
	fmt.Println("End")
}
