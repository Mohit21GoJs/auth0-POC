package handlers

import (
	"encoding/json"
	"net/http"
)

type product struct {
	Id int
	Name string
	Slug string
	Description string
}

/* We will create our catalog of VR experiences and store them in a slice. */
var products = []product{
	product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description : "Shoot your way to the top on 14 different hoverboards"},
	product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind underwater experience"},
	product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
	product{Id: 4, Name: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
	product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR",
	}}

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// Here we are converting the slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})