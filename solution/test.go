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
	length := getFactorial(len(cardsSlice)) / getFactorial(5)
	res := make([][]string, length)
	for i:=0;i<length-1;i++ {
		tmp := make([]string, length)
		for j:=i+1;j<length;j++ {
			for k:=j+1;k<length;k++ {
			    for l:=k+1;l<length;l++ {
        			cnt := 0
        			for m:=l+1;m<length;m++ {
        			    tmpRes := cardsSlice[i] + cardsSlice[j] +cardsSlice[k] + cardsSlice[l] +cardsSlice[m]
            			tmp[cnt] = tmpRes
            			fmt.Println(cnt)
            			cnt++
            		}
        		}
		    }
		}
		fmt.Println(tmp)
		res[i] = tmp
	}
	return res
}
func main() {
    fmt.Println(getFactorial(4))
    //fmt.Println(makeCombinations([]string{"1", "2", "3", "4", "5"}))
    fmt.Println("Hello World")
}