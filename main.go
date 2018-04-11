package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eduartua/workshop-go-cli/solucion/indsearch"
)

func main() {
	prov := `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`
	proverbios := strings.Split(prov, "\n")
	limite := len(proverbios)
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
			fmt.Printf("Proverbio #%d\t--> %s\n", i, proverbios[i-1])
			return
		} else if i > limite || i == 0 {
			fmt.Printf("El número de proverbio debe estar enre 1 y %d ó entre -1 y -%d.", limite, limite)
			return
		}
	} else {
		flag.Parse()
		indexes := indsearch.Search(*palabraPtr, proverbios)
		for _, v := range indexes {
			fmt.Printf("Proverbio #%d\t%s\n", v+1, proverbios[v])
		}
		return
	}
}
