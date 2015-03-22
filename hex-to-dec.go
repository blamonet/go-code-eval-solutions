package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        hex := scanner.Text()
        i, err := strconv.ParseInt(hex, 16, 64)
        
        if err == nil {
        	fmt.Println(i)
        } else {
        	fmt.Println(err)
        }
    }   
}