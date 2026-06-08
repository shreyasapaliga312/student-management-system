package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const MAX_STUDENT_NUMBER = 20

type Class struct {
	gorm.Model
	TotalStudentNumber uint      `gorm:"default:0"`
	Students           []Student `gorm:"many2many:class_students;"`
}

func (Class) TableName() string {
	return "classes"
}

func (class *Class) CreateClass() *Class {
	db.NewRecord(class)
	db.Create(&class)
	return class
}

func (class *Class) IsClassFull() bool {
	return class.TotalStudentNumber == MAX_STUDENT_NUMBER
}

func ClassExist(Id int64) bool {
	var class Class
	db.Where("Id=?", Id).Find(&class)
	return !db.RecordNotFound()
}

func GetClassById(ID int64) (*Class, *gorm.DB, error) {
	var class Class
	db := db.Where("Id=?", ID).Find(&class)

	if db.RecordNotFound() {
		return nil, nil, fmt.Errorf("Class not found")
	}
	if err := db.Error; err != nil {
		return nil, nil, err
	}

	db.Preload("Students").First(&class, class.ID)
	return &class, db, nil
}

func DeleteClass(ID int64) (Class, error) {
	var class Class

	if err := db.First(&class, ID).Error; err != nil {
		return class, err
	}
	
	if class.TotalStudentNumber > 0 { // Only empty classes will be deleted
		return class, fmt.Errorf("the required class still have %v students. Remove them before deleting the class", class.TotalStudentNumber)
	}

	if err := db.Delete(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}

func GetAllClasses() []Class {
	var Classes []Class
	db.Preload("Students").Find(&Classes)
	return Classes
}

func EnrollStudentInClass(studentId uint, classId uint) error { // Ver se aluno e turma existem
	if !ClassExist(int64(classId)) || !StudentExistId(int64(studentId)) {
		return fmt.Errorf("Class or Student does not exist")
	}

	student := Student{}
	if err := db.First(&student, studentId).Error; err != nil {
		return err
	}

	oldClassId := student.ClassId

	student.ClassId = classId
	if err := db.Save(&student).Error; err != nil { // Update and save the student
		return err
	}

	// If the student was already enrolled in a class, remove them from that class
	if oldClassId != 0 {
		if err := RemoveStudentFromClass(student.ID, oldClassId); err != nil {
			return err
		}
	}
	// Add the student to the new class
	newClass := Class{}
	if err := db.Preload("Students").First(&newClass, classId).Error; err != nil {
		return err
	}

	if newClass.IsClassFull() {
		return fmt.Errorf("Class is full")
	}

	newClass.TotalStudentNumber += 1

	newClass.Students = append(newClass.Students, student)

	if err := db.Save(&newClass).Error; err != nil {
		return err
	}

	return nil
}

func RemoveStudentFromClass(studentId uint, oldClassId uint) error {
	if !ClassExist(int64(oldClassId)) || !StudentExistId(int64(studentId)) {
		return fmt.Errorf("Class or Student does not exist")
	}

	student := Student{}
	if err := db.First(&student, studentId).Error; err != nil {
		return err
	}

	if student.ClassId != oldClassId {
		return fmt.Errorf("Student is not in the class")
	}
	student.ClassId = 0
	if err := db.Save(&student).Error; err != nil {
		return err
	}

	oldClass := Class{}
	if err := db.Preload("Students").First(&oldClass, oldClassId).Error; err != nil {
		return err
	}
	// Remove the student from the Students slice
	var updatedStudents []Student
	for _, s := range oldClass.Students {
		if s.ID != studentId {
			updatedStudents = append(updatedStudents, s)
		}
	}

	oldClass.Students = updatedStudents

	oldClass.TotalStudentNumber -= 1

	if err := db.Model(&oldClass).Association("Students").Replace(oldClass.Students).Error; err != nil {
		return err
	}

	if err := db.Save(&oldClass).Error; err != nil {
		return err
	}

	return nil
}
