package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenderStoreArray(t *testing.T) {
	g := NewGenderStoreArray()
	t.Run("1_create_gender", func(t *testing.T) {
		g.CreateGender("Dev")

		want := []Gender{
			{ID: 1, Name: "Admin"},
			{ID: 2, Name: "SE"},
			{ID: 3, Name: "Dev"},
		}
		got := GenderArray
		assert.Equal(t, want, got)
	})

	t.Run("2_GetGender", func(t *testing.T) {
		testValue := []struct {
			input uint
			want  Gender
		}{
			{input: 1, want: Gender{ID: 1, Name: "Admin"}},
			{input: 2, want: Gender{ID: 2, Name: "SE"}},
			{input: 100, want: Gender{}},
		}

		for _, v := range testValue {
			got := g.GetGender(v.input)
			assert.Equal(t, v.want, got)
		}
	})

	t.Run("3_ GetGenderByName", func(t *testing.T) {
		testValue := []struct {
			input string
			want  Gender
		}{
			{input: "Admin", want: Gender{ID: 1, Name: "Admin"}},
			{input: "OP", want: Gender{}},
		}

		for _, v := range testValue {
			got := g.GetGenderByName(v.input)
			assert.Equal(t, v.want, got)
		}
	})

	t.Run("4_GetGenders", func(t *testing.T) {
		want := GenderArray
		got := g.GetGenders()
		assert.Equal(t, want, got)
	})

	t.Run("5_Update", func(t *testing.T) {
		// reset genderArrat
		GenderArray = []Gender{
			{ID: 1, Name: "Admin"},
			{ID: 2, Name: "SE"},
		}
		got, _ := g.Update(1, "Admin1")
		want := "Success"
		assert.Equal(t, want, got)

		want1 := []Gender{
			{ID: 1, Name: "Admin1"},
			{ID: 2, Name: "SE"},
		}
		assert.Equal(t, want1, GenderArray)

		got, _ = g.Update(100, "Admin1")
		want = "Fail"
		assert.Equal(t, want, got)
	})

	t.Run("6_Delete", func(t *testing.T) {
		// reset genderArrat
		GenderArray = []Gender{
			{ID: 1, Name: "Admin"},
			{ID: 2, Name: "SE"},
			{ID: 3, Name: "Dev"},
		}
		_ = GenderArray
		want := []Gender{
			{ID: 1, Name: "Admin"},
			{ID: 3, Name: "Dev"},
		}
		_ = want
		g.Delete(2)
		assert.Equal(t, want, GenderArray)
	})

}
