package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/eduartua/workshop-go-cli/solucion/env"
	"github.com/eduartua/workshop-go-cli/solucion/search"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printProv(i int, ps []string) {
	limite := len(ps)
	if i != 0 && i <= limite {
		fmt.Printf("#%d\t--> %s\n", i, ps[i-1])
	} else if i > limite || i == 0 {
		fmt.Printf("El número debe estar enre 1 y %d.", limite)
	}
	return
}

func checkFlag(s string, xs []string) bool {
	for _, v := range xs {
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	path := env.CheckPath()
	proverbs, err := ioutil.ReadFile(path)
	check(err)
	ps := strings.Split(string(proverbs), "\n")
	findPtr := flag.String("f", "find", "palabra de búsqueda")
	outPtr := flag.String("o", "path", "PATH al cuál debería ser escrito el archivo")
	flag.Parse()

	if checkFlag("env", os.Args) {
		env.ListVars()
		return
	}
	if checkFlag("-o", os.Args) {
		f, err := os.Create(*outPtr)
		check(err)
		defer f.Close()
		if checkFlag("list", os.Args) {
			for _, v := range ps {
				_, err = f.Write([]byte(v + "\n"))
				check(err)
			}
			return
		}
		if checkFlag("-f", os.Args) {
			indexes := search.Search(*findPtr, ps)
			for _, v := range indexes {
				_, err = f.Write([]byte(ps[v] + "\n"))
			}
			return
		}
		return
	}
	if checkFlag("list", os.Args) {
		search.PrintAll(ps)
		return
	}
	if checkFlag("-f", os.Args) {
		indexes := search.Search(*findPtr, ps)
		for _, v := range indexes {
			fmt.Printf("Proverbio #%d\t%s\n", v+1, ps[v])
		}
		return
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argumento no numérico, o flag desconocida (usar -f).")
		return
	}
	if i < 0 {
		i = -i
	}
	printProv(i, ps)
	return
}
