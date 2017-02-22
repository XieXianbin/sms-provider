#smtp

## demo

```go
package main

import (
	"log"

	"github.com/XieXianbin/sms-provider/smtp"
)

func main() {
	s := smtp.New("smtp.exmail.qq.com:25", "notify@a.com", "password")
	log.Println(s.SendMail("notify@a.com", "ulric@b.com;rain@c.com", "这是subject", "这是body,<font color=red>red</font>"))
}
```