package sending_queue

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

type Queue_ struct {
	Url     string
	Methods string
	Value   []byte
}

func init() {
	go sending()
}

var Channel = make(chan Queue_, 10)

func sending() {

	for {
		var request = <-Channel

		var buf bytes.Buffer
		buf.Write(request.Value)
		str := string(request.Value)

		fmt.Println(str)

		client := &http.Client{}
		req, err := http.NewRequest(
			request.Methods, request.Url, &buf,
		)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "MSIE/15.0")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		runtime.Gosched()
	}
}
