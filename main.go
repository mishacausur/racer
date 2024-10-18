package main

import (
	"fmt"
	"time"
)

func main() {
	time.NewTimer(5 * time.Second) // Создаётся таймер  на 5 секунд
	duration := 5 * time.Second
	fmt.Println("Duration:", duration)

	// Выполняем действия в течение заданного промежутка времени
	fmt.Println("Start")
	time.Sleep(duration)
	fmt.Println("End")

	// Вычисляем разницу между моментами времени
	t1 := time.Now()
	time.Sleep(duration)
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Difference:", diff)
}
