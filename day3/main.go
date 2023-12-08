package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)


func main() {
    data, err := os.ReadFile("input.txt") 
    if err != nil {
        log.Println(err)
    }

    s := bufio.NewScanner(bytes.NewBuffer(data))
    grid := [][]rune{}
    sum := 0
    for s.Scan() {
        grid = append(grid, []rune(s.Text())) 
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; i < len(grid[i]); j++ {
            if grid[i][j] != '.' && !unicode.IsDigit(grid[i][j]) {
                sum += getNums(grid, i, j)
            }
        }
    }

    println(sum)
}

func getNums(arr [][]rune, i, j int) int {
    
}
