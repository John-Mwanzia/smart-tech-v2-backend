package protect

import "net/http"

func Protect(next http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) 

}


