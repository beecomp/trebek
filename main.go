package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/beecomp/trebek/server"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func MakePingerServer() http.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		first := true
		ema := 0 * time.Millisecond
		emv := 0 * time.Millisecond
		for {
			_ = conn.WriteMessage(websocket.TextMessage, []byte{})
			st := time.Now()

			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			lat := time.Since(st)
			if first {
				ema = lat
				first = false
			} else {
				d := lat - ema
				ema += d * 1 / 10
				emv = (emv + d*d*1/10) * 9 / 10
			}
			emstd := time.Duration(math.Sqrt(float64(emv)))
			log.Println("pong", lat, ema, emstd)

			time.Sleep(100 * time.Millisecond)
		}
	}
}

func MakeGetJeopardyBoardEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qcs := s.Qcs
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				qcs[i].Qs[j].Q = ""
				qcs[i].Qs[j].A = ""
			}
		}
		v, _ := json.Marshal(s.Qcs)
		_, _ = w.Write(v)
	}
}

func MakeResetJeopardyBoardEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, q := range s.Qs {
			q.Revealed = false
		}
	}
}

type LoadBoardRequest struct {
	Board [5]server.QuestionCategory `json:"board"`
}

func MakeLoadJeopardyBoardEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DASAR ANJENG")
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		req := LoadBoardRequest{}
		d.Decode(&req)
		log.Println(req)
		s.SetQcs(req.Board)
	}
}

func MakeGetQuestionEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := uuid.Parse(mux.Vars(r)["q"])
		if err != nil {
			return
		}
		if _, ok := s.Qs[q]; !ok {
			return
		}
		v, _ := json.Marshal(s.Qs[q])
		_, _ = w.Write(v)
	}

}

func MakeSetQuestionRevealedEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := uuid.Parse(mux.Vars(r)["q"])
		if err != nil {
			return
		}
		if _, ok := s.Qs[q]; !ok {
			return
		}
		s.Qs[q].Revealed = true

		v, _ := json.Marshal(s.Qs[q])
		_, _ = w.Write(v)
	}

}

func MakeSetQuestionUnrevealedEndpoint(s *server.GameServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := uuid.Parse(mux.Vars(r)["q"])
		if err != nil {
			return
		}
		if _, ok := s.Qs[q]; !ok {
			return
		}
		s.Qs[q].Revealed = false

		v, _ := json.Marshal(s.Qs[q])
		_, _ = w.Write(v)
	}
}

func MakeAuthMiddleware(token string) mux.MiddlewareFunc {
	token = fmt.Sprintf("Token %s", token)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Cookie("Authorization"))
			if v, err := r.Cookie("Authorization"); err == nil && v.Value == token {
				next.ServeHTTP(w, r)
				return
			}
			http.Error(w, "403 Unauthorized", http.StatusUnauthorized)
		})
	}
}

type SetTokenRequest struct {
	Token string `json:"token"`
}

func MakeSetTokenEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		req := SetTokenRequest{}
		d.Decode(&req)

		c := &http.Cookie{
			Name:  "Authorization",
			Value: fmt.Sprintf("Token %s", req.Token),
			SameSite: http.SameSiteNoneMode,
			Secure: true,
		}
		http.SetCookie(w, c)
	}
}

func main() {
	gs := server.NewGameServer(server.Jeoqcs)

	r := mux.NewRouter()

	r.Handle("/ping", MakePingerServer())
	auth := MakeAuthMiddleware(os.Getenv("TOKEN"))

	r.Handle("/q/{q}", auth(MakeGetQuestionEndpoint(gs)))
	r.Handle("/q/{q}/reveal", auth(MakeSetQuestionRevealedEndpoint(gs)))
	r.Handle("/q/{q}/unreveal", auth(MakeSetQuestionUnrevealedEndpoint(gs)))

	r.Handle("/board", auth(MakeGetJeopardyBoardEndpoint(gs)))
	r.Handle("/load-board", auth(MakeLoadJeopardyBoardEndpoint(gs)))
	r.Handle("/reset-board", auth(MakeGetJeopardyBoardEndpoint(gs)))
	r.Handle("/set-token", MakeSetTokenEndpoint())

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://jeopardy.bcomp.id"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credentialsOk := handlers.AllowCredentials()
	log.Fatal(http.ListenAndServe(":80", handlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(r)))
}
