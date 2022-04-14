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
package routers

import (
	"net/http"

	"movie_tracker/handlers"
	"movie_tracker/repositories"
	"movie_tracker/services"

	"github.com/labstack/echo/v4"
	"github.com/retail-ai-inc/bean"
)

type Repositories struct {
	exampleRepo repositories.ExampleRepository
	movieRepo   repositories.MovieRepository // added by bean
}

type Services struct {
	exampleSvc services.ExampleService
	movieSvc   services.MovieService // added by bean
}

type Handlers struct {
	exampleHdlr handlers.ExampleHandler
}

func Init(b *bean.Bean) {

	e := b.Echo

	repos := &Repositories{
		exampleRepo: repositories.NewExampleRepository(b.DBConn),
		movieRepo:   repositories.NewMovieRepository(b.DBConn), // added by bean
	}

	svcs := &Services{
		exampleSvc: services.NewExampleService(repos.exampleRepo),
		movieSvc:   services.NewMovieService(repos.movieRepo), // added by bean
	}

	hdlrs := &Handlers{
		exampleHdlr: handlers.NewExampleHandler(svcs.exampleSvc),
	}

	// Default index page goes to above JSON (/json) index page.
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": `movie_tracker 🚀`,
		})
	})

	// IMPORTANT: Just a JSON response index page. Please change or update it if you want.
	e.GET("/json", hdlrs.exampleHdlr.JSONIndex)

	// IMPORTANT: Just a HTML response index page. Please change or update it if you want.
	e.GET("/html", hdlrs.exampleHdlr.HTMLIndex)

	// Example of using validator.
	e.POST("/example", hdlrs.exampleHdlr.Validate)
}
