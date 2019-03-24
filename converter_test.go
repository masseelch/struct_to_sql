package struct_to_sql

import (
	"testing"
)

type InsertI struct {
	ID          uint   `db:"id" sql:"iu" sqli:"iu"`
	Title       string `db:"title" sql:"i" sqli:"i"`
	Description string `sql:"u" sqli:"u"`
}

func TestConverter_cols(t *testing.T) {
	c := New()

	s := InsertI{}

	// Should return error "no flag provided".
	if _, err := c.cols(&s); err == nil || err.Error() != "no flag provided" {
		t.Errorf("Expected an error of 'no flag provided', but got %#v", err)
	}

	// Test 'i' with defaults.
	ev := []string{"id", "title"}
	av, _ := c.cols(&s, "i")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected cols to be '%#v', but got '%#v'", ev, av)
			break
		}
	}

	// Test 'i' with custom flags.
	c.FlagsName = "sqli"
	c.TagName = "json"
	av, _ = c.cols(&s, "i")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected cols to be '%#v', but got '%#v'", ev, av)
			break
		}
	}

	// Test flag 'u'.
	c.FlagsName = KDefaultFlagsName
	c.TagName = KDefaultTagName
	ev = []string{"id", "description"}
	av, _ = c.cols(&s, "u")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected cols to be '%#v', but got '%#v'", ev, av)
			break
		}
	}
}

func TestConverter_vals(t *testing.T) {
	c := New()

	s := InsertI{
		ID: 1,
		Title: "title",
		Description: "description",
	}

	// Should return error "no flag provided".
	if _, err := c.cols(&s); err == nil || err.Error() != "no flag provided" {
		t.Errorf("Expected an error of 'no flag provided', but got %#v", err)
	}

	// Test 'i' with defaults.
	ev := Vals{uint(1), "title"}
	av, _ := c.vals(&s, "i")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected vals to be '%#v', but got '%#v'", ev, av)
			break
		}
	}

	// Test 'i' with custom flags.
	c.FlagsName = "sqli"
	c.TagName = "json"
	av, _ = c.vals(&s, "i")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected vals to be '%#v', but got '%#v'", ev, av)
			break
		}
	}

	// Test flag 'u'.
	c.FlagsName = KDefaultFlagsName
	c.TagName = KDefaultTagName
	ev = Vals{uint(1), "description"}
	av, _ = c.vals(&s, "u")
	for i := range ev {
		if ev[i] != av[i] {
			t.Errorf("Expected vals to be '%#v', but got '%#v'", ev, av)
			break
		}
	}
}
