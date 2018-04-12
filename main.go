package main

import (
	"flag"
	"fmt"
	"io"
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
	limite := len(ps) - 1
	if i != 0 && i <= limite {
		fmt.Printf("Proverbio #%d\t--> %s\n", i, ps[i-1])
	} else if i > limite || i == 0 {
		fmt.Printf("El número de proverbio debe estar enre 1 y %d ó entre -1 y -%d.", limite, limite)
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

//Config configs the writing to a file
type Config struct{ outfile io.Writer }

func main() {
	path := env.CheckPath()
	proverbs, err := ioutil.ReadFile(path)
	check(err)
	ps := strings.Split(string(proverbs), "\n")
	outPtr := flag.String("o", "path", "PATH al cuál debería ser escrito el archivo")
	findPtr := flag.String("f", "find", "palabra de búsqueda")
	flag.Parse()

	//cfg := Config{outfile: os.Stdout}

	if len(os.Args) < 2 {
		return
	}
	if os.Args[1] == "env" {
		env.ListVars()
		return
	}
	if os.Args[1] == "list" {
		if checkFlag(os.Args[2], os.Args) {
			f, err := os.Create(*outPtr)
			check(err)
			defer f.Close()
			for _, v := range ps {
				_, err = f.Write([]byte(v + "\n"))
				check(err)
			}
			return
		}
		search.PrintAll(ps)
		return
	}
	if len(os.Args) != 2 || (os.Args[1] == "-f") {
		fmt.Println(os.Args)
		//flag.Parse()
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
