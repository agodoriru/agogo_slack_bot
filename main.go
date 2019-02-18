package main

import (
	"./todoist"
	"bytes"
	"log"
	"net/http"
)

func main() {

	a := todoist.New(getTodoistApiToken())
	tasks, err := a.GetActiveTaskNames()
	if err != nil {
		log.Fatalln(err)
	}
	content_str := convArrToStr(tasks)

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

func convArrToStr(arr []string) string {
	tasks := ""
	for index := 0; index < len(arr); index++ {
		tasks += "・" + arr[index] + "\n"
	}
	return tasks
}
