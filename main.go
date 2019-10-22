/**
 *  R Digger Media	(https://rdigger.com)
 *
 *
 *  @copyright      2019-present R Digger Media <info@rdigger.com>
 *                  2019-present Appzaib				<info@appzaib.com>
 *	                2019-present Nerdspex				<info@nerdspex.com>
 *	                2019-present Rana Jahanzaib	<iam@ranajahanzaib.com>
 *
 *
 *
 *
 *
 *
 *
 *
 *  @author         Rana Jahanzaib	<iam@ranajahanzaib.com>
 *
 *  @docs           https://docs.appzaib.com/rdigger/api
 *
 *  @blueprints     https://lab.appzaib.com/org/rdigger/api/blueprints
 *
 *  @package       rdigger/api    (https://pkg.appzaib.com/rdigger/api)
 */


package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	nspx "github.com/nspx/core"
	"google.golang.org/appengine"
)

// Init Products var as a slice Product struct
var products []nspx.Product

// Get all Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(products)
}

// Get a single Product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Loop through Products and find with ID
	for _, item := range products {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(&nspx.Product{})

}

func main() {

	// Init Firestore

	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "rdiggerapi"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Store sample data
	// _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	// 	"first": "Ada",
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", err)
	// }

	// Start router
	r := mux.NewRouter()

	// Sample Data
	products = append(products, nspx.Product{ID: "1", Name: "HP 840 G3 Laptop"})

	// Endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

	appengine.Main()

}
