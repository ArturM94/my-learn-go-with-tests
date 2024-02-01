package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
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
			}{"Artur"},
			[]string{"Artur"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Artur", "Dnipro"},
			[]string{"Artur", "Dnipro"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Artur", 29},
			[]string{"Artur"},
		},
		{
			"nested fields",
			Person{
				"Artur",
				Profile{29, "Dnipro"},
			},
			[]string{"Artur", "Dnipro"},
		},
		{
			"pointers to things",
			&Person{
				"Artur",
				Profile{29, "Dnipro"},
			},
			[]string{"Artur", "Dnipro"},
		},
		{
			"slices",
			[]Profile{
				{29, "Dnipro"},
				{34, "London"},
			},
			[]string{"Dnipro", "London"},
		},
		{
			"arrays",
			[2]Profile{
				{29, "Dnipro"},
				{34, "London"},
			},
			[]string{"Dnipro", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
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
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
