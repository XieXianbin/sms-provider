package http

import (
	"net/http"
	"strings"

	"github.com/XieXianbin/sms-provider/config"
	"github.com/XieXianbin/sms-provider/sms/alidayu"
	"github.com/toolkits/smtp"
	"github.com/toolkits/web/param"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this server support to send sms/email/wechat message..."))
}

func sendSMS(w http.ResponseWriter, r *http.Request) {
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
		alidayu.Appkey = cfg.Sms.Appkey
		alidayu.AppSecret = cfg.Sms.Appsecret
		alidayu.IsDebug = cfg.Debug
		result, _ = alidayu.SendBatch(tos, cfg.Sms.Smsfreesignname, cfg.Sms.Smstemplatecode, param)
	}

	if result != nil {
		http.Error(w, "success", http.StatusOK)
	} else {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	cfg := config.Config()
	token := param.String(r, "token", "")
	if cfg.Http.Token != token {
		http.Error(w, "no privilege", http.StatusForbidden)
		return
	}

	tos := param.MustString(r, "tos")
	subject := param.MustString(r, "subject")
	content := param.MustString(r, "content")
	tos = strings.Replace(tos, ",", ";", -1)

	s := smtp.New(cfg.Smtp.Addr, cfg.Smtp.Username, cfg.Smtp.Password)
	err := s.SendMail(cfg.Smtp.From, tos, subject, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Error(w, "success", http.StatusOK)
	}
}

func sendWechat(w http.ResponseWriter, r *http.Request) {
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
		//http.Error(w, result, http.StatusInternalServerError)
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func configProcRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/sender/sms", sendSMS)
	http.HandleFunc("/sender/email", sendMail)
	http.HandleFunc("/sender/wechat", sendWechat)
}
