package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrecodato/go-hardwarestore/pkg/models"
	"github.com/andrecodato/go-hardwarestore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewHardware models.Hardware

func GetHardware(w http.ResponseWriter, r *http.Request) {
	newHardwares := models.GetAllHardwares()
	res, _ := json.Marshal(newHardwares)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHardwareById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hardwareId := vars["hardwareId"]
	ID, err := strconv.ParseInt(hardwareId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	hardwareDetails, _ := models.GetHardwareById(ID)
	res, _ := json.Marshal(hardwareDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateHardware(w http.ResponseWriter, r *http.Request) {
	CreateHardware := &models.Hardware{}
	utils.ParseBody(r, CreateHardware)
	h := CreateHardware.CreateHardware()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteHardware(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hardwareId := vars["hardwareId"]
	ID, err := strconv.ParseInt(hardwareId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	hardware := models.DeleteHardware(ID)
	res, _ := json.Marshal(hardware)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateHardware(w http.ResponseWriter, r *http.Request) {
	var updateHardware = &models.Hardware{}
	utils.ParseBody(r, updateHardware)
	vars := mux.Vars(r)
	hardwareId := vars["hardwareId"]
	ID, err := strconv.ParseInt(hardwareId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	hardwareDetails, db := models.GetHardwareById(ID)
	if updateHardware.Name != "" {
		hardwareDetails.Name = updateHardware.Name
	}
	if updateHardware.Brand != "" {
		hardwareDetails.Brand = updateHardware.Brand
	}
	if updateHardware.Serial != "" {
		hardwareDetails.Serial = updateHardware.Serial
	}
	db.Save(&hardwareDetails)
	res, _ := json.Marshal(hardwareDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
