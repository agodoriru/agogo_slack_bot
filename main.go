package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Due struct {
	Date      string `json:"date"`
	Recurring bool   `json:"recurring"`
	String    string `json:"string"`
}

type j_struct []struct {
	ID           int64         `json:"id"`
	ProjectID    int           `json:"project_id"`
	Content      string        `json:"content"`
	Completed    bool          `json:"completed"`
	LabelIds     []interface{} `json:"label_ids"`
	Order        int           `json:"order"`
	Indent       int           `json:"indent"`
	Priority     int           `json:"priority"`
	CommentCount int           `json:"comment_count"`
	URL          string        `json:"url"`
	Due          struct {
		Recurring bool   `json:"recurring"`
		String    string `json:"string"`
		Date      string `json:"date"`
	} `json:"due,omitempty"`
}

func main() {

	// todoist
	req1, _ := http.NewRequest("GET", "https://beta.todoist.com/API/v8/tasks", nil)

	auth_val := " Bearer " + getTodoistApiToken()
	req1.Header.Set("Authorization", auth_val)

	client1 := &http.Client{}
	res1, _ := client1.Do(req1)

	return_json, _ := ioutil.ReadAll(res1.Body)

	buff := bytes.NewBuffer(return_json)

	json_str := buff.String()
	json_byte := ([]byte)(json_str)

	defer res1.Body.Close()

	// process
	var jj j_struct

	if err := json2.Unmarshal(json_byte, &jj); err != nil {
		fmt.Println(err)
	}

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
