package nextstep

import (
	"encoding/json"
	"github.com/golovpetr/internal/storage"
	"net/http"
)

func NextStep(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var in nextStepIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uid := in.GameID

	prevFrame, ok := storage.GameFields[uid]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	curFrame := make([][]int, len(prevFrame))
	for i := range curFrame {
		curFrame[i] = make([]int, len(prevFrame[0]))
	}

	for i := 0; i < len(prevFrame); i++ {
		for j := 0; j < len(prevFrame[0]); j++ {
			counter := 0
			if i == 0 && j == 0 {
				counter += prevFrame[0][1] +
					prevFrame[1][0] +
					prevFrame[1][1]
			} else if i == 0 && j == len(prevFrame[0])-1 {
				counter += prevFrame[0][len(prevFrame[0])-2] +
					prevFrame[1][len(prevFrame[0])-1] +
					prevFrame[1][len(prevFrame[0])-2]
			} else if i == len(prevFrame)-1 && j == len(prevFrame[0])-1 {
				counter += prevFrame[len(prevFrame)-2][len(prevFrame[0])-1] +
					prevFrame[len(prevFrame)-1][len(prevFrame[0])-2] +
					prevFrame[len(prevFrame)-2][len(prevFrame[0])-2]
			} else if i == len(prevFrame)-1 && j == 0 {
				counter += prevFrame[len(prevFrame)-2][0] +
					prevFrame[len(prevFrame)-1][1] +
					prevFrame[len(prevFrame)-2][1]
			} else if i == 0 {
				counter += prevFrame[i][j-1] +
					prevFrame[i+1][j-1] +
					prevFrame[i+1][j] +
					prevFrame[i+1][j+1] +
					prevFrame[i][j+1]
			} else if j == len(prevFrame[0])-1 {
				counter += prevFrame[i-1][j] +
					prevFrame[i-1][j-1] +
					prevFrame[i][j-1] +
					prevFrame[i+1][j-1] +
					prevFrame[i+1][j]
			} else if i == len(prevFrame)-1 {
				counter += prevFrame[i][j-1] +
					prevFrame[i-1][j-1] +
					prevFrame[i-1][j] +
					prevFrame[i-1][j+1] +
					prevFrame[i][j+1]
			} else if j == 0 {
				counter += prevFrame[i-1][j] +
					prevFrame[i-1][j+1] +
					prevFrame[i][j+1] +
					prevFrame[i+1][j+1] +
					prevFrame[i+1][j]
			} else {
				counter += prevFrame[i-1][j] +
					prevFrame[i-1][j+1] +
					prevFrame[i][j+1] +
					prevFrame[i+1][j+1] +
					prevFrame[i+1][j] +
					prevFrame[i+1][j-1] +
					prevFrame[i][j-1] +
					prevFrame[i-1][j-1]
			}

			if prevFrame[i][j] == 0 && counter == 3 {
				curFrame[i][j] = 1
			} else if prevFrame[i][j] == 1 && counter == 2 || counter == 3 {
				curFrame[i][j] = 1
			} else {
				curFrame[i][j] = 0
			}

			counter = 0
		}
	}

	storage.GameFields[uid] = curFrame

	out := nextStepOut{
		GameField: storage.GameFields[uid],
	}

	outBytes, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	wrote, err := w.Write(outBytes)
	if err != nil || wrote != len(outBytes) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
