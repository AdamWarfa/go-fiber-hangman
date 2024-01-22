package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetWord(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	var wordArray []string
	err = json.NewDecoder(response.Body).Decode(&wordArray)

	return wordArray[0]

}
