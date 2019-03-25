package struct_to_sql

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

const (
	KDefaultFlagsName = "sql"
	KDefaultTagName   = "db"
	kInsertFlag       = "i"
	kUpdateFlag       = "u"
	kQueryFlag        = "q"
)

type Cols []string
type Vals []interface{}

type Converter struct {
	FlagsName string
	TagName   string
}

func New() *Converter {
	return &Converter{
		FlagsName: KDefaultFlagsName,
		TagName:   KDefaultTagName,
	}
}

func (c *Converter) InsertCols(s interface{}) (Cols, error) {
	return c.cols(s, kInsertFlag)
}

func (c *Converter) UpdateCols(s interface{}) (Cols, error) {
	return c.cols(s, kUpdateFlag)
}

func (c *Converter) QueryCols(s interface{}) (Cols, error) {
	return c.cols(s, kQueryFlag)
}

func (c *Converter) InsertVals(s interface{}) (Vals, error) {
	return c.vals(s, kInsertFlag)
}

func (c *Converter) UpdateVals(s interface{}) (Vals, error) {
	return c.vals(s, kUpdateFlag)
}

func (c *Converter) SelectVals(s interface{}) (Vals, error) {
	return c.vals(s, kQueryFlag)
}

func (c *Converter) cols(s interface{}, flags ...string) (Cols, error) {
	if len(flags) == 0 {
		return nil, errors.New("no flag provided")
	}

	// Reflect type.
	t := reflect.TypeOf(s)

	// If the given interface is not of kind ptr raise an error.
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("passed value is not a pointer")
	}

	t = t.Elem()

	cols := make(Cols, 0)

	// Iterate over all available fields and read the tag value.
	for i := 0; i < t.NumField(); i++ {
		// Get the current field.
		field := t.Field(i)

		// Get the field tag value.
		fieldFlags := field.Tag.Get(c.FlagsName)

		// If one of the given flags is contained in the struct flags
		// add the current field to the Cols.
		for _, f := range flags {
			if strings.Contains(fieldFlags, f) {
				col := field.Tag.Get(c.TagName)
				if col == "" {
					col = strings.ToLower(field.Name)
				}
				cols = append(cols, col)
			}
		}
	}

	return cols, nil
}

func (c *Converter) vals(s interface{}, flags ...string) (Vals, error) {
	if len(flags) == 0 {
		return nil, errors.New("no flag provided")
	}

	// Reflect type.
	t := reflect.TypeOf(s)

	// If the given interface is not of kind ptr raise an error.
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("passed value is not a pointer")
	}

	t = t.Elem()
	v := reflect.ValueOf(s).Elem()

	vals := make(Vals, 0)

	// Iterate over all available fields and read the tag value.
	for i := 0; i < t.NumField(); i++ {
		// Get the current field.
		field := t.Field(i)

		// Get the field tag value.
		fieldFlags := field.Tag.Get(c.FlagsName)

		// If one of the given flags is contained in the struct flags
		// add the current field value to the Vals.
		for _, f := range flags {
			if strings.Contains(fieldFlags, f) {
				vals = append(vals, v.Field(i).Interface())
			}
		}
	}

	return vals, nil
}
