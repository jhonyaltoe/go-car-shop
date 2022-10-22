package routes

import (
	"github.com/gorilla/mux"
	carController "github.com/jhony/go-car-shop/api/controllers/car"
	connect "github.com/jhony/go-car-shop/api/database"
)

func CarRoutes() *mux.Router {
	carCollection := connect.NewCollection("CarShop", "cars")
	carCtrl := carController.New(carCollection)

	router := mux.NewRouter()
	router.Host("http://www.google.com")
	router.HandleFunc("/cars", carCtrl.GetAll).Methods("GET")
	router.HandleFunc("/cars", carCtrl.CreateOne).Methods("POST")

	return router
}
