package models

type Candidate struct {
	Filename string  `json:"video"`
	FrameID  string  `json:"frame_name"`
	Score    float32 `json:"score"`
}

func NewCandidate(frame Frame, score float32) *Candidate {
	return &Candidate{
		Filename: frame.GetFilename(),
		FrameID:  frame.GetFrameID() + ".jpg",
		Score:    score,
	}
}
