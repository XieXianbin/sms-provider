# sms-provider
=============

把smtp封装为一个简单http接口，配置到falcon-sender中用来发送报警短信

## 使用方法

```
curl http://$ip:8008/sender/sms -d "tos=13523599999,1352358888&content=报警内容"
```
