package controllers

import (
	"OCRsearch/helpers"
	"context"
	"encoding/csv"
	"net/http"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/gorilla/mux"
)

type Frame struct{}

var frame *Frame

func FrameInstance() *Frame {
	if frame == nil {
		frame = &Frame{}
	}
	return frame
}

type MetaFrame struct {
	Folder    string `json:"folder"` // keyframe / subframe
	FrameName string `json:"frame_name"`
	PTSTime   string `json:"pts_time"`
	FrameID   int    `json:"frame_id"`
}

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type ReturnFormat struct {
	Folder    string `json:"folder"` // keyframe / subframe
	FrameName string `json:"frame_name"`
	RangeTime Range  `json:"range_time"`
	FrameID   string `json:"frame_id"`
}

func ReadFolder(folder string, filename string) ([]MetaFrame, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	r, err := client.Bucket("thangtd1").Object(folder + "/" + filename + ".csv").NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	csvReader := csv.NewReader(r)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = 4

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make([]MetaFrame, 0)

	header := true

	for _, record := range records {
		if header {
			header = false
			continue
		}

		id, _ := strconv.Atoi(record[3])
		frameName := record[0]
		size := 4
		videoNumber, _ := strconv.Atoi(filename[1:3])

		if videoNumber >= 17 || folder[4:] == "subframes" {
			size = 3
		}

		for len(frameName) < size {
			frameName = "0" + frameName
		}

		frameName = frameName + ".jpg"

		metaFrame := MetaFrame{
			Folder:    folder[4:],
			FrameName: frameName,
			PTSTime:   record[1],
			FrameID:   id,
		}
		result = append(result, metaFrame)
	}

	return result, nil
}

func GetMetadata(filename string) ([]MetaFrame, error) {
	keyframes, err := ReadFolder("map-keyframes", filename)
	if err != nil {
		return nil, err
	}
	subframes, err := ReadFolder("map-subframes", filename)
	if err != nil {
		return nil, err
	}

	result := make([]MetaFrame, 0)

	for i, j := 0, 0; i < len(keyframes) || j < len(subframes); {
		if i == len(keyframes) {
			result = append(result, subframes[j])
			j++
			continue
		}
		if j == len(subframes) {
			result = append(result, keyframes[i])
			i++
			continue
		}
		if keyframes[i].FrameID < subframes[j].FrameID {
			result = append(result, keyframes[i])
			i++
		} else {
			result = append(result, subframes[j])
			j++
		}
	}

	return result, nil
}

func FormatResult(mergedFrames []MetaFrame) []ReturnFormat {
	result := make([]ReturnFormat, 0)

	for i, st := 0, "0"; i < len(mergedFrames); i++ {
		returnFormat := ReturnFormat{
			Folder:    mergedFrames[i].Folder,
			FrameName: mergedFrames[i].FrameName,
			FrameID:   strconv.Itoa(mergedFrames[i].FrameID),
			RangeTime: Range{
				Start: st,
				End:   mergedFrames[i].PTSTime,
			},
		}
		st = mergedFrames[i].PTSTime
		result = append(result, returnFormat)
	}

	return result
}

func (c *Frame) FrameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	video := vars["video"]

	mergedFrames, err := GetMetadata(video)

	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	result := FormatResult(mergedFrames)

	helpers.SendResponse(w, http.StatusOK, "OK", result)
}

func GetAnswer(result []ReturnFormat, start float64, end float64) *ReturnFormat {
	var answer *ReturnFormat
	var closestRight *ReturnFormat
	var closestLeft *ReturnFormat

	for i := 0; i < len(result); i++ {
		frameTime, _ := strconv.ParseFloat(result[i].RangeTime.End, 32)

		if start <= frameTime {
			if frameTime <= end {
				answer = &result[i]
				break
			} else if closestRight == nil {
				closestRight = &result[i]
				break
			}
		} else {
			closestLeft = &result[i]
		}
	}

	if answer == nil {
		if closestLeft == nil {
			answer = closestRight
		} else if closestRight == nil {
			answer = closestLeft
		} else {
			leftTime, _ := strconv.ParseFloat(closestLeft.RangeTime.End, 32)
			rightTime, _ := strconv.ParseFloat(closestRight.RangeTime.End, 32)
			if start-leftTime < rightTime-start {
				answer = closestLeft
			} else {
				answer = closestRight
			}
		}
	}

	return answer
}

func (c *Frame) FrameHandler2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	video := vars["video"]
	start, _ := strconv.ParseFloat(vars["start"], 32)
	end, _ := strconv.ParseFloat(vars["end"], 32)

	mergedFrames, err := GetMetadata(video)

	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	result := FormatResult(mergedFrames)

	answer := GetAnswer(result, start, end)

	helpers.SendResponse(w, http.StatusOK, "OK", *answer)
}
