package name

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"github.com/ignite-laboratories/core/enum/gender"
	"github.com/ignite-laboratories/core/sys/log"
	"math/rand"
	"regexp"
	"strings"
)

var moduleName = "name"

func init() {
	reader := csv.NewReader(strings.NewReader(nameDB))
	reader.Comma = '\t'

	i := 0
	for {
		record, err := reader.Read() // Read a single line
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file
			}

			log.Verbosef(moduleName, "error reading name database: %v\n", err)
			panic(err)
		}

		genderFunc := func(s string) gender.Gender {
			if s == "Male" {
				return gender.Male
			} else if s == "Female" {
				return gender.Female
			} else {
				return gender.NonBinary
			}
		}

		entry := Given{
			Name:        strings.TrimSpace(record[0]),
			Description: strings.TrimSpace(record[3]),
			Details: struct {
				Origin string
				Gender gender.Gender
			}{
				Origin: strings.TrimSpace(record[1]),
				Gender: genderFunc(strings.TrimSpace(record[2])),
			},
		}
		Database = append(Database, entry)

		i++
	}

	log.Verbosef(moduleName, "name database loaded\n")
}

//go:embed names.csv
var nameDB string

// Database provides a collection of cultural names for seeding identifiers with.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database!
// https://incompetech.com
//
// Please check his stuff out, he's quite clever!
var Database = make([]Given, 0, 8888)

// New creates a new Given name.  You may optionally provide a description during creation.
func New(name string, description ...string) Given {
	if len(description) > 0 {
		return Given{
			Name: name,
		}
	}
	return Given{
		Name:        name,
		Description: description[0],
	}
}

// Random returns a random name from the Database.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
func Random(database ...[]Given) Given {
	if len(database) > 0 {
		names := database[0]
		return names[rand.Intn(len(names))]
	}
	return Database[rand.Intn(len(Database))]
}

// Filtered returns a random name from the Database which passes the provided predicate check.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
func Filtered(predicate func(Given) bool, database ...[]Given) Given {
	for {
		name := Random(database...)
		if predicate(name) {
			return name
		}
	}
}

// Lookup finds the provided name in the Database, otherwise it returns nil.
func Lookup(name string, caseInsensitive ...bool) (Given, error) {
	for _, n := range Database {
		if len(caseInsensitive) > 0 && caseInsensitive[0] {
			if strings.EqualFold(n.Name, name) {
				return n, nil
			}
		} else {
			if n.Name == name {
				return n, nil
			}
		}
	}
	return Given{}, fmt.Errorf("name not found")
}

/**
tiny
*/

// Tiny is a standard function for returning a name from the Database which satisfies tiny's requirements for implicit naming.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
//
// See tiny.NameFilter for the explicit details in use.
func Tiny(database ...[]Given) Given {
	return Filtered(tinyNameFilter)
}

var usedTinyNames = make(map[string]*Given)

// tinyNameFilter is a standard function for returning a name which satisfies tiny's requirements for implicit naming.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word alpha-explicit names)
//   - At least three characters in length
//   - At least 2¹⁴ unique names before beginning to recycling names
//   - Names are case-sensitive in uniqueness.
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
func tinyNameFilter(name Given) bool {
	var nonAlphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

	if len(usedTinyNames) >= 1<<14 {
		usedTinyNames = make(map[string]*Given)
	}

	if nonAlphaRegex.MatchString(name.Name) && usedTinyNames[name.Name] == nil && len(name.Name) > 2 {
		usedTinyNames[name.Name] = &name
		return true
	}
	return false
}
