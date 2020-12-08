package main

import (
    "net/http/httptest"
    "net/http"
    "testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    server := NewPlayerServer(NewInMemoryPlayerStore())
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    response := httptest.NewRecorder()
    server.ServeHTTP(response, newGetScoreRequest(player))
    assertStatus(t, response.Code, http.StatusOK)

    assertResponseBody(t, response.Body.String(), "3")
}
