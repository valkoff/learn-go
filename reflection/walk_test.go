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
				{2023, "Milan"},
			},
			[]string{"Rome", "Milan"},
		},
		{
			"arrays",
			[2]Profile{
				{1991, "Rome"},
				{2023, "Milan"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{1991, "Rome"}
			aChannel <- Profile{2019, "Milan"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Rome", "Milan"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
