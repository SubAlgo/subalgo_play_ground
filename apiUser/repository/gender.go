package repository

import (
	seterror "apiUser/setError"
	"fmt"
	"log"
)

type RepoGender interface {
	CreateGender(name string) (status string, err error)
	GetGender(id uint) Gender
	GetGenderByName(name string) Gender
	GetGenders() []Gender
	Update(id uint, name string) (status string, err error)
	Delete(id uint) error
}

type Gender struct {
	ID   uint
	Name string
}

var (
	GenderArray = []Gender{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "SE"},
	}
)

func NewGenderStoreArray() RepoGender {
	return &genderStoreArray{}
}

type genderStoreArray struct{}

func (g *genderStoreArray) CreateGender(name string) (status string, err error) {
	newID := GenderArray[len(GenderArray)-1].ID + 1
	GenderArray = append(GenderArray, Gender{ID: newID, Name: name})
	fmt.Println(GenderArray)
	return "Success", nil
}

func (g *genderStoreArray) GetGender(id uint) Gender {
	res := Gender{}
	for _, v := range GenderArray {
		if v.ID == id {
			res = v
		}
	}
	return res
}

func (g *genderStoreArray) GetGenderByName(name string) Gender {
	for _, v := range GenderArray {
		if v.Name == name {
			return v
		}
	}
	return Gender{}
}

func (g *genderStoreArray) GetGenders() []Gender {
	return GenderArray
}

func (g *genderStoreArray) Update(id uint, name string) (status string, err error) {
	for idx, v := range GenderArray {
		if v.ID == id {
			fmt.Println(idx)
			GenderArray[idx].Name = name
			fmt.Println(GenderArray[idx])
			return "Success", nil
		}
	}
	log.Printf("error: repo->genderStoreArray.Update\n")
	return "Fail", seterror.ErrorRepoUpdate
}

func (g *genderStoreArray) Delete(id uint) error {
	for idx, v := range GenderArray {
		if v.ID == id {
			GenderArray = append(GenderArray[:idx], GenderArray[idx+1:]...)
		}
	}
	return nil
}
