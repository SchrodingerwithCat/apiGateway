package demo

import "strings"

type TeacherItem struct {
	Id             int32
	Name           string
	CollegeName    string
	CollegeAddress string
	Email          string
}

func Teacher2TeacherItem(teacher *Teacher) *TeacherItem {
	return &TeacherItem{
		Id:             teacher.Id,
		Name:           teacher.Name,
		Email:          strings.Join(teacher.Email, ","),
		CollegeName:    teacher.College.Name,
		CollegeAddress: teacher.College.Address,
	}
}

func TeacherItem2Teacher(teacher *TeacherItem) *Teacher {
	return &Teacher{
		Id:      teacher.Id,
		Name:    teacher.Name,
		Email:   strings.Split(teacher.Email, ","),
		College: &College{Name: teacher.CollegeName, Address: teacher.CollegeAddress},
	}
}
