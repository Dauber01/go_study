package main

import (
	"fmt"
	"log"
	"net/http"
)

func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, GopherCon 5G")
	})
	/* 	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}() */

	//对于相互之间有关联的调用，使用串行之行
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	//用来对整个函数进行阻塞，防止函数执行完毕
	//select {}
}
