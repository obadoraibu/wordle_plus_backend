package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/domain"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (s *Service) NewWord(c *gin.Context, r *domain.NewWordRequest) (*domain.NewWordResponse, error) {
	var word string
	var definition string
	for {
		w, err := RapidAPIRandomWordCall(r.Length)
		if err != nil {
			logrus.Error("rapidapi random word call failed")
			w, err = s.repo.GetNewWordFromStorage(r.Length)
			if err != nil {
				return nil, err
			}
		}
		definition, err = DictionaryCheck(w)
		if err != nil {
			if errors.Is(err, domain.ErrWordDoesntExist) {
				continue
			} else {
				return nil, err
			}
		}
		word = w
		break
	}

	fmt.Println(word)
	fmt.Println(definition)
	resp := &domain.NewWordResponse{
		Word:       word,
		Definition: definition,
	}

	return resp, nil
}

func (s *Service) CheckWord(c *gin.Context, r *domain.CheckWordRequest) (string, error) {
	definition, err := DictionaryCheck(r.Word)
	if err != nil {
		return "", err
	} else {
		return definition, nil
	}
}

func (s *Service) DailyWord(c *gin.Context) (*domain.DailyWordResponse, error) {
	word, err := s.repo.GetDailyWordFromStorage()
	if err != nil {
		return nil, err
	}

	definition, err := DictionaryCheck(word)
	if err != nil {
		return nil, err
	}
	resp := &domain.DailyWordResponse{
		Word:       word,
		Definition: definition,
	}

	return resp, nil
}

func DictionaryCheck(word string) (string, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", domain.ErrWordDoesntExist
	}
	var dictionaryResponse domain.DictionaryResponse
	err = json.Unmarshal(body, &dictionaryResponse)
	if err != nil {
		return "", err
	}
	if len(dictionaryResponse) > 0 && len(dictionaryResponse[0].Meanings) > 0 && len(dictionaryResponse[0].Meanings[0].Definitions) > 0 {
		return dictionaryResponse[0].Meanings[0].Definitions[0].Definition, nil
	}
	return "", domain.ErrWordDoesntExist
}

func RapidAPIRandomWordCall(length int) (string, error) {
	url := fmt.Sprintf("https://random-word-api.herokuapp.com/word?length=%d", length)
	res, err := http.Get(url)

	if err != nil {
		return "", domain.ErrRapidAPIRandomWord
	}

	defer res.Body.Close()

	var words []string
	err = json.NewDecoder(res.Body).Decode(&words)

	if err != nil || len(words) == 0 {
		return "", domain.ErrRapidAPIRandomWord
	}

	return words[0], nil
}
