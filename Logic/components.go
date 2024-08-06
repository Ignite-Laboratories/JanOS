package Logic

/**
BinaryData
*/

type BinaryDataSet struct {
	Components[BinaryData]
}

type BinaryData struct {
	Data []byte
}

/**
FileMetaData
*/

type FileMetaDataSet struct {
	Components[FileMetaData]
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
	Components[Renderable]
}

type Renderable struct {
}
