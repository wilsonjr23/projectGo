// Echo2 exibe seus argumentos de linha de comando.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for ind, arg := range os.Args[:] {
		fmt.Println(ind, arg)
	}
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
