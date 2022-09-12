package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var names []string = readMemberList()
	var grpSiz int = 3
	randomizeAndAssign(names, grpSiz)
}

func readMemberList() []string {
	//open student_names file
	file, err := os.Open("student_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Create Slice
	var students []string
	//Add student names from text file to slice
	for scanner.Scan() {
		students = append(students, scanner.Text())
	}

	return students
}

func randomizeAndAssign(names []string, grpSiz int) {
	//----------Inputs---------
	//names := []string{"Sam", "Bill", "Devon", "Erin", "Fali", "Gabby", "Hunter", "Isaac", "Julia", "Kaylee"}

	//----------Vars-----------
	var numStu int = len(names)
	var numGrp int = (numStu / grpSiz)
	var grpRm int = numStu - (grpSiz * numGrp)

	//----------Spots----------
	grpRang := makeRange(1, numGrp)
	grpsL := []int{}

	for i := 0; i < grpSiz; i++ {
		grpsL = append(grpsL, grpRang...)
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
	//Make an array of arrays. Each array in the array is a team.

	test := [][]string{}

	for i := 0; i < numGrp; i++ {
		groupArray := []string{}
		test = append(test, groupArray)
	}

	//This gives a random teams spot out. grpsL is the available spots for the team
	//G is the number of the team
	for n := 0; n < numStu; n++ {
		rand.Seed(time.Now().UnixNano())
		RN := rand.Intn(len(grpsL))
		G := grpsL[RN]
		grpsL[RN] = grpsL[len(grpsL)-1]
		grpsL[len(grpsL)-1] = 0
		grpsL = grpsL[:len(grpsL)-1]
		test[G-1] = append(test[G-1], names[n])
	}
	fmt.Println("test:")
	fmt.Println(test)

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func printTeams() {}
