## msg-provider


=============

Open-Falcon 告警组件 -- 支持短信、邮件、微信告警。

### 查看文档

```
godoc -http=:8080
```

### SMS

使用方法:

短信模版

```
告警内容：
触发值：
IP地址：
所属团队：
告警时间：
```

```
curl http://$ip:8008/sender/sms -d "tos=13523599999,1352358888&context=报警内容"
```

http://127.0.0.1:8008/sender/sms?tos=13523591108&context=sms_content&name=test

### EMAIL

使用方法:

Email模版

```
告警内容：
触发值：
IP地址：
所属团队：
告警时间：
```

```
curl http://$ip:8008/sender/mail -d "tos=me@xiexianbin.cn,10972062@qq.com&subject=demo_subject&content=email_content"
```

http://127.0.0.1:8008/sender/email?tos=me@xiexianbin.cn,10972072@qq.com&subject=demo_subject&content=email_content

### Wechat

_该方法仅支持已认证的微信公众号_

支持微信公众号[模版消息](https://mp.weixin.qq.com/wiki/17/304c1885ea66dbedf7dc170d84999a9d.html)接受告警信息。[详细信息点此](https://mp.weixin.qq.com/debug/cgi-bin/readtmpl?t=tmplmsg/faq_tmpl)

微信告警模版：

```
{{title.DATA}}
告警主机：{{endpoint.DATA}}
告警条件：{{metric.DATA}}
告警标签：{{tags.DATA}}
触发条件：{{reason.DATA}}
告警最大次数：{{max.DATA}}
当前告警次数：{{current.DATA}}
告警备注：{{note.DATA}}
告警时间：{{timestamp.DATA}}
{{remark.DATA}}
```

微信发送消息：

```
https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=GXD-2TpFjGPPb4lZ2b7bZGb9YpsdAOszkPHEnevlvaQNN90s6RCqpvMWM0uSR4U6zlde2OWwjO6K_KdNwt3MfUWD1PnBHwfdDXgAUHJI63K7p-PVCCza5zJF8wHd0_kVPUDiAAAUZI


{
    "touser": "oNvaVuH1Km6fGz3i8V2iHIR9aspk", 
    "template_id": "IJldjH11vTrcaKAwJ2RaGARWuPwA4gLO91b7VZL6n6Q", 
    "url": "http://openfalcon.xiexianbin.cn:5050/template/view/2", 
    "data": {
        "title": {
            "value": "告警平台异常", 
            "color": "#173177"
        }, 
        "endpoint": {
            "value": "xiexianbin_cn", 
            "color": "#173177"
        }, 
        "metric": {
            "value": "net.port.listen", 
            "color": "#bd2636"
        }, 
        "tags": {
            "value": "port=443", 
            "color": "#bd2636"
        }, 
        "reason": {
            "value": "all(#3): 0==0", 
            "color": "#173177"
        }, 
        "max": {
            "value": "6", 
            "color": "#173177"
        }, 
        "current": {
            "value": "5", 
            "color": "#173177"
        }, 
        "note": {
            "value": "网站异常", 
            "color": "#bd2636"
        }, 
        "timestamp": {
            "value": "2017-02-22 22:12:00", 
            "color": "#173177"
        }, 
        "remark": {
            "value": "请相关运维人员尽快恢复该故障，谢谢！", 
            "color": "#173177"
        }
    }
}
```

msg-provider请求方式：

```
curl http://$ip:8008/sender/wechat -d "tos=openid1,openid1&title=title_1&endpoint=endpoint_1&metric=metric_1&tags=tags_1&reason=reason_1&max=max_1&current=current_1&note=note_1&timestamp=timestamp_1&remark=remark_1&url=url_1"
```

http://127.0.0.1:8008/sender/wechat?tos=oNvaVuH1Km6fGz3i8V2iHIR9aspk&title=title_1&endpoint=endpoint_1&metric=metric_1&tags=tags_1&reason=reason_1&max=max_1&current=current_1&note=note_1&timestamp=timestamp_1&remark=remark_1&url=http://openfalcon.xiexianbin.cn:5050/template/view/2




