package main

import (
	"log"
	"net/http"
	"strings"
	"github.com/mdillenk/gsm/src/nathealth/services/FTDIServices"
	"github.com/mdillenk/gsm/src/nathealth/services/ISerialServices"
)



func fDTIController(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, "/api/v1/serial/ftdi") {
		var index = strings.Index(r.RequestURI, "v1/serial/ftdi") + 14
		runes := []rune(r.RequestURI)
		operationName := string(runes[index:len(r.RequestURI)])

		switch operationName {
		case "/connect":
			FTDIServices.Connect(w,r)
		case "/send":
			FTDIServices.Send(w,r)
		case "/receive":
			FTDIServices.Receive(w,r)
		case "/scan":
			FTDIServices.Scan(w,r)
		case "/scanwithcontrolcharacter":
			FTDIServices.ScanWithControlCharacter(w,r)
		case "/disconnect":
			FTDIServices.Disconnect(w,r)
		case "/escorts":
			FTDIServices.GetAllEscorts(w,r)
		case "/count":
			FTDIServices.GetDeviceCount(w, r)
		default:
			FTDIServices.GetAll(w,r)


		}

	}
}

func iSerialController(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, "v1/serial/iserial") {
		var index = strings.Index(r.RequestURI, "v1/serial/iserial") + 17
		runes := []rune(r.RequestURI)
		operationName := string(runes[index:len(r.RequestURI)])


		switch operationName {
			case "/connect":
				ISerialServices.Connect(w,r)
			case "/send":
				ISerialServices.Send(w,r)
			case "/receive":
				ISerialServices.Receive(w,r)
			case "/scan":
				ISerialServices.Scan(w,r)
			case "/scanwithcontrolcharacter":
				ISerialServices.ScanWithControlCharacter(w,r)
			case "/disconnect":
				ISerialServices.Disconnect(w,r)
			default:
				ISerialServices.GetAll(w,r)
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s: %s", "Fatal Error! ", msg, err)
	}
}

func main() {

	httpMux := http.NewServeMux()

	httpMux.Handle("/api/v1/serial/ftdi", http.HandlerFunc(fDTIController))
	httpMux.Handle("/api/v1/serial/iserial", http.HandlerFunc(iSerialController))

}
