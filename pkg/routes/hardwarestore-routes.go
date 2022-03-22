package routes

import (
	"github.com/andrecodato/go-hardwarestore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterHardwareStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/hardware/", controllers.CreateHardware).Methods("POST")
	router.HandleFunc("/hardware/", controllers.GetHardware).Methods("GET")
	router.HandleFunc("/hardware/{hardwareId}", controllers.GetHardwareById).Methods("GET")
	router.HandleFunc("/hardware/{hardwareId}", controllers.UpdateHardware).Methods("PUT")
	router.HandleFunc("/hardware/{hardwareId}", controllers.DeleteHardware).Methods("DELETE")

}
