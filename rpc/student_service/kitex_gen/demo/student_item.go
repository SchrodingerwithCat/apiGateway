package demo

import "strings"

type StudentItem struct {
	Id             int32
	Name           string
	CollegeName    string
	CollegeAddress string
	Email          string
}

func Student2StudentItem(student *Student) *StudentItem {
	return &StudentItem{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func StudentItem2Student(student *StudentItem) *Student {
	return &Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}
