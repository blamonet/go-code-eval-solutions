package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
	"sort"
	"strconv"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        var numberfloatarr []float64
        fmt.Println(scanner.Text())
        numbers := strings.Split(scanner.Text(), " ")

        for _, value := range numbers {
            var n, _ = strconv.ParseFloat(value, 64)
            numberfloatarr = append(numberfloatarr, n)
        }

        sort.Float64s(numberfloatarr[:])

        fmt.Println(strings.Trim(fmt.Sprintf("%.3f", numberfloatarr), "[]"))
        fmt.Println()
    }   
}
