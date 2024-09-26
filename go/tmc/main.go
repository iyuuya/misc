package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	result, err := calc(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	minutes := int(result.Minutes())
	seconds := int(result.Seconds()) % 60
	fmt.Printf("%02d:%02d\n", minutes, seconds)
}

func calc(args []string) (time.Duration, error) {
	if len(args) < 3 {
		return 0, errors.New("not enough arguments")
	}

	var result time.Duration
	for i, arg := range args {
		if i%2 == 1 {
			if arg != "+" && arg != "-" {
				return 0, errors.New("unsupported operator")
			}
			continue
		}

		formats := []string{
			"04:05", // 4分5秒
			"04:5",  // 4分5秒
			"4:05",  // 4分5秒
			"4:5",   // 4分5秒
			"5",     // 5秒
		}
		var t time.Time
		var err error
		for _, format := range formats {
			t, err = time.Parse(format, arg)
			if err == nil {
				break
			}
		}
		if err != nil {
			return 0, err
		}

		duration := time.Duration(t.Minute())*time.Minute +
			time.Duration(t.Second())*time.Second +
			time.Duration(t.Nanosecond())*time.Nanosecond

		if i%2 == 0 || args[i-1] == "+" {
			result += duration
		} else if args[i-1] == "-" {
			result -= duration
		}
	}

	return result, nil
}
