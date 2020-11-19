package main

import (
	"fmt"
	"net/rpc"
)

func GetStudentName() string {
	var studentName string
	fmt.Print("Student: ")
	fmt.Scanln(&studentName)
	return studentName
}

func GetSubject() string {
	var subject string
	fmt.Print("Subject: ")
	fmt.Scanln(&subject)
	return subject
}

func GetGrade() float64 {
	var grade float64
	fmt.Print("Grade: ")
	fmt.Scanln(&grade)
	return grade
}

func AddStudentGradeBySubject() map[string]map[string]float64 {
	var subjects = make(map[string]map[string]float64)
	var subject = GetSubject()
	var studentName = GetStudentName()
	var grade = GetGrade()
	student := make(map[string]float64)
	student[studentName] = grade
	subjects[subject] = student

	return subjects
}

func client() {
	
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Add Student Grade By Subject")
		fmt.Println("2) Get Student Average")
		fmt.Println("3) Get Students General Average")
		fmt.Println("4) Get Subject Average")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var result string
			err = c.Call("Server.AddStudentGradeBySubject", AddStudentGradeBySubject(), &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AddStudentGradeBySubject =", result)
			}
		case 2:
			var studentName = GetStudentName()
			var result string
			err = c.Call("Server.GetStudentAverage", studentName, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.GetStudentAverage =", result)
			}
		case 3:
			var result string
			err = c.Call("Server.GetStudentsGeneralAverage", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.GetStudentsGeneralAverage =", result)
			}
		case 4:
			var subject = GetSubject()
			var result string
			err = c.Call("Server.GetSubjectAverage", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.GetSubjectAverage =", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
