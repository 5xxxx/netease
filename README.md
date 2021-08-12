netease是对网易云信服务端API进行封装，netease已经在线上环境使用。
代码中已经将网易的文档全部拷贝下来方便查阅。

[网易文档](https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%8E%A5%E5%8F%A3%E6%A6%82%E8%BF%B0)

## 注意 
由于本人业务需求，所有接口返回200时只会提取部分参数。

使用示例

```
import (
	"fmt"
	"github.com/5xxxx/netease"
)

func main() {
	n := netease.NewNetEaseIM("appkey", "secret")
	token, err := n.CreateAccount(netease.Account{
		Accid: "xx",
		Name:  "xx",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
```


## License

[Apache License 2.0](https://github.com/5xxxx/netease-im/blob/master/LICENSE)
