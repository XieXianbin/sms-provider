package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/XieXianbin/msg-provider/config"
	"github.com/XieXianbin/msg-provider/sms/alidayu"
	"github.com/XieXianbin/msg-provider/sms/aliyun"
	"github.com/XieXianbin/msg-provider/smtp"
	"github.com/XieXianbin/msg-provider/wechat/mp/core"
	"github.com/XieXianbin/msg-provider/wechat/mp/message/template"
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
	name := param.MustString(r, "name")
	context := param.MustString(r, "context")

	provider := cfg.Sms.Provider

	moblieList := strings.Split(tos, ",")
	h := int(len(moblieList) / 200)
	result := make([]*alidayu.Result, 0, h)

	switch provider {
	case "alidayu":
		sms_param := []string{"{\"context\": \"", context, "\"}", "{\"name\": \"", name, "\"}"}
		param := strings.Join(sms_param, "")
		alidayu.Appkey = cfg.Sms.Appkey
		alidayu.AppSecret = cfg.Sms.Appsecret
		alidayu.IsDebug = cfg.Debug
		result, _ = alidayu.SendBatch(tos, cfg.Sms.Smssignname, cfg.Sms.Smstemplatecode, param)

		if result != nil {
			log.Println("sms send to:", tos, "result is success.")
			http.Error(w, "success", http.StatusOK)
		} else {
			log.Println("sms send to:", tos, "result is fail.")
			http.Error(w, "error", http.StatusInternalServerError)
		}

	case "aliyun":
		aliyun.HttpDebugEnable = cfg.Debug
		c := aliyun.New(cfg.Sms.Appkey, cfg.Sms.Appsecret)
		// send to more than one person
		moblieList := strings.Split(tos, ",")
		paramstring, error1 := json.Marshal(map[string]string{
			"name":    name,
			"context": context,
		})

		if error1 != nil {
			log.Println("json error:", error1)
			http.Error(w, error1.Error(), http.StatusOK)
		}

		log.Println("--", string(paramstring))

		e, err := c.SendMulti(moblieList, cfg.Sms.Smssignname, cfg.Sms.Smstemplatecode, string(paramstring))
		if err != nil {
			log.Println("sms send to:", tos, "result is failed.", err, e.Error())
			http.Error(w, e.Error(), http.StatusOK)
		} else {
			log.Println("sms send to:", tos, "result is success. request id:", e.GetRequestId())
			http.Error(w, "success", http.StatusOK)
		}

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
		log.Println("smtp send to:", tos, "result is fail:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("smtp send to:", tos, "result is success.")
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
	tos = strings.Replace(tos, ",", ";", -1)

	accessTokenServer := core.NewDefaultAccessTokenServer(cfg.Wechat.Appid, cfg.Wechat.Appsecret, nil)
	wechatClient := core.NewClient(accessTokenServer, nil)

	title := template.DataItem{
		Value: param.MustString(r, "title"),
		Color: "#173177",
	}
	endpoint := template.DataItem{
		Value: param.MustString(r, "endpoint"),
		Color: "#173177",
	}
	metric := template.DataItem{
		Value: param.MustString(r, "metric"),
		Color: "#bd2636",
	}
	tags := template.DataItem{
		Value: param.MustString(r, "tags"),
		Color: "#bd2636",
	}
	reason := template.DataItem{
		Value: param.MustString(r, "reason"),
		Color: "#173177",
	}
	max := template.DataItem{
		Value: param.MustString(r, "max"),
		Color: "#173177",
	}
	current := template.DataItem{
		Value: param.MustString(r, "current"),
		Color: "#173177",
	}
	note := template.DataItem{
		Value: param.MustString(r, "note"),
		Color: "#bd2636",
	}
	timestamp := template.DataItem{
		Value: param.MustString(r, "timestamp"),
		Color: "#173177",
	}
	remark := template.DataItem{
		Value: param.MustString(r, "remark"),
		Color: "#173177",
	}

	data := map[string]template.DataItem{
		"title":     title,
		"endpoint":  endpoint,
		"metric":    metric,
		"tags":      tags,
		"reason":    reason,
		"max":       max,
		"current":   current,
		"note":      note,
		"timestamp": timestamp,
		"remark":    remark,
	}

	msg := &template.TemplateMessage2{
		ToUser:     param.MustString(r, "tos"),
		TemplateId: cfg.Wechat.Templateid,
		URL:        param.MustString(r, "url"),
		Data:       data,
	}

	msgid, err := template.Send(wechatClient, msg)
	if err != nil {
		log.Println("wechat send to:", tos, ", msgid:", msgid, "result is err:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("wechat send to:", tos, ", msgid:", msgid, "result is success.")
		http.Error(w, "success", http.StatusOK)
	}
}

func configProcRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/sender/sms", sendSMS)
	http.HandleFunc("/sender/email", sendMail)
	http.HandleFunc("/sender/wechat", sendWechat)
}
