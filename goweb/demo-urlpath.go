package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var courseList []Course

func init() {
	courseJSON := `[
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
	err := json.Unmarshal([]byte(courseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range courseList {
		if highestID < course.Id {
			highestID = course.Id
		}
	}

	return highestID + 1
}

func findID(ID int) (*Course, int) {
	for i, course := range courseList {
		if course.Id == ID {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("course ID %d not found", ID), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPut:
		var updatecourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &updatecourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatecourse.Id != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		course = &updatecourse
		courseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}

}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newcourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, &newcourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newcourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newcourse.Id = getNextID()
		courseList = append(courseList, newcourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course/", courseHandler)
	http.HandleFunc("/course", coursesHandler)
	http.ListenAndServe(":8000", nil)
}
