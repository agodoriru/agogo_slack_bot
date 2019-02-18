package todoist

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

type api struct {
	token  string
	client *http.Client
}

func New(apiKey string) *api {
	Api := &api{apiKey, &http.Client{}}
	return Api
}

func (api *api) sendRequest(taskUrl string) (data []byte, err error) {

	base, err := url.Parse(todoistURL)
	if err != nil {
		return nil, err
	}

	path, err := url.Parse(taskUrl)
	if err != nil {
		return nil, err
	}

	url := base.ResolveReference(path).String()

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := api.getAuthHeader()
	request.Header.Set("Authorization", auth)

	response, err := api.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (api *api) getAuthHeader() string {
	return " Bearer " + api.token
}

func (api *api) GetActiveTasks() (tasks []j_struct, err error) {
	response, err := api.sendRequest(taskURL)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(response).String()
	json_byte := ([]byte)(buff)

	if err := json2.Unmarshal(json_byte, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (api *api) GetActiveTaskNames() (taskNames []string, err error) {
	tasks, err := api.GetActiveTasks()
	if err != nil {
		return nil, err
	}

	for index := 0; index < len(tasks); index ++ {
		taskNames = append(taskNames, tasks[index].Content)
	}
	return taskNames, nil
}