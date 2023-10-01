package models

type Candidate struct {
	Filename string
	FrameID  string
	Score    int
}

func NewCandidate(frame Frame, score int) *Candidate {
	return &Candidate{
		Filename: frame.GetFilename(),
		FrameID:  frame.GetFrameID(),
		Score:    score,
	}
}
