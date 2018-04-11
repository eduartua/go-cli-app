package env

import (
	"fmt"
	"os"
	"strings"
)

const path string = "PROVERBS_FILE"

//ListVars list all environment variables.
func ListVars() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

//CheckPath returns the path if the PROVERBS_FILE exists.
func CheckPath() string {
	var p string
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == path {
			p = pair[1]
			return p
		}
	}
	fmt.Println("No existe la variable de entorno PROVERBS_FILE, usando PATH por defecto ./proverbs.txt")
	return "proverbs.txt"
}
