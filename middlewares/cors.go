package middlewares

import "net/http"

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if req.Method == http.MethodOptions {
			w.Write([]byte("allowed"))
			return
		}

		next(w, req)
	}

}
