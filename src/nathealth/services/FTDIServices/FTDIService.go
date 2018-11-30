package FTDIServices
/*
#include <stdlib.h>
#include <stdio.h>
#cgo LDFLAGS: -lftd2xx
#include <../../../../lib/ftd2xx.h>
*/
import "C"
//import "unsafe"
import "net/http"

func Connect(w http.ResponseWriter, r *http.Request){

}

func Send(w http.ResponseWriter, r *http.Request){

}

func Receive(w http.ResponseWriter, r *http.Request){

}

func Scan(w http.ResponseWriter, r *http.Request){

}

func ScanWithControlCharacter(w http.ResponseWriter, r *http.Request){

}

func Disconnect(w http.ResponseWriter, r *http.Request){

}

func GetAll(w http.ResponseWriter, r *http.Request){

}

func GetAllEscorts(w http.ResponseWriter, r *http.Request){

}

func GetDeviceCount(w http.ResponseWriter, r *http.Request) (int,int) {
	var deviceCount C.DWORD=0 // NOTE: Don't listen to the IDE when it claims this type can be omitted
	ftStatus := C.FT_CreateDeviceInfoList(&deviceCount)
	return int(ftStatus),int(deviceCount)
}

