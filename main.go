package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

//Resp is a simple structure for the http response
type Resp struct {
	Code int
	Time string
	Data struct {
		Hostname    string
		IPs         []string
		RequestAddr string
		Host        string
	}
}

func getFunc(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/json")

	resp := Resp{Code: http.StatusOK, Time: time.Now().Format(time.RFC3339)}

	if resp.Data.Hostname, err = os.Hostname(); err != nil {
		fmt.Println("Error getting hostname")
	}

	resp.Data.Host = r.Host

	resp.Data.RequestAddr = r.RequestURI

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces")
	}

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Println("Unable to get addresses for interface", iface.Name)
		}
		for _, addr := range addresses {
			resp.Data.IPs = append(resp.Data.IPs, addr.String())
		}

	}

	j, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Unable to Marshal json")
	}
	fmt.Fprint(w, string(j))

}

func main() {
	port := flag.Int("p", 8080, "Listen port")
	flag.Parse()

	fmt.Println("Starting service ...")

	http.HandleFunc("/", getFunc)

	listen := fmt.Sprintf(":%v", *port)
	http.ListenAndServe(listen, nil)
}
