package main

import(
	"net/http"
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Username string   `json:"username"`
	Email    string   `json:"username"`
	Password string   `json:"username"`
}

func main() {
	http.HandleFunc("/a",handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w, "Hello World!!")
	userinfo := UserInfo{
		Username: "hoge",
		Email:    "1222@gmail.com",
		Password: "password",
	}
	res, err := json.Marshal(userinfo)
	fmt.Println(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.Write(res)
	fmt.Fprint(w, userinfo)
    return
}
