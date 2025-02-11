package main

import (
    "fmt"
    "time"
)

func mcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func mcm(a, b int) int {
    return abs(a*b) / mcd(a, b)
}

func mcmVarios(nums ...int) int {
    resultado := 1
    for _, num := range nums {
        resultado = mcm(resultado, num)
    }
    return resultado
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    startTime := time.Now()
		for i := 0; i < 300000; i++ {
    	mcmVarios(12321, 5674, 123, 821)
		}
    endTime := time.Now()
    fmt.Printf("%v\n", endTime.Sub(startTime).Seconds() * 1000)
}
