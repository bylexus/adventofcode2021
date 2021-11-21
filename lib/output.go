package lib

import (
	"fmt"
	"time"
)

func AnsiBold(input interface{}) string {
	return fmt.Sprintf("\033[31;1;4m%v\033[0m", input)
}

func Highlight(input interface{}) string {
	return fmt.Sprintf("\033[1;4m%v\033[0m", input)
}

func OutputTitle(input string) {
	outStr := fmt.Sprintf("%v", input)
	fmt.Printf("\033[1;4m%v\033[0m\n\n", outStr)
}

func OutputSolution(nr int, duration time.Duration, solution interface{}) {
	fmt.Printf("Solution %v: %v\n", nr, Highlight(solution))
	fmt.Printf("Took: %v ms\n\n", float64(duration.Nanoseconds())/1000.0/1000.0)
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Since(start)
	return end
}
