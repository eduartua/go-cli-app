package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/eduartua/workshop-go-cli/solucion/indsearch"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	proverbs, err := ioutil.ReadFile("./proverbs.txt")
	check(err)
	ps := strings.Split(string(proverbs), "\n")
	limite := len(ps) - 1
	palabraPtr := flag.String("f", "", "palabra de búsqueda")

	if len(os.Args) < 2 {
		return
	} else if len(os.Args) == 2 && (os.Args[1] != "-f") {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argumento no numérico.")
			return
		} else if i < 0 {
			i = -i
		}

		if i != 0 && i <= limite {
			fmt.Printf("Proverbio #%d\t--> %s\n", i, ps[i-1])
			return
		} else if i > limite || i == 0 {
			fmt.Printf("El número de proverbio debe estar enre 1 y %d ó entre -1 y -%d.", limite, limite)
			return
		}
	} else {
		flag.Parse()
		if os.Args[2] == "env" {
			for _, e := range os.Environ() {
				pair := strings.Split(e, "=")
				fmt.Println(pair[0])
			}
			return
		}
		indexes := indsearch.Search(*palabraPtr, ps)
		for _, v := range indexes {
			fmt.Printf("Proverbio #%d\t%s\n", v+1, ps[v])
		}
		return
	}
}
