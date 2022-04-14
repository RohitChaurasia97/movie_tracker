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
package repositories

import (
	"context"
	"errors"
	models "movie_tracker/models"
	"time"

	// "github.com/retail-ai-inc/bean/helpers"
	// "github.com/getsentry/sentry-go"

	"github.com/retail-ai-inc/bean"
)

type MovieRepository interface {
	GetAllMoviesRepo(c context.Context) ([]models.Movie, error)
	AddMovieRepo(c context.Context, movie models.Movie) ([]models.Movie, error)
	CompleteMovieRepo(c context.Context, idToComplete string) ([]models.Movie, error)
	DeleteMoviesRepo(c context.Context, idToDelete string) ([]models.Movie, error)
	UpdateMoviesRepo(c context.Context, idToUpdate string, movie models.Movie) ([]models.Movie, error)
}

func NewMovieRepository(dbDeps *bean.DBDeps) *DbInfra {
	return &DbInfra{dbDeps}
}

func (db *DbInfra) GetAllMoviesRepo(c context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	err := db.Conn.MasterMySQLDB.Find(&movies).Error
	if err != nil {
		return []models.Movie{}, err
	}
	return movies, nil
}

func (db *DbInfra) AddMovieRepo(c context.Context, movie models.Movie) ([]models.Movie, error) {
	err := db.Conn.MasterMySQLDB.Create(&movie).Error
	var movies []models.Movie
	db.Conn.MasterMySQLDB.Find(&movies)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (db *DbInfra) CompleteMovieRepo(c context.Context, idToComplete string) ([]models.Movie, error) {
	var toComplete models.Movie
	db.Conn.MasterMySQLDB.Find(&toComplete, "ID=?", idToComplete)

	if toComplete.ID == 0 {
		return []models.Movie{}, errors.New("id not found to complete")
	}

	toComplete.Completed = true
	err := db.Conn.MasterMySQLDB.Save(&toComplete).Error

	var movies []models.Movie
	db.Conn.MasterMySQLDB.Find(&movies)

	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (db *DbInfra) DeleteMoviesRepo(c context.Context, idToDelete string) ([]models.Movie, error) {
	var toDoDelete models.Movie
	db.Conn.MasterMySQLDB.Find(&toDoDelete, "ID=?", idToDelete)

	if toDoDelete.ID == 0 {
		return []models.Movie{}, errors.New("id not found to delete")
	}

	err := db.Conn.MasterMySQLDB.Delete(&toDoDelete).Error

	var movies []models.Movie
	db.Conn.MasterMySQLDB.Find(&movies)

	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (db *DbInfra) UpdateMoviesRepo(c context.Context, idToUpdate string, movie models.Movie) ([]models.Movie, error) {
	var movieToFind models.Movie

	db.Conn.MasterMySQLDB.Find(&movieToFind, "ID=?", idToUpdate)

	if movieToFind.ID == 0 {
		return []models.Movie{}, errors.New("id not found to update")
	} else {
		movie.ID = movieToFind.ID
		movie.Completed = movieToFind.Completed
		movie.CreatedAt = movieToFind.CreatedAt
		movie.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := db.Conn.MasterMySQLDB.Save(&movie).Error
		var movies []models.Movie
		db.Conn.MasterMySQLDB.Find(&movies)
		if err != nil {
			return movies, err
		}
		return movies, nil
	}
}
