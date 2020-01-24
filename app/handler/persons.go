package handler

import (
	"encoding/json"
	"net/http"

	"github.com/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllPersons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	persons := []model.Person{}
	db.Find(&persons)
	respondJSON(w, http.StatusOK, persons)
}

func CreatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	person := model.Person{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, person)
}

func GetPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func UpdatePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func DeletePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	if err := db.Delete(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisablePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	person.Disable()
	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

func EnablePerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	person := getPersonOr404(db, name, w, r)
	if person == nil {
		return
	}
	person.Enable()
	if err := db.Save(&person).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, person)
}

// getPersonOr404 gets a person instance if exists, or respond the 404 error otherwise
func getPersonOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Person {
	person := model.Person{}
	if err := db.First(&person, model.Person{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &person
}
