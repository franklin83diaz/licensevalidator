package test

import (
	"fmt"
	"net/http"
)

func SeverHttp() {

	//Listen server port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//get parameter
		jwt := r.URL.Query().Get("q")

		//if the parameter is empty, return a bad request
		if jwt == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}

		dtoReq, err := CheckJWT(publicKey, jwt)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		protectedID := dtoReq.GetProtectedId()
		serialNumber := dtoReq.GetSerialNumber()

		fmt.Println("Request received in server:")
		fmt.Println("ProtectedID:", protectedID)
		fmt.Println("SerialNumber:", serialNumber)

		if protectedID == "1501d9052e005977a3a1281f8f4a689681f82da09229005889e0fcc5cf365e57" && serialNumber == "X1d23" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("License is valid"))

			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not has this serial number registered"))

	})

	//server running
	http.ListenAndServe(":8080", nil)
}
