package models

type Candidate struct {
	Frame *Frame
	Score int
}

func NewCandidate(frame *Frame, score int) *Candidate {
	return &Candidate{
		Frame: frame,
		Score: score,
	}
}
