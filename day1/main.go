package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	res, err := solve(data)
	if err != nil {
		panic(err)
	}

	println(res)

}

func solve(buf []byte) (uint64, error) {
	s := bufio.NewScanner(bytes.NewReader(buf))

	res := uint64(0)
	for s.Scan(){
		val, err := proccessInput(s.Bytes())
		if err != nil {
			return 0, err
		}

		res += val
	}

	return res, nil
}

func proccessInput(line []byte) (uint64, error) {
	firstDigit := findFirstDigit(line)
	lastDigit := findLastDigit(line)

	return (uint64(firstDigit-'0') * 10) + uint64(lastDigit-'0'), nil
}

func findFirstDigit(line []byte) byte {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return line[i]
		}
	}

	return '0'
}

func findLastDigit(line []byte) byte {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return line[i]
		}
	}
	
	return '0'
}

















