package main

import (
	"bytes"
	"net/http"
)

func main() {
	jj, _ := getActiveTask()
	content_arr := []string{}
	content_str := ""
	for index := 0; index < len(jj); index++ {

		content_arr = append(content_arr, jj[index].Content)
		content_str += " * " + jj[index].Content + "\n"
	}

	content_header := `\n *remaining tasks* \n`
	content_footer := `\n footer \n`
	content := content_header + content_str + content_footer
	json := `{"text": "` + content + `   \n"}`
	// slack
	req2, _ := http.NewRequest("POST", getSlackWebhookUrl(), bytes.NewBuffer([]byte(json)))
	req2.Header.Set("Content-Type", "application/json")
	client2 := &http.Client{}
	res2, _ := client2.Do(req2)
	defer res2.Body.Close()
}
