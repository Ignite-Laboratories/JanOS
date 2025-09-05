package given

import (
	"core/enum/gender"
	"core/sys/given/format"
	"core/sys/id"
	log "core/sys/log"
	_ "embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"
)

// TODO: Allow replacing nameDB and surnameDB with a file from atlas configuration

var moduleName = "name"

//go:embed nameDB.tsv
var nameDBRaw string
var nameDB = make([]Name, 0, 8888)

//go:embed surnameDB.txt
var surnameDBRaw string
var surnameDB = make([]string, 0, 8888)

var tinyNameCount = 0

func init() {
	initNameDB()
	initSurnameDB()

	for _, g := range nameDB {
		lower := strings.ToLower(g.Name)
		if len(lower) > 2 && nonAlphaRegex.MatchString(lower) {
			tinyNameCount++
		}
	}

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

		entry := Name{
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

// New creates a new given.Name.  You may optionally provide a description during creation.
func New(name string, description ...string) Name {
	if len(description) > 0 {
		return Name{
			Name:        name,
			Description: description[0],
		}
	}
	return Name{
		Name: name,
	}
}

type tokenSet struct {
	sync.Mutex
	used map[Name]struct{}
}

var tokens = make(map[uint64]*tokenSet)
var lock sync.Mutex

func getTokenSet(t uint64) *tokenSet {
	lock.Lock()
	defer lock.Unlock()

	ts := tokens[t]
	if ts == nil {
		ts = &tokenSet{used: make(map[Name]struct{})}
		tokens[t] = ts
	}
	return ts
}

// Random generates a random formatted name guaranteed to be unique to the token number.  If no token is provided,
// a random token is generated along with the name.
//
// NOTE: Uniqueness is only guaranteed up to the available data set per token before "rolling over" the unique entry table to a fresh one.
func Random[T format.Format](token ...uint64) (Name, uint64) {
	var t uint64
	if len(token) > 0 {
		t = token[0]
	} else {
		t = id.Next()
	}
	ts := getTokenSet(t)

	var uniqueness int
	switch any(T("")).(type) {
	case format.NameDB:
		if len(nameDB) == 0 {
			panic("nameDB is empty")
		}
		uniqueness = len(nameDB)
	case format.SurnameDB:
		if len(nameDB) == 0 {
			panic("surnameDB is empty")
		}
		uniqueness = len(surnameDB)
	case format.Tiny:
		if len(nameDB) == 0 {
			panic("nameDB is empty")
		}
		uniqueness = int(float64(tinyNameCount) * 0.9)
	case format.Multi, format.Default: // NOTE: Default can be moved between case statements
		// NOTE: We limit this down 'slightly' from the full width for a slight performance boost under a VERY intermittent loading condition.
		uniqueness = int(float64(len(nameDB)*len(surnameDB)) * 0.9)
	default:
		panic("unknown name format")
	}

	for {
		name := random[T]()

		ts.Lock()
		if len(ts.used) >= uniqueness {
			ts.used = make(map[Name]struct{})
		}
		if _, exists := ts.used[name]; !exists {
			ts.used[name] = struct{}{}
			ts.Unlock()
			return name, t
		}
		ts.Unlock()
	}

}

var nonAlphaRegex = regexp.MustCompile(`^[A-Za-z]+$`)

func random[T format.Format]() Name {
	switch any(T("")).(type) {
	case format.NameDB:
		return nameDB[rand.Intn(len(nameDB))]
	case format.SurnameDB:
		return Name{Name: surnameDB[rand.Intn(len(surnameDB))]}
	case format.Tiny:
		for {
			name := random[format.NameDB]()
			lower := strings.ToLower(name.Name)
			if nonAlphaRegex.MatchString(lower) && len(name.Name) > 2 {
				return name
			}
		}
	case format.Multi, format.Default: // NOTE: Default can be moved between case statements
		name := nameDB[rand.Intn(len(nameDB))]
		last := surnameDB[rand.Intn(len(surnameDB))]
		name.Name += " " + last
		return name
	default:
		// Just return a random name from the NameDB
		return random[format.NameDB]()
	}
}

// Lookup finds the provided name in the provided database, otherwise it returns nil and an error.  You may optionally
// provide whether the search should be case sensitive.
//
// NOTE: This will only look up names from the NameDB and SurnameDB databases as all others are dynamically generated.
//
// See Format.
func Lookup[T format.Format](name string, caseInsensitive ...bool) (Name, error) {
	switch any(T("")).(type) {
	case format.NameDB:
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
	case format.SurnameDB:
		for _, n := range surnameDB {
			if len(caseInsensitive) > 0 && caseInsensitive[0] {
				if strings.EqualFold(n, name) {
					return Name{Name: n}, nil
				}
			} else {
				if n == name {
					return Name{Name: n}, nil
				}
			}
		}
	}
	return Name{}, fmt.Errorf("name not found")
}
