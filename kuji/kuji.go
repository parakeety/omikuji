package kuji

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var results []string = []string{"大吉", "中吉", "小吉", "凶", "大凶"}

type Omikuji struct {
	Message string `json:"message"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func kujiPicker(i int, t time.Time) *Omikuji {
	var message string
	if isLuckyDay(t) {
		message = "大吉"
	} else {
		message = results[i]
	}

	return &Omikuji{
		Message: message,
	}
}

func isLuckyDay(t time.Time) bool {
	if t.Month() != 1 {
		return false
	}

	switch t.Day() {
	case 1, 2, 3:
		return true
	default:
		return false
	}
}

type OmikujiHandler func(w http.ResponseWriter, r *http.Request, i int, t time.Time)

func NewOmikujiHandler() OmikujiHandler {
	return omikuji
}

func omikuji(w http.ResponseWriter, r *http.Request, i int, t time.Time) {
	kuji := kujiPicker(i, t)
	b, err := json.Marshal(kuji)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}

func (f OmikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(results))
	t := time.Now()
	f(w, r, i, t)
}
