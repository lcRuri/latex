package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// TODO generate middleware implement function, delete after code implementation
		//auth := r.Header.Get("Authorization")
		////认证为空
		//if auth == "" {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	w.Write([]byte("Unauthorized 未授权"))
		//	return
		//}
		//
		////认证出错
		//uc, err := helper.AnalyzeToken(auth)
		//if err != nil {
		//	log.Printf(" helper.AnalyzeToken(auth) err:%v", err)
		//	w.WriteHeader(http.StatusUnauthorized)
		//	w.Write([]byte(err.Error()))
		//	return
		//}
		//
		////往header设置头部信息
		//r.Header.Set("UserId", string(rune(uc.Id)))
		//r.Header.Set("UserIdentity", uc.Identity)
		//r.Header.Set("UserName", uc.Name)
		//next(w, r)
		//// Passthrough to next handler if need
		next(w, r)
	}
}
