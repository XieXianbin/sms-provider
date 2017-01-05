package http

import (
	"net/http"
	"strings"

	"github.com/XieXianbin/sms-provider/config"
	"github.com/XieXianbin/sms-provider/sms/alidayu"
	"github.com/toolkits/web/param"
)

func configProcRoutes() {

	http.HandleFunc("/sender/sms", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", ";", -1)

		provider := cfg.Sms.Provider

		moblieList := strings.Split(tos, ",")
		h := int(len(moblieList) / 200)
		result := make([]*alidayu.Result, 0, h)

		switch provider {
		case "alidayu":
			sms_param := []string{"{\"content\": \"", content, "\"}"}
			param := strings.Join(sms_param, "")
			result, _ = alidayu.SendBatch(tos, cfg.Sms.Smsfreesignname, cfg.Sms.Smstemplatecode, param)
		}

		if result != nil {
			http.Error(w, "success", http.StatusOK)
		} else {
//			http.Error(w, result, http.StatusInternalServerError)
			http.Error(w, "error", http.StatusInternalServerError)
		}
	})

}
