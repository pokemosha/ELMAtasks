package services

import (
	"ELMAcourses/config"
	"ELMAcourses/model"
	"ELMAcourses/tasks"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

func SolveTask(w http.ResponseWriter, r *http.Request) {
	taskName := chi.URLParam(r, "taskName")
	if taskName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := http.Get(config.AddrPublic + "/tasks/" + taskName)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(resp.StatusCode)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	taskData := [][]interface{}{}

	err = json.Unmarshal(content, &taskData)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("%v", taskData)
	result := []interface{}{}
	switch taskName {
	case "Циклическая ротация":
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			value := int(taskData[i][1].(float64))
			result = append(result, tasks.CycleRotation(array, value))
		}
	case "Чудные вхождения в массив":
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.WonderfulArrayEntries(array))
		}
	case "Проверка последовательности":
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.CheckSubsequence(array))
		}
	case "Поиск отсутствующего элемента":
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.FindingMissingItem(array))
		}
	default:
		w.Write([]byte("Invalid task name"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%v", result)
	var data = model.TaskData{}
	data.Username = config.Username
	data.Task = taskName
	for i := range taskData {
		data.Results = append(data.Result.Results, result[i])
		data.Payloads = append(data.Payloads, taskData[i])
	}
	err = CheckSolution(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetSolution() {

}

func CheckSolution(data model.TaskData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	val, err := http.Post(config.AddrPublic+config.PostReq, "application/json",
		bytes.NewBuffer(body))
	log.Printf("%s", string(body))
	if err != nil {
		return err
	}
	if val.StatusCode != http.StatusOK {
		return errors.New("Response is not ok")
	}

	defer val.Body.Close()

	content, err := ioutil.ReadAll(val.Body)
	if err != nil {
		return err
	}
	var sol = model.Solution{}
	err = json.Unmarshal(content, &sol)
	if err != nil {
		return err
	}
	log.Printf("%v", sol)
	return nil
}

func ConvertToArray(value interface{}) []int {
	array := value.([]interface{})
	result := make([]int, len(array))
	for i := range array {
		result[i] = int(array[i].(float64))
	}
	return result
}

func CreateServer() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/task/{taskName}", SolveTask)
	return router
}
