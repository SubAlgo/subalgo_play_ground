package services

import "apiUser/repository"

type ServiceGender interface {
	Create(name string) (string, error)
	GetRow(id uint) repository.Gender
	GetRowByName(name string) repository.Gender
	GetRows() []repository.Gender
	Update(id uint, name string) (string, error)
	Delete(id uint) error
}

func NewServiceGender() ServiceGender {
	return &serviceGender{
		repoGender: repository.NewGenderStoreArray(),
	}
}

type serviceGender struct {
	repoGender repository.RepoGender
}

func (s *serviceGender) Create(name string) (string, error) {
	return s.repoGender.CreateGender(name)
}

func (s *serviceGender) GetRow(id uint) repository.Gender {
	return s.repoGender.GetGender(id)
}

func (s *serviceGender) GetRowByName(name string) repository.Gender {
	return s.repoGender.GetGenderByName(name)
}

func (s *serviceGender) GetRows() []repository.Gender {
	return s.repoGender.GetGenders()
}

func (s *serviceGender) Update(id uint, name string) (string, error) {
	return s.repoGender.Update(id, name)
}

func (s *serviceGender) Delete(id uint) error {
	return s.repoGender.Delete(id)
}
