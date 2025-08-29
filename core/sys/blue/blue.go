package blue

import "core/sys/num"

// "The Blue Note" is a fixed value from 0-7 improvised on startup.
func Note() num.Note {
	return num.Random[num.Note]()
}
