package basic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rss/usecase"
	"strconv"
)

type (
	Application interface {
		Run(port int)
	}
	app struct {
		ucase usecase.Usecase
	}
)

func NewApplication(uc usecase.Usecase) Application {
	return &app{
		ucase: uc,
	}
}

func (a *app) Run(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		trends, err := a.ucase.GetTrends(context.Background())
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(trends)
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(response)
	})

	// ポート8080番でサーバーを起動する
	log.Println(fmt.Sprintf("http://localhost:%d", port))
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		panic(err)
	}
}
