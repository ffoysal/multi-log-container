package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	logDir = "logs"
	path1  = "logs/log1.log"
	path2  = "logs/log2.log"
	path3  = "logs/log3.log"
	path4  = "logs/log4.log"
)

func writeToFile(fileName string, data string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		t := time.Now()
		fmt.Fprintf(w, "Welcome to Ottawa Docker Meetup, "+r.URL.Query().Get("name"))
		writeToFile(path1, t.Format("2006/01/02 15:04:05.999Z ")+" [INFO] "+r.RemoteAddr+" First Log - "+r.URL.Query().Get("name")+"\n")
		writeToFile(path2, t.Format("2006/01/02 15:04:05.999Z ")+" [INFO] "+r.RemoteAddr+" Second Log - "+r.URL.Query().Get("name")+"\n")
		writeToFile(path3, t.Format("2006/01/02 15:04:05.999Z ")+" [INFO] "+r.RemoteAddr+" Third Log - "+r.URL.Query().Get("name")+"\n")
		writeToFile(path4, t.Format("2006/01/02 15:04:05.999Z ")+" [INFO] "+r.RemoteAddr+" Fourth Log - "+r.URL.Query().Get("name")+"\n")
	} else {
		fmt.Fprintf(w, "Welcome to Ottawa Docker Meetup")
	}

}

func main() {
	os.Mkdir("logs", os.ModePerm)
	http.HandleFunc("/", helloWorld)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Start and Using port:", listener.Addr().(*net.TCPAddr).Port)

	panic(http.Serve(listener, nil))
}
