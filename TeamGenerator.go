package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//----------Inputs---------
	names := []string{"a", "b", "d", "e", "f", "g", "h", "i", "j", "k"}
	//var grpSiz int = 3
	var numGrp int = 3

	//fmt.Println(names)
	//fmt.Println(grpSiz)
	//fmt.Println("-------")
	//----------Vars-----------
	var numStu int = len(names)
	//var numGrp int = (numStu / grpSiz)
	var grpSiz int = numStu/numGrp
	//var grpRm int = numStu - (grpSiz * numGrp)
	var grpRm int = numStu % numGrp

	//fmt.Println(numStu)
	//fmt.Println(numGrp)
	//fmt.Println(grpSiz)
	//fmt.Println(grpRm)
	//fmt.Println("-------")
	//----------Spots----------
	grpRang := makeRange(1, numGrp)
	grpsL := []int{}

	for i := 0; i < grpSiz; i++ {
		grpsL = append(grpsL, grpRang...)
		//fmt.Println(grpsL)
	}

	for grpRm > 0 {
		for r := 1; r < numGrp+1; r++ {
			if grpRm > 0 {
				grpsL = append(grpsL, r)
				grpRm = grpRm - 1
			}
		}

	}

	//----------Assigner-------

	for n := 0; n < numStu; n++ {
		rand.Seed(time.Now().UnixNano())
		RN := rand.Intn(len(grpsL))
		G := grpsL[RN]
		grpsL[RN] = grpsL[len(grpsL)-1]
		grpsL[len(grpsL)-1] = 0
		grpsL = grpsL[:len(grpsL)-1]
		fmt.Println(names[n] + " : " + strconv.Itoa(G))

	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}