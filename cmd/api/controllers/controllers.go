package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oluu-web/hngx-stage2/cmd/api/models"
	"github.com/oluu-web/hngx-stage2/cmd/api/utilities"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	// validation checks
	err = models.ValidateFields(person)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	isDuplicate, err := models.CheckDuplicate(person.Name)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}
	if isDuplicate {
		utilities.ErrorJSON(w, fmt.Errorf("duplicate person"))
		return
	}

	_, err = models.CreateNewPerson(person)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, person, "person")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")
	person, err := models.GetPerson(name)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, person, "person")
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")

	var updatedPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	isDuplicate, err := models.CheckDuplicate(updatedPerson.Name)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}
	if isDuplicate {
		utilities.ErrorJSON(w, fmt.Errorf("duplicate person"))
		return
	}

	err = models.UpdatePerson(name, updatedPerson)
	if err != nil {
		utilities.ErrorJSON(w, err)
	}

	utilities.WriteJSON(w, http.StatusOK, "Updated Successfully", "Success")
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")

	err := models.DeletePerson(name)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, "Person deleted successfully", "Success")
}
