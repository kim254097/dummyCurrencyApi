package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kim254097/dummyApiCurrency/db"
)

var currencyList = db.AllCurrency{
	{
		ID:      "ARG",
		Name:    "Peso",
		Country: "Argentina",
	},
	{
		ID:      "USD",
		Name:    "Dolar",
		Country: "Estados Unidos",
	},
	{
		ID:      "CNY",
		Name:    "Yuan",
		Country: "China",
	},
	{
		ID:      "INR",
		Name:    "Rupia",
		Country: "India",
	},
}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var newCurrency db.Currency
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the currency name and country only in order to create")
	}

	json.Unmarshal(reqBody, &newCurrency)
	currencyList = append(currencyList, newCurrency)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCurrency)
}

func GetOneCurrency(w http.ResponseWriter, r *http.Request) {
	currencyID := mux.Vars(r)["id"]

	for _, singleCurrency := range currencyList {
		if singleCurrency.ID == currencyID {
			json.NewEncoder(w).Encode(singleCurrency)
		}
	}
}

func GetAllCurrency(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(currencyList)
}

func UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	currencyID := mux.Vars(r)["id"]
	var updatedCurrency db.Currency

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the currency name and country only in order to update")
	}
	json.Unmarshal(reqBody, &updatedCurrency)

	for i, singleCurrency := range currencyList {
		if singleCurrency.ID == currencyID {
			singleCurrency.Name = updatedCurrency.Name
			singleCurrency.Country = updatedCurrency.Country
			currencyList = append(currencyList[:i], singleCurrency)
			json.NewEncoder(w).Encode(singleCurrency)
		}
	}
}

func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	currencyID := mux.Vars(r)["id"]

	for i, singleCurrency := range currencyList {
		if singleCurrency.ID == currencyID {
			currencyList = append(currencyList[:i], currencyList[i+1:]...)
			fmt.Fprintf(w, "The currency with ID %v has been deleted successfully", currencyID)
		}
	}
}
