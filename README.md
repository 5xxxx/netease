```
import (
	"fmt"
	"github.com/NSObjects/netease"
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

[Apache License 2.0](https://github.com/NSObjects/netease-im/blob/master/LICENSE)
