package ropository

import "fmt"

type RepoGender interface {
	CreateGender(name string) (status string)
	GetGender(id uint) Gender
	GetGenderByName(name string) Gender
	GetGenders() []Gender
	Update(id uint, name string) (status string)
	Delete(id uint)
}

type Gender struct {
	ID   uint
	Name string
}

var (
	genderArray = []Gender{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "SE"},
	}
)

func NewGenderStoreArray() RepoGender {
	return &genderStoreArray{}
}

type genderStoreArray struct{}

func (g *genderStoreArray) CreateGender(name string) (status string) {
	newID := genderArray[len(genderArray)-1].ID + 1
	genderArray = append(genderArray, Gender{ID: newID, Name: name})
	fmt.Println(genderArray)
	return "Success"
}

func (g *genderStoreArray) GetGender(id uint) Gender {
	res := Gender{}
	for _, v := range genderArray {
		if v.ID == id {
			res = v
		}
	}
	return res
}

func (g *genderStoreArray) GetGenderByName(name string) Gender {
	for _, v := range genderArray {
		if v.Name == name {
			return v
		}
	}
	return Gender{}
}

func (g *genderStoreArray) GetGenders() []Gender {
	return genderArray
}

func (g *genderStoreArray) Update(id uint, name string) (status string) {
	for idx, v := range genderArray {
		if v.ID == id {
			fmt.Println(idx)
			genderArray[idx].Name = name
			fmt.Println(genderArray[idx])
			return "Success"
		}
	}
	return "Fail"
}

func (g *genderStoreArray) Delete(id uint) {
	for idx, v := range genderArray {
		if v.ID == id {
			genderArray = append(genderArray[:idx], genderArray[idx+1:]...)
		}
	}
}
