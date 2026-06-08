package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model           // Include following fields to the DB: ID, CreatedAt, UpdatedAt e DeletedAt
	Name         string  `json:"name" validate:"required,min=2,max=120"`
	ClassId      uint    `gorm:"default:0" json:"class_id"`
	TotalFaults  uint    `gorm:"default:0" json:"total_faults" validate:"gte=0"`
	GradeAverage float64 `json:"grade_average" validate:"required,gte=0,lte=100"`
}

func (Student) TableName() string {
	return "students"
}

func StudentExistId(Id int64) bool {
	var student Student
	db.Where("Id=?", Id).Find(&student)
	return !db.RecordNotFound()
}

func StudentExistName(name string) bool {
	var student Student
	db.Where("name = ?", name).First(&student)
	return !db.RecordNotFound()
}

func (s *Student) CreateStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentbyId(ID int64) (*Student, *gorm.DB) {
	var s Student
	db := db.Where("Id=?", ID).Find(&s)

	return &s, db
}

func DeleteStudent(ID int64) (Student, error) {
	var s Student
	// Find the student with the given ID
	if err := db.First(&s, ID).Error; err != nil {
		// Return an error if the student is not found
		return s, err
	}
	
	if s.ClassId != 0 {
		RemoveStudentFromClass(uint(ID), s.ClassId)
	}

	// Delete the student
	if err := db.Delete(&s).Error; err != nil {
		// Return an error if there was an issue deleting the student
		return s, err
	}
	// Return the deleted student
	return s, nil
}
