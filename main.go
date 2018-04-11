package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/eduartua/workshop-go-cli/solucion/env"
	"github.com/eduartua/workshop-go-cli/solucion/indsearch"
)

func check(e error) {
	if e != nil {
		fmt.Println("No existe la variable de entorno PROVERBS_FILE")
		panic(e)
	}
}

func printProv(i int, ps []string) {
	limite := len(ps) - 1
	if i != 0 && i <= limite {
		fmt.Printf("Proverbio #%d\t--> %s\n", i, ps[i-1])
	} else if i > limite || i == 0 {
		fmt.Printf("El número de proverbio debe estar enre 1 y %d ó entre -1 y -%d.", limite, limite)
	}
	return
}

func main() {
	path := env.CheckPath()
	proverbs, err := ioutil.ReadFile(path)
	check(err)
	ps := strings.Split(string(proverbs), "\n")
	palabraPtr := flag.String("f", "", "palabra de búsqueda")

	if len(os.Args) < 2 {
		return
	} else if len(os.Args) == 2 && (os.Args[1] != "-f") {
		if os.Args[1] == "env" {
			env.ListVars()
			return
		}

		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argumento no numérico, o flag desconocida (usar -f).")
			return
		} else if i < 0 {
			i = -i
		}
		printProv(i, ps)
	} else {
		flag.Parse()
		indexes := indsearch.Search(*palabraPtr, ps)
		for _, v := range indexes {
			fmt.Printf("Proverbio #%d\t%s\n", v+1, ps[v])
		}
		return
	}
}
