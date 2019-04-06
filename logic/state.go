package logic

import "net/http"

type State struct {
}

func NewState() *State {
	return &State{}
}

func (g *State) Handler() http.Handler {
	return nil
}
