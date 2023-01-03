package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostTodo(t *testing.T) {
	assert := assert.New(t)

	todo := Todo{Name: "aaa"}
	bytes, _ := json.Marshal(todo)
	s := string(bytes)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(s)) //  /todos 경로 테스트
	mux := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(mux)
	n.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	data, _ := io.ReadAll(res.Body)
	log.Println(string(data))
}
