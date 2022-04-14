// MIT License

// Copyright (c) The RAI Authors

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package services

import (
	"context"
	models "movie_tracker/models"
	repositories "movie_tracker/repositories"
	// "github.com/getsentry/sentry-go"
	// "github.com/retail-ai-inc/bean/helpers"
)

type MovieService interface {
	GetAllMoviesService(ctx context.Context) ([]models.Movie, error)
	AddMovieService(ctx context.Context, movie models.Movie) ([]models.Movie, error)
	CompleteMovieService(ctx context.Context, idToDelete string) ([]models.Movie, error)
	DeleteMoviesService(ctx context.Context, idToDelete string) ([]models.Movie, error)
	UpdateMoviesService(ctx context.Context, idToDelete string, updateMovie models.Movie) ([]models.Movie, error)
}

type movieService struct {
	movieRepository repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) *movieService {
	return &movieService{movieRepo}
}

func (service *movieService) GetAllMoviesService(ctx context.Context) ([]models.Movie, error) {
	return service.movieRepository.GetAllMoviesRepo(ctx)
}

func (service *movieService) AddMovieService(ctx context.Context, movie models.Movie) ([]models.Movie, error) {
	return service.movieRepository.AddMovieRepo(ctx, movie)
}

func (service *movieService) CompleteMovieService(ctx context.Context, idToUpdate string) ([]models.Movie, error) {
	return service.movieRepository.CompleteMovieRepo(ctx, idToUpdate)
}

func (service *movieService) DeleteMoviesService(ctx context.Context, idToDelete string) ([]models.Movie, error) {
	return service.movieRepository.DeleteMoviesRepo(ctx, idToDelete)
}

func (service *movieService) UpdateMoviesService(ctx context.Context, idToDelete string, updateMovie models.Movie) ([]models.Movie, error) {
	return service.movieRepository.UpdateMoviesRepo(ctx, idToDelete, updateMovie)
}
