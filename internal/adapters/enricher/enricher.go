package enricher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"sync"
)

type configer interface {
	GetAgeApiURL() string
	GetGenderApiURL() string
	GetNationalityApiURL() string
}

type enricher struct {
	ageUrl         string
	genderUrl      string
	nationalityUrl string
	client         *http.Client
}

func New(cfg configer) *enricher {
	return &enricher{
		ageUrl:         cfg.GetAgeApiURL(),
		genderUrl:      cfg.GetGenderApiURL(),
		nationalityUrl: cfg.GetNationalityApiURL(),
		client:         &http.Client{},
	}
}

func (e enricher) Enrich(ctx context.Context, name string) (*dto.EnrichDataDTO, error) {
	errCh := make(chan error)
	resCh := make(chan *dto.EnrichDataDTO)

	newCtx, cansel := context.WithCancel(ctx)
	defer cansel()

	go func() {
		defer close(errCh)
		defer close(resCh)

		var (
			age         *dto.AgeDTO
			gender      *dto.GenderDTO
			nationality *dto.NationalityDTO
			err         error
		)

		w := &sync.WaitGroup{}
		w.Add(3)

		go func() {
			defer w.Done()
			age, err = e.getAge(newCtx, name)
			fmt.Println(err)
			if err != nil {
				errCh <- err
				return
			}
		}()

		go func() {
			defer w.Done()
			gender, err = e.getGender(newCtx, name)
			if err != nil {
				errCh <- err
				return
			}
		}()

		go func() {
			defer w.Done()
			nationalities, err := e.getNationality(newCtx, name)
			if err != nil {
				errCh <- err
				return
			}
			if len(nationalities.Country) == 0 {
				errCh <- domain_err.ErrNoNationality
				return
			}
			nationality = &dto.NationalityDTO{CountryId: nationalities.Country[0].CountryId, Probability: nationalities.Country[0].Probability}
		}()

		w.Wait()
		if age != nil && gender != nil && nationality != nil {
			resCh <- &dto.EnrichDataDTO{
				Age:         age.Age,
				Gender:      gender.Gender,
				Nationality: nationality.CountryId,
			}
		}
	}()

	select {
	case <-ctx.Done():
		return nil, domain_err.ErrorTimeOut
	case err := <-errCh:
		return nil, err
	case res := <-resCh:
		return res, nil
	}
}

func (e enricher) getAge(ctx context.Context, name string) (*dto.AgeDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", e.ageUrl, name), nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("errror to get age")
	}
	defer resp.Body.Close()
	ageDto := &dto.AgeDTO{}
	err = json.NewDecoder(resp.Body).Decode(&ageDto)
	if err != nil {
		return nil, err
	}

	return ageDto, nil
}

func (e enricher) getGender(ctx context.Context, name string) (*dto.GenderDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", e.genderUrl, name), nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("errror to get gender")
	}
	defer resp.Body.Close()
	genderDto := &dto.GenderDTO{}
	err = json.NewDecoder(resp.Body).Decode(&genderDto)
	if err != nil {
		return nil, err
	}

	return genderDto, nil
}

func (e enricher) getNationality(ctx context.Context, name string) (*dto.NationalitiesDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", e.nationalityUrl, name), nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("errror to get nationality")
	}
	defer resp.Body.Close()
	nationalitiesDto := &dto.NationalitiesDTO{}
	err = json.NewDecoder(resp.Body).Decode(&nationalitiesDto)
	if err != nil {
		return nil, err
	}

	return nationalitiesDto, nil
}
