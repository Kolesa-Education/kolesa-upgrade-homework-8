package main

import "fmt"

func getFactorial(n int)int{
    var res int = 1
    for i:=1;i<=n;i++{
        res *= i
    }
    return res
}

func makeCombinations(cardsSlice []string)[][]string{
	length := len(cardsSlice)
	res := make([][]string, length)
	tmp := make([]string, length)
	a := 0
	for i:=0;i<length-1;i++ {
		for j:=i+1;j<length;j++ {
			for k:=j+1;k<length;k++ {
			    for l:=k+1;l<length;l++ {
        			cnt := 0
        			for m:=l+1;m<length;m++ {
        			    tmpRes := cardsSlice[i] + cardsSlice[j] +cardsSlice[k] + cardsSlice[l] +cardsSlice[m]
            			//tmp[cnt] = 
            			res[i][cnt] = tmpRes
            			fmt.Println(res)
            			cnt++
            			a++
            		}
            		
        		}
		    }
		}
		fmt.Println(tmp)
		
	}
	fmt.Println(a)
	return res
}
func main() {
    //fmt.Println(getFactorial(4))
    fmt.Println(makeCombinations([]string{"1", "2", "3", "4", "5"}))
    fmt.Println("Hello World")
}