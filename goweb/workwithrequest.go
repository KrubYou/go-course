package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Course struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id":1,
			"name":"Golang",
			"price":2999,
			"instructor":"Mr.A"
		},
		{
			"id":2,
			"name":"Python",
			"price":3999,
			"instructor":"Mr.B"
		},
		{
			"id":3,
			"name":"Java",
			"price":4999,
			"instructor":"Mr.C"
		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList {
		if highestID < course.Id {
			highestID = course.Id
		}
	}

	return highestID + 1
}

func main() {

	http.HandleFunc("/course", CourseHandler)
	http.ListenAndServe(":8000", nil)
}
