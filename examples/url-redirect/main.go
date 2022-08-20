package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://kite.trade/connect/login?api_key=exas5p7qo7tula9z&v=3", 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// func main() {
// 	client := &http.Client{}
// 	url := "https://kite.trade/connect/login?api_key=exas5p7qo7tula9z&v=3"
// 	http.Redirect(w, r, url, 301)
// }

// reqest, err := http.NewRequest("GET", url, nil)
// if err != nil {
// 	panic(err)
// }
// response, _ := client.Do(reqest)
// fmt.Println(response.Status)
// fmt.Println(response.Location())

// res, err := client.Head(url)
// fmt.Printf(" %v\n", res)

// }
