package services

import (
	"context"
	// "github.com/getsentry/sentry-go"
	// "github.com/retail-ai-inc/bean/helpers"
	"movie_tracker/repositories"
)

type MovieService interface {
	MovieServiceExampleFunc(ctx context.Context) (string, error)
}

type movieService struct {
	movieRepository repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) *movieService {
	return &movieService{
		movieRepository: movieRepo,
	}
}

func (service *movieService) MovieServiceExampleFunc(ctx context.Context) (string, error) {
	// IMPORTANT - If you wanna trace the performance of your handler function then uncomment following 3 lines
	// span := sentry.StartSpan(ctx, "http.service")
	// span.Description = helpers.CurrFuncName()
	// defer span.Finish()
	return "MovieService", nil
}
