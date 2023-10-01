package models

type Frame interface {
	GetFilename() string
	GetFrameID() string
	String() string
}
