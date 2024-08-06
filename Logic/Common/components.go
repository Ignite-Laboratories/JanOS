package Common

import "github.com/Ignite-Laboratories/JanOS/Logic"

/**
BinaryData
*/

type BinaryDataSet struct {
	Logic.Components[BinaryData]
}

type BinaryData struct {
	Data []byte
}

/**
FileData
*/

type FileDataSet struct {
	Logic.Components[FileData]
}

type FileData struct {
	Path      string
	Name      string
	Extension string
}
