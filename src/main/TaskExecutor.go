package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main(){
	fileHandler := &FileHandler{SharedFilePath: "arquivo.txt"}

	fileHandler.CreateFile()

	var n int
	var e int
	var t int

	fmt.Print("Digite a potência para o número de tarefas:  ")
  	fmt.Scan(&n)
    fmt.Println("Digite o percentual de tarefas de escrita: ")
	fmt.Scan(&e)
	fmt.Println("Digite o número de threads a serem criadas: ")
	fmt.Scan(&t)

	var allTasks []Task

	for i := 0; i < int((float64(E)/100.0)*math.Pow(10, float64(n))); i++ {
		cost := rand.Float64() / 100.0
		value := rand.Intn(10)
		allTasks = append(allTasks, NewTask(i, cost, false, value))
	}

	for i := int((float64(E)/100.0)*math.Pow(10, float64(n))); i < math.Pow(10, float64(n)); i++ {
		cost := rand.Float64() / 100.0
		value := rand.Intn(10)
		allTasks = append(allTasks, NewTask(i, cost, true, value))
	}

	rand.Shuffle(len(allTasks), func(i, j int) {
		allTasks[i], allTasks[j] = allTasks[j], allTasks[i]
	})
}
