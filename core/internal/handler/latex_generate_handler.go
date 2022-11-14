package handler

import (
	"net/http"

	"core/internal/logic"
	"core/internal/svc"
	"core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LatexGenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.LatexGenerateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLatexGenerateLogic(r.Context(), svcCtx)
		//file, _, err := r.FormFile("file")
		//if err != nil {
		//	log.Printf("r.FormFile(file) err:%v", err)
		//	return
		//}
		//
		//content, err := ioutil.ReadAll(file)
		//if err != nil {
		//	log.Fatal("ioutil.ReadAll(file)", err)
		//}

		resp, err := l.LatexGenerate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
