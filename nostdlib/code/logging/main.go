package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("App started")
	id := 6
	num := []int{1,4,8,45}
	desc := "ComplexCalc"
	step := "step 1"
	err := fmt.Errorf("error calculating result")
	log.Printf("ID: %d (%s) errored %v on %v", id, desc, err, num)
	log.Printf("ID: %d (%s) completed step %q", id, desc, step, num)

	log.Printf("ID: %d persisted", id)

	log.Printf("ID: (%s) %d finalised results", desc, id)
	log.Printf("ID: %d (%s) completed step %q", id, desc, step, num)


	// 2019/01/23 11:28:16 ID: 6 (ComplexCalc) errored error calculating result on [1 4 8 45]

}


func DoComplexWork(id int, num []int, desc string) (int, error){
	res, err := calc(num)
	if err != nil {
		log.Printf("ID: %d (%s) errored %v on %v", id, desc, err, num)
	}
	log.Println()
	return res, nil
}

func calc(num []int) (int, error) {
	if len(num) == 0{
		return 0, fmt.Errorf("no numbers supplied")
	}
	res := 0
	for _, i := range num {
		res += i
	}
	return res, nil
}

