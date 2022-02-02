package startgame

import (
	"encoding/json"
	"fmt"
	"github.com/golovpetr/internal/storage"
	"github.com/google/uuid"
	"net/http"
)

func StartGamePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = fmt.Fprint(w, "Unsupported method")
		return
	}

	defer r.Body.Close()

	var in startGameIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Incorrect input data")
		return
	}

	uid := uuid.New().String()

	storage.GameFields[uid] = in.StartField

	out := startGameOut{
		GameID: uid,
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
