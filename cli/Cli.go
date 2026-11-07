package cli

import (
	"flag"
	"fmt"
	"os"

	"football-stats/analyzer"
)

func Cli(Args []string) {

	subcommand := os.Args[1]

	if subcommand == "help" {
		PrintHelp()
	} else if subcommand == "version" {
		fmt.Println("football-stats v1.0.0 (Lab 3)")
	} else if subcommand == "analyze" {
		fs := flag.NewFlagSet("analyze", flag.ExitOnError)
		count := fs.Int("count", 0, "Кількість значень (5-100)")
		competition := fs.String("competition", "", "Назва турніру")
		seed := fs.Int64("seed", -1, "Seed (необов'язково)")

		fs.Parse(os.Args[2:])

		validComp, err := analyzer.Validate(*count, *competition)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Помилка валідації:", err)
			os.Exit(2)
		}

		var seedPtr *int64
		if *seed != -1 {
			seedPtr = seed
		}

		res, vals, err := analyzer.RunAnalyzer(*count, validComp, seedPtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Помилка:", err)
			os.Exit(1)
		}

		fmt.Println("Дані:", vals)
		fmt.Println("Сума:", res.Sum)
		fmt.Println("Мінімум:", res.Min)
		fmt.Println("Максимум:", res.Max)
		fmt.Printf("Середнє: %.2f\n", res.Avg)

		fmt.Println("\033[1m\033[33mПовторювані значення (через мапу):\033[0m")
		for val, freq := range res.Frequencies {
			fmt.Println("\033[36mЗначення\033[0m", val, "\033[36mзустрічається\033[0m", freq, "\033[36mраз(ів)\033[0m")
		}
	} else {
		PrintHelp()
		os.Exit(2)
	}
}

func PrintHelp() {
	fmt.Println("Використання: go run main.go [analyze|help|version]")
	fmt.Println("  analyze --count <5-100> --competition <назва> [--seed <число>]")
}
