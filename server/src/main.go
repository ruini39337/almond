package main

import (
	"fmt"
	"math"
)

func main(){
	b := 1.999999999999

	//a,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", b), 64)
	b = math.Floor(b)

	fmt.Println(b)
}
