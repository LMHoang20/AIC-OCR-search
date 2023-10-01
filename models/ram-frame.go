package models

type RAMFrame struct {
	filename string
	frameID  int
}

func NewRAMFrame(filename string, frameID int) *RAMFrame {
	return &RAMFrame{
		filename: filename,
		frameID:  frameID,
	}
}

func (f *RAMFrame) GetFilename() string {
	return f.filename
}

func (f *RAMFrame) GetFrameID() int {
	return f.frameID
}
