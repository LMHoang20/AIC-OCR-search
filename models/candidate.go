package models

type Candidate struct {
	Filename string  `json:"filename"`
	FrameID  string  `json:"frame_id"`
	Score    float32 `json:"score"`
}

func NewCandidate(frame Frame, score float32) *Candidate {
	return &Candidate{
		Filename: frame.GetFilename(),
		FrameID:  frame.GetFrameID(),
		Score:    score,
	}
}
