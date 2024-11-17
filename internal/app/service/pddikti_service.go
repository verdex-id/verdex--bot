package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"github.com/verdex-id/verdex-bot/internal/app/model"
)

type PDDIKTIService struct {
}

func NewPDDIKTIService() *PDDIKTIService {
	return &PDDIKTIService{}
}

func (s *PDDIKTIService) GetManyUniversity(page int) (*model.PDDIKTIUniversityResponse, error) {
	uri, _ := url.Parse("https://pddikti.kemdikbud.go.id/api/v2/pt/search/filter")

	params := uri.Query()

	// Basic URL
	// https://pddikti.kemdikbud.go.id/api/v2/pt/search/filter?limit=15&page=1&akreditasi=&jenis=&provinsi=&status=&

	params.Set("limit", "10")
	params.Set("page", fmt.Sprint(page))
	params.Set("status", "A") // A = Active

	uri.RawQuery = params.Encode()

	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, uri.String(), nil)
	req.Header.Set("x-api-key", viper.GetString("pddikti.api_key"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var pddiktiUniversityResponse *model.PDDIKTIUniversityResponse
	json.Unmarshal(body, &pddiktiUniversityResponse)

	return pddiktiUniversityResponse, nil
}
