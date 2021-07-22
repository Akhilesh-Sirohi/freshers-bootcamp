package main

import (
	"fmt"
)


func count(data [5]string, char chan rune, done chan bool){
	ans:=0;
	ch:=<-char
	for i:=0;i<len(data);i++{
		for j:=0;j<len(data[i]);j++{
			if ch == rune(data[i][j]){
				ans++;
			}
		}
	}
	if ans>0 {
		fmt.Printf("%q:  ", ch)
		fmt.Println(ans);
	}
	done<-true
}
func main() {
	//why if a charater is two time the result seems distorted
	data :=[5]string{"quick", "brown", "fox", "lazy", "dog"}
	fmt.Println(data)
	char:=make(chan rune,26);
	done:=make(chan bool,26);
	for i:=0;i<26;i++{
		go count(data,char,done)
	}
	for i:='a';i<='z';i++{
		char<-i
	}
	for i:=0;i<26;i++{
		<-done
	}
}

