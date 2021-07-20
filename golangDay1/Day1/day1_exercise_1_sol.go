package main

import (
	"errors"
	"fmt"
)

type matrix struct {
	number_of_rows int
	number_of_cols int
	elements [][]int
	//elements:=make([][]int, number_of_rows,number_of_cols)
}

func getRows(name matrix) int{
	return name.number_of_rows;
}
func getColumns(name matrix) int{
	return name.number_of_cols;
}
func setElement(name matrix,i int, j int, value int) {
	name.elements[i][j]=value;
}
func addMarix(name1 matrix, name2 matrix) matrix{
	if getRows(name1)!=getRows(name2) || getColumns(name1)!=getColumns(name2) {
		//errors.New constructs a basic error value with the given error message.
		errors.New("Can't add both matrix as their dimensions are different")
	} else {
		ans := name1
		for i := 0; i < name1.number_of_rows; i++ {
			for j := 0; j < name1.number_of_cols; j++ {
				ans.elements[i][j] = name1.elements[i][j] + name2.elements[i][j]
			}
		}
		return ans;
	}
	ans:= matrix{0,0,make ([][]int,1)}
	return ans
}
func main() {
	x := matrix{15, 16, make([][]int, 3)}
	//setElement(x,1,1,12)
	y := matrix{16, 17, make([][]int, 3)}
	fmt.Println(getRows(x))
	fmt.Println(getColumns(x))
	fmt.Println(x.elements)
	fmt.Println(addMarix(x,y))
}