package main

import (
	"bufio"
	"fmt"
	"os"

	"duration2"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		d, err := duration2.Parse(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		fmt.Printf("%s = %s (%d ns)\n", line, d, d.Nanoseconds())
	}
}
