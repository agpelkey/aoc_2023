package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
    red = 12
    green = 13
    blue = 14
)


func main() {
    data, err := os.ReadFile("input.txt") 
    if err != nil {
        log.Fatal("error opening file")
    }


    fmt.Println(solve(data))
}


func solve(buf []byte) int {
    /*
    s := bufio.NewScanner(bytes.NewReader(buf))

    for s.Scan() {
        line := s.Text()
        processLine(line)
    }
    return 0, nil
    */

    sum := 0

    s := bufio.NewScanner(bytes.NewReader(buf))

    for s.Scan() {
        id, isOk := processLine(s.Text())

        if isOk {
            sum += id
        }
    }

    return sum

}


func processLine(line string) (int, bool) {

    isValid := true

    parse := strings.Split(line, ":")
    gameInt := strings.Fields(parse[0])
    gameId, err := strconv.Atoi(gameInt[1])
    if err != nil {
        log.Fatal("error getting game id")
    }

    draws := strings.Split(parse[1], ";")
    for _, val := range draws {
        balls := strings.Split(val, ",")
        for _, ball := range balls {
            ballSplit := strings.Split(strings.TrimSpace(ball), " ")
            ballCount, err := strconv.Atoi(ballSplit[0])
            ballColor := ballSplit[1]

            if err != nil {
                log.Fatal(err)
            }

            switch ballColor {
            case "red":
                isValid = isValid && ballCount <= red
            case "green":
                isValid = isValid && ballCount <= green
            case "blue":
                isValid = isValid && ballCount <= blue
            }

        }
    }

    return gameId, isValid
}

/*
99
 9 green
 7 blue
 1 green
 3 red
 4 blue
*/








