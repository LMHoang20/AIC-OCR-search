package models

type Frame interface {
	GetFilename() string
	GetFrameID() int
}
