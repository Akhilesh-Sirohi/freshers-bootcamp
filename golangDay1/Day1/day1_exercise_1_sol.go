package main

import (
	"fmt"
	"math/rand"
)

type matrix struct {
	number_of_rows int
	number_of_cols int
	elements [][]int
	//elements:=make([][]int, number_of_rows,number_of_cols)
}

func createMatrix(number_of_rows int, number_of_cols int) matrix{
	elements:= make([][]int, number_of_rows);
	for i:=0;i<number_of_rows;i++{
		elements[i]= make([] int,number_of_cols);
		for j:=0;j<number_of_cols;j++{
			elements[i][j]=rand.Intn(10);
		}
	}
	return matrix{number_of_rows,number_of_cols,elements}
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
func addMatrix(name1 matrix, name2 matrix) matrix{
	ans:= matrix{0,0,make ([][]int,1)}
	if getRows(name1)!=getRows(name2) || getColumns(name1)!=getColumns(name2) {
		//errors.New constructs a basic error value with the given error message.
		//fmt.Println("Failed: Can't add both matrix as their dimensions are different"
		//x=errors.New("Failed: Can't add both matrix as their dimensions are different")
	} else {
		ans = name1
		for i := 0; i < name1.number_of_rows; i++ {
			for j := 0; j < name1.number_of_cols; j++ {
				ans.elements[i][j] = name1.elements[i][j] + name2.elements[i][j]
			}
		}
	}
	return ans;
}
func print(name matrix) {
	fmt.Printf("{")
	for i := 0; i < name.number_of_rows; i++ {
		fmt.Printf("%s", "{")
		for j := 0; j < name.number_of_cols; j++ {
			fmt.Printf("%d,", name.elements[i][j])
		}
		fmt.Printf("},")
	}
	fmt.Printf("}")
}
func main() {
	x := createMatrix(4,5);
	//setElement(x,1,1,12)
	y := createMatrix(4,5);
	setElement(x,0,0,12);
	setElement(y,0,0,13);
	fmt.Println(getRows(x))
	fmt.Println(getColumns(x))
	fmt.Println(x.elements)
	print(addMatrix(x,y))
	fmt.Println(addMatrix(createMatrix(4,5),createMatrix(6,3)))
}