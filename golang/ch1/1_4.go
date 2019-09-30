// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            counts[input.Text()]["STDIN"]++
        }
    } else{
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            input := bufio.NewScanner(f)
            for input.Scan() {
                text := input.Text()
                if counts[text] == nil {
                    counts[text] = make(map[string]int)
                }
                counts[text][arg]++
            }
            f.Close()
        }
    }
    for line, file := range counts {
        lineCount := 0
        files := ""
        for filename, n := range file {
            lineCount += n
            files += " " + filename
        }
        if lineCount > 1 {
            fmt.Printf("%d\t%s\t%s\n", lineCount, line, files)
        }
    }
}
