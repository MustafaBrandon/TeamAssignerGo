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
	fmt.Println("Enter the group size: ")
	var grpSiz int
	fmt.Scanln(&grpSiz)

	var teams [][]string = randomizeAndAssign(names, grpSiz)

	//Output the teams to a teams.txt file
	printTeams(teams)

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

func randomizeAndAssign(names []string, grpSiz int) [][]string {
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

	teams := [][]string{}

	for i := 0; i < numGrp; i++ {
		groupArray := []string{}
		teams = append(teams, groupArray)
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
		teams[G-1] = append(teams[G-1], names[n])
	}

	return teams

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func printTeams(teams [][]string) {
	f, err := os.Create("teams.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 0; i < len(teams); i++ {
		teamString := fmt.Sprintf("Team %d\n", i+1)
		f.WriteString(teamString)
		for j := 0; j < len(teams[i]); j++ {
			_, err2 := f.WriteString(teams[i][j] + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}
		}
		f.WriteString("\n")
	}
}
