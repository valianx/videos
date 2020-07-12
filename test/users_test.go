package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valianx/videos/config"
	usuarios "github.com/valianx/videos/internal/domain/models/usuario"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	config.Drop()
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestUsersPOSTRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	x := usuarios.CreateUserInput{
		Nombre:   "test",
		Email:    "test@test.cl",
		Password: "123",
	}
	s, err := json.Marshal(x)
	req := bytes.NewBuffer(s)
	resp, err := http.Post(fmt.Sprintf("%s/users", ts.URL), "application/json", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestUsersGetRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/users", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
func TestUsersGetIDRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/users/1", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

//patch test
func TestUserPATCHRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()
	x := usuarios.UpdateUserInput{
		Nombre:   "test editado",
		Email:    "test@test.cl",
		Password: "123",
	}

	s, err := json.Marshal(x)

	req := bytes.NewBuffer(s)

	resp, err := PatchRequest(fmt.Sprintf("%s/users/1", ts.URL), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

/*
func TestLoginRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	x := handlers.Login{
		Email:    "test@test.cl",
		Password: "123",
	}
	s, err := json.Marshal(x)
	req := bytes.NewBuffer(s)
	resp, err := http.Post(fmt.Sprintf("%s/login", ts.URL), "application/json", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
*/

func TestUsuariosDELETERoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	resp, err := DeleteRequest(fmt.Sprintf("%s/users/1", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
