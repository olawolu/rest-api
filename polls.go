package main

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// poll defines the structure of a poll with 5 fields
type poll struct {
	ID      bson.ObjectId  `bson:"_id" json:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
	APIKey  string         `json:"apikey"` // shouldn't be done in production
}

func (s *Server) handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handlePollsGet(w, r)
		return
	case "POST":
		s.handlePollsPost(w, r)
		return
	case "DELETE":
		s.handlePollsDelete(w, r)
		return
	}
	// not found
	respondHTTPErr(w, r, http.StatusNotFound)
}

// Reading polls
func (s *Server) handlePollsGet(w http.ResponseWriter, r *http.Request) {
	var q *mgo.Query
	var result []*poll

	// Create copy of the database connection
	session := s.db.Copy()
	defer session.Close()

	// create an object referring to the polls collection
	c := session.DB("ballots").C("polls")
	p := NewPath(r.URL.Path)

	// build an mgo.Query object by parsing the path
	if p.HasID() {
		// get specific poll
		q = c.FindId(bson.ObjectIdHex(p.ID))
	} else {
		// get all polls
		q = c.Find(nil)
	}
	if err := q.All(&result); err != nil {
		respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
		return
	}
	respond(w, r, http.StatusOK, &result)
}

func (s *Server) handlePollsPost(w http.ResponseWriter, r *http.Request) {

	// create a copy of the database connection
	session := s.db.Copy()
	defer session.Close()

	// create object refer

	respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
}

func (s *Server) handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("not implemented"))
}
