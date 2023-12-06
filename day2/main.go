package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)


func main() {
    data, err := os.ReadFile("input.txt") 
    if err != nil {
        log.Fatal(err)
    }


    solve(data)
}


func solve(buf []byte) (uint64, error) {
    s := bufio.NewScanner(bytes.NewReader(buf))

    fmt.Print(s)
    for s.Scan() {
        //exodia := make(map[string]string)

    }
    return 0, nil
}
