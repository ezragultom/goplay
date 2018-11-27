package main

import (
	"dockerize/conn"
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	n, _ := io.WriteString(w, "Hello world!")
	fmt.Println("hello", n)
}

func set(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	query := r.URL.Query().Get("query")

	redis := conn.RedisClient("master")
	redisHandler := conn.NewRedisHandler(redis)
	val, err := redisHandler.Set(key, query)
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("SET! key = %s AND query = %s AND %v", key, query, val)
	}
	n, _ := io.WriteString(w, msg)
	fmt.Println("SET", n)
}

func get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	redis := conn.RedisClient("master")
	redisHandler := conn.NewRedisHandler(redis)
	val, err := redisHandler.Get(key)
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("GET! key = %s AND value = %s", key, val)
	}
	n, _ := io.WriteString(w, msg)
	fmt.Println("SET", n)
}

func main() {
	fmt.Println("SERVER START")
	http.HandleFunc("/", hello)
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.ListenAndServe("0.0.0.0:8000", nil)

}
