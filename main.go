package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/footballgoals"
)

func main() {

	fmt.Println("\033[36mЛабораторна робота № 2. Варіант 18\033[0m")
	fmt.Println("\033[33mПоказник: Загальна кількість голів у матчі\033[0m")
	fmt.Println("\033[32mОдиниця вимірювання: голи\033[0m")
	fmt.Println("\033[35mДіапазон: 0-12\033[0m")
	fmt.Println("\033[34mКатегорії: без голів (0), низька (1-2), середня (3-4), висока (від 5)\033[0m")

	var count int
	fmt.Print("\033[32mВведіть кількість елементів (від 5 до 100): \033[0m")
	_, err := fmt.Scan(&count)
	if err != nil || count < 5 || count > 100 {
		fmt.Println("\033[31mПомилка: некоректне значення кількості.\033[0m")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Print("\033[36mВведіть назву турніру/ліги (competition): \033[0m")
	competitionInput, _ := reader.ReadString('\n')
	competition := strings.TrimSpace(competitionInput)
	if competition == "" {
		fmt.Println("\033[31mПомилка: параметр competition не може бути порожнім.\033[0m")
		return
	}

	values, err := footballgoals.Generate(count, competition)
	if err != nil {
		fmt.Println("\033[31mПомилка генератора:\033[0m", err)
		return
	}

	fmt.Println("\033[35mОтримані дані:\033[0m", values)

	if len(values) == 0 {
		fmt.Println("\033[31mЗріз даних порожній.\033[0m")
		return
	}

	min := values[0]
	max := values[0]
	sum := 0

	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}

	average := float64(sum) / float64(len(values))

	fmt.Println("\033[31mМінімум:\033[0m", min)
	fmt.Println("\033[34mМаксимум:\033[0m", max)
	fmt.Println("\033[35mСума:\033[0m", sum)
	fmt.Println("\033[33mСереднє арифметичне:\033[0m", average)

	BezGoliv := 0
	Nyzka := 0
	Serednya := 0
	Vysoka := 0

	for _, v := range values {
		if v == 0 {
			BezGoliv++
		} else if v >= 1 && v <= 2 {
			Nyzka++
		} else if v >= 3 && v <= 4 {
			Serednya++
		} else if v >= 5 {
			Vysoka++
		}
	}

	fmt.Println("\033[1m\033[35mРозподіл за категоріями:\033[0m")
	fmt.Println("\033[36mБез голів (0):\033[0m", BezGoliv)
	fmt.Println("\033[32mНизька результативність (1-2):\033[0m", Nyzka)
	fmt.Println("\033[33mСередня результативність (3-4):\033[0m", Serednya)
	fmt.Println("\033[31mВисока sрезультативність (від 5):\033[0m", Vysoka)

	frequency := make(map[int]int)
	for _, v := range values {
		frequency[v]++
	}

	fmt.Println("\033[1m\033[33mПовторювані значення (через мапу):\033[0m")
	for val, freq := range frequency {
		fmt.Println("\033[36mЗначення\033[0m", val, "\033[36mзустрічається\033[0m", freq, "\033[36mраз(ів)\033[0m")
	}

	var seed int64 = 20
	fmt.Println("\033[1m\033[36mЕксперимент з seed \033[0m")
	res1, _ := footballgoals.GenerateWithSeed(count, competition, seed)
	res2, _ := footballgoals.GenerateWithSeed(count, competition, seed)

	otherCompetition := competition + "_SuperCup"
	res3, _ := footballgoals.GenerateWithSeed(count, otherCompetition, seed)

	fmt.Println("\033[32mВиклик з seed 1: \033[0m", res1)
	fmt.Println("\033[32mВиклик з seed 2: \033[0m", res2)
	fmt.Println("\033[33mВиклик з іншим competition: \033[0m", res3)

	countBig3 := 0

	fmt.Println("Голи, більші за 3:")
	for _, v := range values {
		if v > 3 {
			fmt.Println(v)
			countBig3++
		}
	}
	fmt.Println("К-сть голів які більші за 3:", countBig3)

	percentage := (float64(countBig3) / float64(len(values))) * 100.0
	fmt.Println("Відсоток від усього масиву:", percentage)

	if countBig3 > len(values)/2 {
		fmt.Println("Результат: Кількість таких чисел БІЛЬША за половину масиву.")
	} else {
		fmt.Println("Результат: Кількість таких чисел НЕ БІЛЬША за половину масиву.")
	}
}
