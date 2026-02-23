package services

import (
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

var (
	ErrStudentNotFound = errors.New("student not found")
	ErrInvalidInput    = errors.New("invalid input")
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	if err := s.validateForCreate(student); err != nil {
		return err
	}
	return s.Repo.Create(student)
}
func (s *StudentService) UpdateStudent(id string, student models.Student) (*models.Student, error) {

	if err := s.validateForUpdate(student); err != nil {
		return nil, err
	}

	updated, err := s.Repo.Update(id, student)
	if err != nil {
		return nil, ErrStudentNotFound
	}

	return updated, nil
}

func (s *StudentService) DeleteStudent(id string) error {
	err := s.Repo.Delete(id)
	if err != nil {
		return ErrStudentNotFound
	}

	return nil
}

func (s *StudentService) validateForCreate(st models.Student) error {
	if st.Id == "" {
		return ErrInvalidInput
	}
	if st.Name == "" {
		return ErrInvalidInput
	}
	if st.GPA < 0.0 || st.GPA > 4.0 {
		return ErrInvalidInput
	}
	return nil
}

func (s *StudentService) validateForUpdate(st models.Student) error {
	if st.Name == "" {
		return ErrInvalidInput
	}
	if st.GPA < 0.0 || st.GPA > 4.0 {
		return ErrInvalidInput
	}
	return nil
}
