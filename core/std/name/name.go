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

//go:embed nameDB.tsv
var nameDBRaw string
var nameDB = make([]Given, 0, 8888)

//go:embed surnameDB.txt
var surnameDBRaw string
var surnameDB = make([]string, 0, 8888)

func init() {
	initNameDB()
	initSurnameDB()

	log.Verbosef(moduleName, "name database loaded\n")
}

func initSurnameDB() {
	reader := csv.NewReader(strings.NewReader(surnameDBRaw))
	reader.Comma = '\t'

	i := 0
	for {
		record, err := reader.Read() // Read a single line
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file
			}

			log.Verbosef(moduleName, "error reading surname database: %v\n", err)
			panic(err)
		}

		surnameDB = append(surnameDB, strings.TrimSpace(record[0]))

		i++
	}
}

func initNameDB() {
	reader := csv.NewReader(strings.NewReader(nameDBRaw))
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
		nameDB = append(nameDB, entry)

		i++
	}
}

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

// Random generates a random name using the provided type format.
//
// If you'd prefer a random name from your own name database, provide it as a parameter
//
// See Format.
func Random[T Format]() Given {
	switch any(T("")).(type) {
	case NameDB:
		return nameDB[rand.Intn(len(nameDB))]
	case SurnameDB:
		return Given{
			Name: surnameDB[rand.Intn(len(surnameDB))],
		}
	case Tiny:
		for {
			name := Random[NameDB]()
			if tinyNameFilter(name) {
				return name
			}
		}
	case Multi, Default: // NOTE: Default can be moved between case statements
		name := nameDB[rand.Intn(len(nameDB))]
		last := surnameDB[rand.Intn(len(surnameDB))]
		name.Name += " " + last
		return name
	default:
		// Just return a random name from the NameDB
		return Random[NameDB]()
	}
}

// Lookup finds the provided name in the provided database, otherwise it returns nil and an error.
//
// NOTE: This will only look up names from the NameDB and SurnameDB databases.
//
// See Format.
func Lookup[T Format](name string, caseInsensitive ...bool) (Given, error) {
	switch any(T("")).(type) {
	case NameDB:
		for _, n := range nameDB {
			if len(caseInsensitive) > 0 && caseInsensitive[0] {
				if strings.EqualFold(string(n.Name), name) {
					return n, nil
				}
			} else {
				if string(n.Name) == name {
					return n, nil
				}
			}
		}
	case SurnameDB:
		for _, n := range surnameDB {
			if len(caseInsensitive) > 0 && caseInsensitive[0] {
				if strings.EqualFold(n, name) {
					return Given{Name: n}, nil
				}
			} else {
				if n == name {
					return Given{Name: n}, nil
				}
			}
		}
	}
	return Given{}, fmt.Errorf("name not found")
}

/**
tiny
*/

var usedTinyNames = make(map[string]*Given)

// !!!CRITICAL NOTE: Please update Tiny if you make a change to this!
func tinyNameFilter(name Given) bool {
	lower := strings.ToLower(name.Name)
	var nonAlphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

	if len(usedTinyNames) >= 1<<14 {
		usedTinyNames = make(map[string]*Given)
	}

	if nonAlphaRegex.MatchString(lower) && usedTinyNames[lower] == nil && len(name.Name) > 2 {
		usedTinyNames[lower] = &name
		return true
	}
	return false
}
