package Spark

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
FileMetaData
*/

type FileMetaDataSet struct {
	Logic.Components[FileMetaData]
}

type FileMetaData struct {
	Path      string
	Name      string
	Extension string
}

/**
Renderable
*/

type RenderableSet struct {
	Logic.Components[Renderable]
}

type Renderable struct {
}
