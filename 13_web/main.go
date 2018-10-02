// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(r)
// 	fmt.Fprintf(w, "<h1>Hello! Matta!!</h1>")
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	// urlParams := r.URL.Query()
// 	name := strings.TrimPrefix(r.URL.Path, "/greet/")
// 	fmt.Fprintf(w, "<h1>Welocome!"+name+"!!</h1>")
// }
// func main() {
// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/greet/", greet)
// 	// http.HandleFunc("/")
// 	fmt.Println("Server is starting..")
// 	http.ListenAndServe(":3000", nil)
// 	fmt.Println("Served is running on port 3000")
// }
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather -- This is a sample wather interface
type Weather struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LallLong     string `json:"latt_long"`
}

func (we *Weather) Write(w io.Writer) {
	b, _ := json.Marshal(*we)
	w.Write(b)
}

// type Test struct {
// 	Weather []
// 	Status string
// }

type Test []Weather

func index(res http.ResponseWriter, req *http.Request) {

	// data := []Weather{}
	var data []Weather
	res.Header().Set("Content-Type", "application/json")

	url := "https://www.metaweather.com/api/location/search/?query=san"

	// Building the request
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		log.Fatal("Read Error: ", readErr)
		return
	}

	defer response.Body.Close()

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(data)

	// for _, d := range data {
	// 	fmt.Println(d)
	// }
	c, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(c))
	// k := json.Marshal(data)
	res.Write(c)
}
func main() {

	http.HandleFunc("/", index)

	fmt.Println("Server is starting..")

	http.ListenAndServe(":5000", nil)

}
