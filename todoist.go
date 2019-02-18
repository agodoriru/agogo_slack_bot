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

func getActiveTask() (j_struct, error) {
	request, _ := http.NewRequest("GET", "https://beta.todoist.com/API/v8/tasks", nil)

	auth_val := " Bearer " + getTodoistApiToken()
	request.Header.Set("Authorization", auth_val)

	client1 := &http.Client{}
	response, _ := client1.Do(request)

	return_json, _ := ioutil.ReadAll(response.Body)

	buff := bytes.NewBuffer(return_json)

	json_str := buff.String()
	json_byte := ([]byte)(json_str)

	defer response.Body.Close()

	var jj j_struct

	if err := json2.Unmarshal(json_byte, &jj); err != nil {
		fmt.Println(err)
	}
	return jj, nil
}
