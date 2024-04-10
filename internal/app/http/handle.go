package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dragaoazuljr/ice-breaker-go/internal/app"
)

func GetIceBreakers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
	  http.Error(w, "name is required", http.StatusBadRequest)
		return
  }

	content, err := app.IceBreaker(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("llm content: ", content)

	jsonM, err := json.Marshal(content)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonM)
}
