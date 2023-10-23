package service

import (
	"context"
	"people_service/internal/domain/dto"
	"people_service/pkg/logger"
	mock_enricher "people_service/pkg/mocks/enricher"
	mock_storage "people_service/pkg/mocks/storage"
	"people_service/pkg/validator"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

type dependeses struct {
	storage  *mock_storage.MockPersonStorage
	enricher *mock_enricher.MockEnricher
}

func TestService_AddPerson(t *testing.T) {
	tests := []struct {
		name        string
		deadline    time.Duration
		inp         *dto.AddPersonRawDTO
		enrichData  *dto.EnrichDataDTO
		preparation func(*dependeses, *dto.AddPersonRawDTO, *dto.EnrichDataDTO, context.Context)
		out         int64
		err         error
	}{
		{
			name:     "valid",
			deadline: time.Second * 10,
			inp: &dto.AddPersonRawDTO{
				Name:    "Danila",
				Surname: "Ivashenko",
			},
			enrichData: &dto.EnrichDataDTO{
				Age:         42,
				Gender:      "male",
				Nationality: "RU",
			},
			preparation: func(d *dependeses, data *dto.AddPersonRawDTO, enrichData *dto.EnrichDataDTO, ctx context.Context) {
				personAdd := &dto.AddPersonDTO{
					Name: data.Name,
					Surname: data.Surname,
					Patronymic: data.Patronymic,
					Age: enrichData.Age,
					Gender: enrichData.Gender,
					Nationality: enrichData.Nationality,
				}
				d.storage.EXPECT().AddPerson(ctx, personAdd).Return(int64(1), nil)
				d.enricher.EXPECT().Enriche(ctx, data.Name).Return(enrichData, nil)
			},
			out: 1,
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			validator := validator.NewValidator()
			dependeses := &dependeses{
				storage:  mock_storage.NewMockPersonStorage(ctl),
				enricher: mock_enricher.NewMockEnricher(ctl),
			}
			logger := logger.SetupLogger("test")
			ctx, cansel := context.WithTimeout(context.Background(), tt.deadline)
			defer cansel()
			if tt.preparation != nil {
				tt.preparation(dependeses, tt.inp, tt.enrichData, ctx)
			}
			service := New(dependeses.storage, dependeses.enricher, logger, validator)

			result, err := service.AddPerson(ctx, tt.inp)
			if result != tt.out {
				t.Errorf("got %d, want %d", result, tt.out)
			}
			if err != tt.err {
				t.Errorf("got %v, want %v", err, tt.err)
			}
		})
	}
}
