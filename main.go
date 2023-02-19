package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "root"
// 	dbname   = "tmp"
// )

type product struct {
	gorm.Model
	Name     string `gorm:"typevarchar(100);"`
	Price    float64
	Quantity int
}

// type allProducts []product

// var products = allProducts{
// 	{
// 		ID:       "1",
// 		Name:     "Introduction to Golang",
// 		Price:    9.99,
// 		Quantity: 20,
// 	},
// }
// var db *sql.DB
// var err error

// func init() {
// 	// postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 	// 	"password=%s dbname=%s sslmode=disable",
// 	// 	host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname))
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully connected!")
// }

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome home!")

// }

// func createProduct(w http.ResponseWriter, r *http.Request) {
// 	var newProduct product
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newProduct)
// 	products = append(products, newProduct)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newProduct)
// }
// func getOneProduct(w http.ResponseWriter, r *http.Request) {
// 	productID := mux.Vars(r)["id"]

// 	for _, singleProduct := range products {
// 		if singleProduct.ID == productID {
// 			json.NewEncoder(w).Encode(singleProduct)
// 		}
// 	}
// }
// func getAllProducts(w http.ResponseWriter, r *http.Request) {
// 	res, err := db.Exec("select * from Produs;")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res)
// 	json.NewEncoder(w).Encode(products)
// }

// func updateEvent(w http.ResponseWriter, r *http.Request) {
// 	productID := mux.Vars(r)["id"]
// 	var updatedProduct product

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}
// 	json.Unmarshal(reqBody, &updatedProduct)

// 	for i, singleProduct := range products {
// 		if singleProduct.ID == productID {
// 			singleProduct.Name = updatedProduct.Name
// 			singleProduct.Price = updatedProduct.Price
// 			singleProduct.Quantity = updatedProduct.Quantity
// 			products = append(products[:i], singleProduct)
// 			json.NewEncoder(w).Encode(singleProduct)
// 		}
// 	}
// }

// func deleteEvent(w http.ResponseWriter, r *http.Request) {
// 	productID := mux.Vars(r)["id"]

// 	for i, singleEvent := range products {
// 		if singleEvent.ID == productID {
// 			products = append(products[:i], products[i+1:]...)
// 			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", productID)
// 		}
// 	}
// }

var db *gorm.DB
var err error

func main() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	pass := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, pass, dbPort)

	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()
	db.AutoMigrate(&product{})

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// router.HandleFunc("/prd", createProduct).Methods("POST")
	// router.HandleFunc("/prd/{id}", getOneProduct).Methods("GET")
	// router.HandleFunc("/prds", getAllProducts).Methods("GET")
	// router.HandleFunc("/prds/{id}", updateEvent).Methods("PUT")
	// router.HandleFunc("/prds/{id}", deleteEvent).Methods("DELETE")
	// log.Fatal(http.ListenAndServe(":8080", router))

}
