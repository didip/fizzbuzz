// package main
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func buffered() {
	containerSize := 4500
	container := make([]string, containerSize)
	i := 1

	for {
		if i%15 == 0 {
			container[i%containerSize] = "FizzBuzz"
		} else if i%5 == 0 {
			container[i%containerSize] = "Buzz"
		} else if i%3 == 0 {
			container[i%containerSize] = "Fizz"
		} else {
			container[i%containerSize] = strconv.Itoa(i)
		}

		if container[containerSize-1] != "" {
			fmt.Println(strings.Join(container, "\n"))
		}

		i++
	}
}

func naive() {
	i := 1

	for {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(strconv.Itoa(i))
		}

		i++
	}
}

func main() {
	modePtr := flag.String("mode", "naive", "which mode?")
	flag.Parse()

	switch *modePtr {
	case "naive":
		naive()
	case "buffered":
		buffered()
	}
}
