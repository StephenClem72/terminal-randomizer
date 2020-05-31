package main

import (
	"fmt"
	"strings"
  "os"
	"os/exec"
	"time"
	"math/rand"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	clear_terminal()
	args := os.Args[1:]
	rotations := 50
	pluck_results(args, rotations)
}

func pluck_results(results []string, rotations int){
	for i := 0; i < rotations; i++ {
		index := rand.Int() % len(results)
		value := strings.Title(results[index])
		if i == (rotations - 1) {
			ascii_print(value)
			break
		}
		fmt.Println(value)
		time.Sleep(100 * time.Millisecond)
		clear_terminal()
	}
}

func ascii_print(value string){
	cap_value := strings.ToUpper(value)
	fmt.Println(figure.NewFigure(cap_value, "dotmatrix", true))
}

func clear_terminal(){
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}
