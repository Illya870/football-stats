package analyzer

import (
	"errors"
	"strings"

	"github.com/compilyator/VNTU_GoLang/labs/Lab2Data/footballgoals"
)

type AnalysisResult struct {
	Sum, Min, Max int
	Avg           float64
	Frequencies   map[int]int
}

// Validate перевіряє вхідні параметри
func Validate(count int, competition string) (string, error) {
	competition = strings.TrimSpace(competition)
	if count < 5 || count > 100 {
		return "", errors.New("count must be between 5 and 100")
	}
	if competition == "" || len(competition) > 100 {
		return "", errors.New("competition cannot be empty")
	}
	return competition, nil
}

// RunAnalyzer запускає генерацію та підрахунок
func RunAnalyzer(count int, competition string, seed *int64) (AnalysisResult, []int, error) {
	var values []int
	var err error
	if seed != nil {
		values, err = footballgoals.GenerateWithSeed(count, competition, *seed)
	} else {
		values, err = footballgoals.Generate(count, competition)
	}
	if err != nil || len(values) == 0 {
		return AnalysisResult{}, nil, errors.New("generation error or empty data")
	}

	sum, min, max := 0, values[0], values[0]
	freq := make(map[int]int)

	for _, v := range values {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		freq[v]++
	}

	res := AnalysisResult{
		Sum:         sum,
		Min:         min,
		Max:         max,
		Avg:         float64(sum) / float64(len(values)),
		Frequencies: freq,
	}

	return res, values, nil
}
