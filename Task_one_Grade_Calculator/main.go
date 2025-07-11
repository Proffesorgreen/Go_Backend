package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Subjects struct {
	name           string
	grade          int
	alphabet_grade string
}

var subjects []Subjects
var m = map[string]float64{
	"A+": 4.0,
	"A":  4.0,
	"A-": 3.75,
	"B+": 3.5,
	"B":  3.0,
	"B-": 2.75,
	"C+": 2.5,
	"C":  2.0,
	"D":  1.75,
	"F":  0,
}

func average_grade(subjects []Subjects) float64 {
	total_sum := 0.0
	for i := range subjects {
		total_sum += m[subjects[i].alphabet_grade]
	}
	return total_sum / float64(len(subjects))
}

func alphabet_grade(subjects []Subjects) {
	for i := range subjects {
		if subjects[i].grade >= 90 {
			subjects[i].alphabet_grade = "A+"
		} else if subjects[i].grade >= 85 && subjects[i].grade < 90 {
			subjects[i].alphabet_grade = "A"
		} else if subjects[i].grade >= 80 && subjects[i].grade < 85 {
			subjects[i].alphabet_grade = "A-"
		} else if subjects[i].grade >= 75 && subjects[i].grade < 80 {
			subjects[i].alphabet_grade = "B+"
		} else if subjects[i].grade >= 70 && subjects[i].grade < 75 {
			subjects[i].alphabet_grade = "B"
		} else if subjects[i].grade >= 65 && subjects[i].grade < 70 {
			subjects[i].alphabet_grade = "B-"
		} else if subjects[i].grade >= 60 && subjects[i].grade < 65 {
			subjects[i].alphabet_grade = "C+"
		} else if subjects[i].grade >= 55 && subjects[i].grade < 60 {
			subjects[i].alphabet_grade = "C"
		} else if subjects[i].grade >= 50 && subjects[i].grade < 55 {
			subjects[i].alphabet_grade = "D"
		} else {
			subjects[i].alphabet_grade = "F"
		}
	}
}

func printStudent(s Subjects) {
	output := fmt.Sprintf(`
======== Student Report ========
Name   : %s
Grade    : %d
Alphabet_grade : %.2s
===============================
`, s.name, s.grade, s.alphabet_grade)

	fmt.Println(output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num_subjects int

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	for {
		fmt.Print("Enter the number of subjects you taken: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)

		if err != nil || n <= 0 {
			fmt.Print("Invalid input")
			continue
		}
		num_subjects = n
		break
	}

	for i := 0; i < num_subjects; i++ {
		var grade_subject int

		fmt.Printf("Enter the name of the %d subject: ", i+1)
		name_subject, _ := reader.ReadString('\n')
		name_subject = strings.TrimSpace(name_subject)

		for {
			fmt.Printf("Enter the grade of the %d subject: ", i+1)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			n, err := strconv.Atoi(input)

			if err != nil || 0 > n || n > 100 {
				fmt.Println("Invalid Input")
				continue
			}
			grade_subject = n
			break
		}
		subjects = append(subjects, Subjects{name: name_subject, grade: grade_subject})
	}
	alphabet_grade(subjects)

	fmt.Println("Student_name: ", name)
	fmt.Printf("Average_Grade %f", average_grade(subjects))
	for i := range subjects {
		printStudent(subjects[i])
	}
	fmt.Printf("%-10s %-5s %-5s\n", "Name", "Grade", "Alphabet_Grade")
	for i := range subjects {
		fmt.Printf("%-10s %-5d %-5.2s\n", subjects[i].name, subjects[i].grade, subjects[i].alphabet_grade)
	}
}
