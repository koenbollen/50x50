package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/koenbollen/50x50/logic"
	"github.com/koenbollen/50x50/storage"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

const gridSize = 50
const sessionLifetime = 1 * time.Minute

type request struct {
	Session string `json:"session"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}

type response struct {
	Session string `json:"session"`
	Grid    []int  `json:"grid"`
	Width   int    `json:"width"`
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	store := storage.NewMemoryStore(sessionLifetime)

	http.HandleFunc("/state/create", func(w http.ResponseWriter, r *http.Request) {
		session := store.CreateSession()
		grid := logic.NewGrid(gridSize, gridSize)
		err := store.StoreGrid(session, grid.Data)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(&response{
			Session: session,
			Grid:    grid.Data,
			Width:   grid.Width,
		})
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/state/move", func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			panic(err)
		}

		data, err := store.LoadGrid(req.Session)
		if err != nil {
			panic(err)
		}
		grid := logic.FromData(gridSize, data)
		grid.Increment(req.X, req.Y)

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&response{
			Session: req.Session,
			Grid:    grid.Data,
			Width:   grid.Width,
		})
		if err != nil {
			panic(err)
		}
	})

	addr := ":8080"
	logger.Info("listening", zap.String("addr", addr))
	err = http.ListenAndServe(addr, cors.AllowAll().Handler(http.DefaultServeMux))
	logger.Fatal("failed to setup http server", zap.Error(err))
}
