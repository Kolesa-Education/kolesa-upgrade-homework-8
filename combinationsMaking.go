package main

import (
    "strconv"
    "strings"
    "math"
    )

func getFactorial(n int) int {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func makeCombinations(cardsSlice []string) string {
	length := len(cardsSlice)
	res := ""
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				for l := k + 1; l < length; l++ {
					for m := l + 1; m < length; m++ {
						tmpRes := cardsSlice[i] + "," + cardsSlice[j] + "," + cardsSlice[k] + "," + cardsSlice[l] + "," + cardsSlice[m]
						res += tmpRes + ";"
					}
				}
			}
		}

	}
	//fmt.Println(res)
	return res
}


func getBinaryCombinations(n int, k int)[]string {
    resLength := getFactorial(n) / (getFactorial(k) * getFactorial(n-k))
    res := make([]string, resLength)
    operationLength := int(math.Pow(2, float64(n)))
    cnt := 0
    for i:=0;i<operationLength;i++{
        binaryValue := strconv.FormatInt(int64(i), 2)
        if strings.Count(binaryValue, "1") != k{
            continue
        }
        for len(binaryValue) < n{
            binaryValue = "0"+binaryValue
        }
        res[cnt] = binaryValue
        cnt++
    }
    return res
}

func getStringFromMask(str string, mask string)string{
    var res string = ""
    for i:=0; i< len(mask); i++ {
        if mask[i] == '1' {
            res += string(str[i])
        }
    }
    return res
}

func makeCombinationsFromMasks(inputStr string)[]string {
    n := len(inputStr)
	combinations := getBinaryCombinations(n, 5)
	res := make([]string, len(combinations))
	cnt := 0
    for i:=0; i< len(combinations); i++ {
        res[cnt] = getStringFromMask(inputStr, combinations[i])
        cnt++
    }
    return res
}

