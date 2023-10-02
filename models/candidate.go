package models

type Candidate struct {
	Filename string `json:"filename"`
	FrameID  string `json:"frame_id"`
	Score    int    `json:"score"`
}

func NewCandidate(frame Frame, score int) *Candidate {
	return &Candidate{
		Filename: frame.GetFilename(),
		FrameID:  frame.GetFrameID(),
		Score:    score,
	}
}
