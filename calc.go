package main

import (
	"./calc"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	env := calc.Env{}
	for _, arg := range os.Args[1:] {
		scanner := new(calc.Scanner)
		body, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		scanner.Init(string(body))
		for _, statement := range calc.Parse(scanner) {
			s, err := calc.Evaluate(statement, env)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(s)
		}
	}
}
