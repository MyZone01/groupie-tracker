package main

import (
	handlers "groupietracker/lib/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSiteAvailable(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.Index(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestArtistList(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/artists", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.ArtistList(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestArtistInfos(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/artist/5", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.ArtistInfos(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestRouteNotExist(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/toto", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.ArtistList(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", res.Status)
	}
}

func TestRouteGetArtistInfoWithInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/artist/100", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.ArtistInfos(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Bad Request; got %v", res.Status)
	}
}

func TestRouteGetArtistInfoWithNoID(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/artist", nil)
	if err != nil {
		t.Fatalf("Could not created request : %v", err)
	}
	rec := httptest.NewRecorder()
	handlers.ArtistInfos(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Bad Request; got %v", res.Status)
	}
}
