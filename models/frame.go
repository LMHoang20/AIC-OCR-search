package models

type Frame struct {
	filename string
	frameID  int
}

func NewFrame(filename string, frameID int) *Frame {
	return &Frame{
		filename: filename,
		frameID:  frameID,
	}
}

func (f *Frame) Filename() string {
	return f.filename
}

func (f *Frame) FrameID() int {
	return f.frameID
}
