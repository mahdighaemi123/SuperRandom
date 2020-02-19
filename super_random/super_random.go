package super_random

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var randomCount int64= 0
var a = []float64{1,3,9,12,18,123,222,223,4545,4646,58585,58586,606060,606062,7777777,88884444,88888888,99933399,999999999}

func getMax(input int)(max int){
	strInput :=toString(input)
	max = int(math.Pow(10, float64(len(strInput))) - 1)
	return
}

// Random get Random value Between min and max With 1 algorithm
func Random(min int, max int) float64 {

	value := max - min

	rand.Seed(time.Now().UTC().UnixNano())
	randomValue := rand.Float64() * float64(value)

	return float64(min) + randomValue
}

// SuperRandom get Random value Between min and max With 5 algorithm
func SuperRandom(min int ,max int ) (result int) {
	value := max - min
	algorithm := make([]float64,0)
	sum := float64(0)

	randomElementPosition := int(Random(0,len(a)-1))
	randomElementPosition2 := int(Random(randomElementPosition,len(a)-1))

	algorithm = append(algorithm, float64(randomCount) / float64(getMax(int(randomCount))) )
	algorithm = append(algorithm, Random(0,1))
	algorithm = append(algorithm, toFloat(getStrCutOfTime()) / float64(getMax(int(toFloat(getStrCutOfTime())))))
	algorithm = append(algorithm, a[randomElementPosition]/float64(getMax(int(a[randomElementPosition]))))
	algorithm = append(algorithm, a[randomElementPosition]/a[randomElementPosition2])

	for i:=0;i<len(algorithm);i++{
		sum+=algorithm[i]
		//fmt.Printf("%.2f \n", algorithm[i])
	}

	sum /= float64(len(algorithm))
	result = min + int(sum * float64(value))

	randomCount++

	return
}

func getStrCutOfTime() string {
	strTime:= toString(getTime())
	return strTime[len(strTime)-3:]
}

func getTime() int {
	return int(time.Now().UnixNano() / 1000)
}

func toString(intValue int) string {
	return strconv.Itoa(intValue)
}

func toFloat(myString string) float64 {
	value, err := strconv.ParseFloat(myString, 64)
	if err != nil {
		fmt.Println("Error", "toFloat", err)
	}
	return value
}