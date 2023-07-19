package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Year int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Valerio"},
			[]string{"Valerio"},
		},
		{
			"struct with two string field",
			struct {
				Name    string
				Surname string
			}{"Valerio", "Cofano"},
			[]string{"Valerio", "Cofano"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Year int
			}{"Valerio", 1991},
			[]string{"Valerio"},
		},
		{
			"nested fields",
			Person{
				"Valerio",
				Profile{
					1991,
					"Rome",
				},
			},
			[]string{"Valerio", "Rome"},
		},
		{
			"pointers to things",
			&Person{
				"Valerio",
				Profile{1991, "Rome"},
			},
			[]string{"Valerio", "Rome"},
		},
		{
			"slices",
			[]Profile{
				{1991, "Rome"},
				{2018, "Milan"},
			},
			[]string{"Rome", "Milan"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
