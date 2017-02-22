## msg-provider
=============

Open-Falcon 告警组件 -- 支持短信、邮件、微信告警。

### SMS

使用方法:

```
curl http://$ip:8008/sender/sms -d "tos=13523599999,1352358888&content=报警内容"
```

http://127.0.0.1:8008/sender/sms?tos=13523591108&content=sms_content

### EMAIL

使用方法:

```
curl http://$ip:8008/sender/mail -d "tos=me@xiexianbin.cn,10972062@qq.com&subject=demo_subject&content=email_content"
```

curl http://127.0.0.1:8008/sender/mail -d "tos=me@xiexianbin.cn,10972062@qq.com&subject=demo_subject&content=email_content"

http://127.0.0.1:8008/sender/email?tos=me@xiexianbin.cn&subject=demo_subject&content=email_content

### Wechat



