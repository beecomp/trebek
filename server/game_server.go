package server

import "github.com/google/uuid"

type GameServer struct {
	Qcs [5]QuestionCategory
	Qs  map[uuid.UUID]*Question
}

func (gs *GameServer) SetQcs(qcs [5]QuestionCategory) {
	gs.Qcs = qcs
	gs.Qs = map[uuid.UUID]*Question{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			q := &gs.Qcs[i].Qs[j]
			gs.Qs[q.Id] = q
		}
	}
}

func NewGameServer(qcs [5]QuestionCategory) *GameServer {
	gs := GameServer{Qcs: qcs, Qs: map[uuid.UUID]*Question{}}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			q := &gs.Qcs[i].Qs[j]
			gs.Qs[q.Id] = q
		}
	}
	return &gs
}

type QuestionCategory struct {
	Name string      `json:"name"`
	Qs   [5]Question `json:"qs"`
}

type Question struct {
	Id uuid.UUID `json:"id"`
	Q  string    `json:"q"`
	A  string    `json:"a"`

	Revealed bool `json:"revealed"`
}
