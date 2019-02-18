package main

import (
	"bytes"
	json2 "encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Due struct {
	Date      string `json:"date"`
	Recurring bool   `json:"recurring"`
	String    string `json:"string"`
}

type j_struct struct {
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

const todoistURL = "https://beta.todoist.com/"
const taskURL = "API/v8/tasks"

func getAuthHeader() (auth string) {
	return " Bearer " + getTodoistApiToken()
}

func getActiveTasks() (tasks []j_struct, err error) {

	base, err := url.Parse(todoistURL)
	if err != nil {
		return nil, err
	}

	api, err := url.Parse(taskURL)
	if err != nil {
		return nil, err
	}

	url := base.ResolveReference(api).String()

	request, _ := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := getAuthHeader()
	request.Header.Set("Authorization", auth)

	client1 := &http.Client{}
	response, err := client1.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return_json, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(return_json).String()
	json_byte := ([]byte)(buff)

	if err := json2.Unmarshal(json_byte, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func getContent() ([]string, error) {
	j, err := getActiveTasks()
	if err != nil {
		return nil, err
	}

	tasks := []string{}
	for index := 0; index < len(j); index++ {
		tasks = append(tasks, j[index].Content)
	}
	return tasks, nil
}

func convArrToStr(arr []string) string {
	tasks := ""
	for index := 0; index < len(arr); index++ {
		tasks += "・" + arr[index] + "\n"
	}
	return tasks
}
