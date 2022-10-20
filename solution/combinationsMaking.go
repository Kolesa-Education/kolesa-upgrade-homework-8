package solution

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
		//res[i] = make([]string, length)

		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				for l := k + 1; l < length; l++ {
					cnt := 0
					//tmp := make([]string, length)
					for m := l + 1; m < length; m++ {
						tmpRes := cardsSlice[i] + "," + cardsSlice[j] + "," + cardsSlice[k] + "," + cardsSlice[l] + "," + cardsSlice[m]
						res += tmpRes + ";"
						//tmp[cnt] = tmpRes
						cnt++
						//fmt.Println(res)
					}
					//res[i] = tmp
				}
			}
		}

	}
	fmt.Println(res)
	return res
}