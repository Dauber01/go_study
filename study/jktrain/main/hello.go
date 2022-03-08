package main

import (
	"bytes"
	"strings"

	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.RequestURI)
		ans := r.Method
		log.Printf(ans)
		s, _ := ioutil.ReadAll(r.Body)
		log.Printf("%s", s)
		r.Body.Close()
		/* var v interface{}
		json.Unmarshal(s, &v)
		data := v.(map[string]interface{})
		log.Printf("%v", data) */
		url := "http://apiin.im.baidu.com/api/msg/groupmsgsend?access_token=d181c1f04fa3e78c253fd2d61a691d5a3"
		mm := `{"message": {"body":`
		nn := `}}`
		s = []byte(strings.Join([]string{mm, string(s), nn}, ""))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(s))
		if err != nil {
			log.Printf("%v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("%v", err)
		}
		k, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s", k)
		resp.Body.Close()
		//fmt.Println(w, "Hello, GopherCon 5G")
	})
	/* 	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}() */

	//对于相互之间有关联的调用，使用串行之行
	if err := http.ListenAndServe(":8111", nil); err != nil {
		log.Fatal(err)
	}

	//用来对整个函数进行阻塞，防止函数执行完毕
	//select {}
}
