package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var total_rating int =0;
func getRating(student int,wg *sync.WaitGroup) {
	defer wg.Done()
	//x:=rand.Intn(10)
	time.Sleep(1*time.Second)
	rating:= rand.Intn(10)
	total_rating+=rating
	fmt.Println("rating given by student ",student," is ",rating);
}

func main() {
	var wg sync.WaitGroup

	total_student:=5
	for student := 1;  student<= total_student; student++ {
		wg.Add(1)
		go getRating(student,&wg);
	}
	wg.Wait()
	fmt.Println("average rating on scale of 10 is: ", total_rating/total_student)
}
