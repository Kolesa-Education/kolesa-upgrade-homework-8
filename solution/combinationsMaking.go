package solution

func makeCombinations(cardsSlice []string)[][]string{
	var length := len(cardsSlice)
	var res := make([][]string, length)
	for i:=0;i<length-1;i++ {
		var tmp := make([]string, length-i)
		for j:=i+1;length;j++{
			tmp[i] = cardsSlice[i]
		}
		res[i] = tmp
	}
	return res
}

/*
a b c d

ab ac ad
bc bd
cd

*/
