# github.com/rovergulf/kanyerest

[kanye.rest](https://kanye.rest) Go client library

### Install library
```shell
go get github.com/rovergulf/kanyerest
```

### Example usage
```go

import (
	"context"
	"net/http"
	
	"github.com/rovergulf/kanyerest"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	httpClient := &http.Client{Timeout: 5 * time.Second}
	
	yeClient := kanyerest.NewYeClient(httpCLient)
	
	res, err := yeClient.GetQuote()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(res.Quote)
}

```
