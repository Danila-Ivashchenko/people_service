package enricher

import (
	"context"
	"people_service/internal/domain/dto"
)

type Enricher interface {
	Enriche(context.Context, string) (dto.EnrichDataDTO, error)
}
