package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var total_rating int =0
func getRating(student int ,wg *sync.WaitGroup) {
	defer wg.Done()
	//x:=rand.Intn(10)
	time.Sleep(1*time.Second)
	rating:= rand.Intn(10)
	fmt.Println(student," ",rating);
	total_rating+=rating
}

func main() {

	var wg sync.WaitGroup

	for student := 1;  student<= 5; student++ {
		wg.Add(1)
		go getRating(student,&wg)
		//total_rating+=rating;
	}

	wg.Wait()
	fmt.Println(total_rating)
}
