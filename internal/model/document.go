package model

type Chunk struct {
	Index int
	Text  string
}

type ProcessedDeocument struct {
	FileName string
	Chunks   []Chunk
}
