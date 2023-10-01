package models

type RAMFrame struct {
	filename string
	frameID  string
}

func NewRAMFrame(filename string, frameID string) *RAMFrame {
	return &RAMFrame{
		filename: filename,
		frameID:  frameID,
	}
}

func NewRAMFrameFromString(s string) *RAMFrame {
	return &RAMFrame{
		filename: s[:8],
		frameID:  s[9:],
	}
}

func (f *RAMFrame) GetFilename() string {
	return f.filename
}

func (f *RAMFrame) GetFrameID() string {
	return f.frameID
}

func (f *RAMFrame) String() string {
	return f.filename + ":" + f.frameID
}
