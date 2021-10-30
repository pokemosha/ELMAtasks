package services

import (
	"ELMAcourses/config"
	"ELMAcourses/model"
	"ELMAcourses/tasks"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

const FindingMissingItem = "Поиск отсутствующего элемента"
const CycleRotation = "Циклическая ротация"
const CheckSubsequence = "Проверка последовательности"
const WonderfulArrayEntries = "Чудные вхождения в массив"

func SolveTask(w http.ResponseWriter, r *http.Request) {
	taskName := chi.URLParam(r, "taskName")
	if taskName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	solution, err := GetSolution(taskName)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ans, err := json.Marshal(solution)
	w.Write(ans)
	w.WriteHeader(http.StatusOK)
}

func SolveAllTasks(w http.ResponseWriter, r *http.Request) {
	var solArray []model.Solution

	solutionCS, err := GetSolution(CheckSubsequence)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	solArray = append(solArray, solutionCS)

	solutionCR, err := GetSolution(CycleRotation)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	solArray = append(solArray, solutionCR)

	solutionFMI, err := GetSolution(FindingMissingItem)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	solArray = append(solArray, solutionFMI)

	solutionWAE, err := GetSolution(WonderfulArrayEntries)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	solArray = append(solArray, solutionWAE)
	ans, err := json.Marshal(solArray)
	w.Write(ans)

}

func GetSolution(taskName string) (model.Solution, error) {
	resp, err := http.Get(config.AddrPublic + "/tasks/" + taskName)
	if err != nil {
		return model.Solution{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.Solution{}, fmt.Errorf("%s", resp.StatusCode)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Solution{}, err
	}
	taskData := [][]interface{}{}

	err = json.Unmarshal(content, &taskData)
	if err != nil {
		return model.Solution{}, err
	}
	log.Printf("%v", taskData)
	result := []interface{}{}
	switch taskName {
	case CycleRotation:
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			value := int(taskData[i][1].(float64))
			result = append(result, tasks.CycleRotation(array, value))
		}
	case WonderfulArrayEntries:
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.WonderfulArrayEntries(array))
		}
	case CheckSubsequence:
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.CheckSubsequence(array))
		}
	case FindingMissingItem:
		for i := range taskData {
			array := ConvertToArray(taskData[i][0])
			result = append(result, tasks.FindingMissingItem(array))
		}
	default:
		return model.Solution{}, err
	}
	log.Printf("%v", result)
	var data = model.TaskData{}
	data.Username = config.Username
	data.Task = taskName
	for i := range taskData {
		data.Results = append(data.Result.Results, result[i])
		data.Payloads = append(data.Payloads, taskData[i])
	}
	sol, err := CheckSolution(data)
	if err != nil {
		return model.Solution{}, err
	}
	return sol, nil
}

func CheckSolution(data model.TaskData) (model.Solution, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return model.Solution{}, err
	}

	val, err := http.Post(config.AddrPublic+config.PostReq, "application/json",
		bytes.NewBuffer(body))
	log.Printf("%s", string(body))
	if err != nil {
		return model.Solution{}, err
	}
	if val.StatusCode != http.StatusOK {
		return model.Solution{}, errors.New("Response is not ok")
	}

	defer val.Body.Close()

	content, err := ioutil.ReadAll(val.Body)
	if err != nil {
		return model.Solution{}, err
	}
	var sol = model.Solution{}
	err = json.Unmarshal(content, &sol)
	if err != nil {
		return model.Solution{}, err
	}
	log.Printf("%v", sol)
	return sol, nil
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
	router.Get("/tasks", SolveAllTasks)
	return router
}
