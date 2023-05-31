package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"myapp/model"
	"myapp/utils/httpResp"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "add student handler")
	var cour model.Course
	fmt.Println(cour)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cour)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json data")
		return
	}
	// studPointer:=&stud
	// studpointer.Create()
	dbErr := cour.Create()
	if dbErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, dbErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Course data added"})

}

func GetCour(w http.ResponseWriter, r *http.Request) {
	// get url parameter
	cid := mux.Vars(r)["cid"]
	coId, idErr := getcourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := model.Course{CId: coId}

	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

// helper function to receive string sid as a input and return a string s
func getcourseId(userIdParam string) (string, error) {
	return userIdParam, nil
}

func getCourId(courIdParam string) (string, error) {
	return courIdParam, nil
}

func UpdateCour(w http.ResponseWriter, r *http.Request) {
	old_cid := mux.Vars(r)["cid"]
	old_coId, idErr := getCourId(old_cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var cour model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cour); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	updateErr := cour.Update(old_coId)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, cour)
	}
}

// delete course
func DeleteCour(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	coId, idErr := getCourId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := model.Course{CId: coId}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// get all course
// getall (get request)
func GetAllCour(w http.ResponseWriter, r *http.Request) {
	courses, getErr := model.GetALLCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
