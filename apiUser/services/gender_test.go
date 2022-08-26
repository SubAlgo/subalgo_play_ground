package services_test

import (
	"apiUser/repository"
	"apiUser/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	s := services.NewServiceGender()

	t.Run("1_create", func(t *testing.T) {
		res, err := s.Create("Dev")

		assert.Equal(t, "Success", res)
		assert.Equal(t, nil, err)
	})

	t.Run("2_GetRow", func(t *testing.T) {
		got := s.GetRow(1)
		want := repository.Gender{
			ID: 1, Name: "Admin",
		}
		assert.Equal(t, want, got)
	})

	t.Run("3_GetRowByName", func(t *testing.T) {
		got := s.GetRowByName("Admin")
		want := repository.Gender{
			ID: 1, Name: "Admin",
		}
		assert.Equal(t, want, got)

		got = s.GetRowByName("Test")
		want = repository.Gender{}
		assert.Equal(t, want, got)
	})

	t.Run("4_GetRows", func(t *testing.T) {
		got := s.GetRows()
		want := repository.GenderArray
		assert.Equal(t, want, got)
	})

	t.Run("5_Update", func(t *testing.T) {
		id := uint(1)
		got, err := s.Update(id, "Admin01")
		want := "Success"

		assert.Equal(t, want, got)
		assert.Equal(t, nil, err)
	})

	t.Run("6_Delete", func(t *testing.T) {
		id := uint(1)
		got := s.Delete(id)
		assert.Equal(t, nil, got)
	})
}
