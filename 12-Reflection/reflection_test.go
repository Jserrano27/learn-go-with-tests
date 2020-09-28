package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "California"},
			[]string{"Chris", "California"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 28},
			[]string{"Chris"},
		},
		{
			"Struct with nested field",
			Person{
				"Chris",
				Profile{33, "California"},
			},
			[]string{"Chris", "California"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "California"},
			},
			[]string{"Chris", "California"},
		},
		{
			"Slices",
			[]Profile{
				{28, "California"},
				{25, "New York"},
			},
			[]string{"California", "New York"},
		},
		{
			"Arrays",
			[2]Profile{
				{28, "California"},
				{25, "New York"},
			},
			[]string{"California", "New York"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("Expected %v, but got %v instead", test.Expected, got)
			}
		})
	}

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("With channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{35, "California"}
			aChannel <- Profile{30, "New York"}
			close(aChannel)
		}()

		expected := []string{"California", "New York"}
		got := []string{}
		Walk(aChannel, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})

	t.Run("With function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{35, "California"}, Profile{30, "New York"}
		}

		got := []string{}
		expected := []string{"California", "New York"}

		Walk(aFunction, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, val := range haystack {
		if val == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("Expected %q not found", needle)
	}
}
