package handlers

import (
	"net/http"
	"testing"
)

var (
	getBoardUrl = "/api/"
)

func testGetBoards(t *testing.T) {
	t.Parallel()
	request, _ := http.NewRequest("GET", getBoardUrl, nil)
	cookie := &http.Cookie{Name: ""}
	request.AddCookie()
}
