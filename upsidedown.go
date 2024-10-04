package main

import "fmt"

func main() {
	a := [4]string{
		"1",
		"2",
		"3",
		"4",
	}
	//finalList := []string{}

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)
}

//tasks := startList[0:3]
//length := len(startList)
//for i := range length {
//	fmt.Println(tasks[i])
//}
