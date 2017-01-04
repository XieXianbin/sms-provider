package alidayu

import (
	"strings"
	"time"
)

// 设置参数常量
const (
	methodSendSMS    = "alibaba.aliqin.fc.sms.num.send"
	methodCallTTS    = "alibaba.aliqin.fc.tts.num.singlecall"
	methodCallVoice  = "alibaba.aliqin.fc.voice.num.singlecall"
	methodCallDouble = "alibaba.aliqin.fc.voice.num.doublecall"

	// 正式环境
	httpsURL = "http://gw.api.taobao.com/router/rest"
	// "http://gw.api.taobao.com/router/rest"
	// https://eco.taobao.com/router/rest

	// 沙箱环境
	sandboxHTTPSURL = "https://gw.api.tbsandbox.com/router/rest"

	calledShowNum = "4008221620"
)

var (
	// IsDebug切换到沙箱环境
	IsDebug bool

	Appkey    string
	AppSecret string

	cm = &commonModel{
		// // TOP分配给应用的AppKey
		// AppKey: Appkey,
		//API协议版本，可选值：2.0
		V: "2.0",
		//签名的摘要算法，可选值为：hmac，md5。
		SignMethod: "md5",
		//响应格式。默认为xml格式，可选值：xml，json
		Format: "json",
	}

	sm = &smsModel{
		// 类型.normal：短信
		SmsType: methodSendSMS,
	}
)

// SendOnce 短信单条发送的接口
// moblie-手机号码
// signname-短信签名
// templatecode-短信模板
// param-传入参数
// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错
func SendOnce(moblie, signname, templatecode, param string) (*Result, error) {
	sm.RecNum = moblie
	sm.SmsFreeSignName = signname
	sm.SmsTemplateCode = templatecode
	sm.SmsParam = param
	cm.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	res, err := messasg(cm, sm)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendBatch 短信单条发送的接口
// moblie-手机号码 ','隔开
// signname-短信签名
// templatecode-短信模板
// param-传入参数
// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错
func SendBatch(moblie, signname, templatecode, param string) ([]*Result, error) {
	// 建立电话号码数组
	moblieList := strings.Split(moblie, ",")

	h := int(len(moblieList) / 200)
	// 判断是否有多余值
	if (len(moblieList) % 200) != 0 {
		h++
	}
	// 声明返回值
	resultArrays := make([]*Result, 0, h)
	var cutMoblieArrays []string
	var cutStringStr string
	if h > 0 {
		for i := 1; i <= h; i++ {
			if i == h {
				cutMoblieArrays = moblieList[(i-1)*200:]
			} else {
				cutMoblieArrays = moblieList[(i-1)*200 : (i)*200]
			}
			cutStringStr = strings.Join(cutMoblieArrays, ",")
			cutResult, err := SendOnce(cutStringStr, signname, templatecode, param)
			if err != nil {
				return nil, err
			}
			resultArrays = append(resultArrays, cutResult)
		}
	}
	return resultArrays, nil
}

// SendLecall 文本转语音的接口
// moblie-手机号码
// templatecode-短信模板
// param-传入参数
// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错
func SendLecall(mobile, templatecode, param string) (*Result, error) {
	cm.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	lm := &lecallModel{
		// 被叫号码
		CalledShowNum: calledShowNum,
		TtsParam:      param,
		TtsCode:       templatecode,
		CalledNum:     mobile,
	}

	res, err := lecall(cm, lm)
	if err != nil {
		return nil, err
	}
	return res, nil
}
