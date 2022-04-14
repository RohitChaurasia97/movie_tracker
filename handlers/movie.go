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
package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	models "movie_tracker/models"
	services "movie_tracker/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type MovieHandler interface {
	GetAllMovies(c echo.Context) error
	AddMovie(c echo.Context) error
	CompleteMovie(c echo.Context) error
	DeleteMovie(c echo.Context) error
	UpdateMovie(c echo.Context) error
}

type movieHandler struct {
	movieService services.MovieService
}

func NewMovieHandler(movieSvc services.MovieService) *movieHandler {
	return &movieHandler{movieSvc}
}

func (handler *movieHandler) GetAllMovies(c echo.Context) error {
	ctx := context.Background()
	output, err := handler.movieService.GetAllMoviesService(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "templates/movie", map[string]interface{}{
		"movies": output,
	})
}

func (handler *movieHandler) AddMovie(c echo.Context) error {
	ctx := context.Background()
	var movie models.Movie

	err := c.Bind(&movie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	output, err := handler.movieService.AddMovieService(ctx, movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "templates/movie", map[string]interface{}{
		"movies": output,
	})
}

func (handler *movieHandler) CompleteMovie(c echo.Context) error {
	ctx := context.Background()
	idToComplete := c.Param("idToComplete")

	output, err := handler.movieService.CompleteMovieService(ctx, idToComplete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "templates/movie", map[string]interface{}{
		"movies": output,
	})
}

func (handler *movieHandler) DeleteMovie(c echo.Context) error {
	ctx := context.Background()

	idToDelete := c.Param("idToDelete")

	output, err := handler.movieService.DeleteMoviesService(ctx, idToDelete)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "templates/movie", map[string]interface{}{
		"movies": output,
	})
}

func (handler *movieHandler) UpdateMovie(c echo.Context) error {
	var updateMovie models.Movie
	ctx := context.Background()
	idToUpdate := c.Param("idToUpdate")

	type jsonparams struct {
		Title string `json:"title"`
	}
	bytesRead, _ := ioutil.ReadAll(c.Request().Body)

	var param jsonparams
	jsonUnmarshalError := json.Unmarshal(bytesRead, &param)
	if jsonUnmarshalError != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": jsonUnmarshalError.Error(),
		})
	}

	updateMovie.Title = param.Title
	updateMovie.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	output, err := handler.movieService.UpdateMoviesService(ctx, idToUpdate, updateMovie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "templates/movie", map[string]interface{}{
		"movies": output,
	})
}
