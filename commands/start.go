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
package commands

import (
	"movie_tracker/middlewares"
	"movie_tracker/routers"
	"movie_tracker/validations"

	"github.com/retail-ai-inc/bean"
	berror "github.com/retail-ai-inc/bean/error"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	host       string
	port       string
	startWeb   bool
	startQueue bool

	// startCmd represents the start command.
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the web service",
		Long:  `Start the web service base on the configs in env.json`,
		Run:   start,
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	defaultHost := bean.BeanConfig.HTTP.Host
	defaultPort := bean.BeanConfig.HTTP.Port
	startCmd.Flags().StringVar(&host, "host", defaultHost, "host address")
	startCmd.Flags().StringVar(&port, "port", defaultPort, "port number")
	startCmd.Flags().BoolVarP(&startWeb, "web", "w", false, "start with web service")
	startCmd.Flags().BoolVarP(&startQueue, "queue", "q", false, "start with job queue worker pool service")
}

func start(cmd *cobra.Command, args []string) {
	b := bean.New()
	// Add custom validation to the default validator.
	b.UseValidation(
		// Example:
		validations.Example,
	)

	// Set custom middleware in here.
	b.UseMiddlewares(
		// Example:
		middlewares.Example("example middleware"),
	)

	// Set custom error handler function here.
	// Bean use a error function chain inside the default http error handler,
	// so that it can easily add or remove the different kind of error handling.
	b.UseErrorHandlerFuncs(
		berror.ValidationErrorHanderFunc,
		berror.APIErrorHanderFunc,
		berror.EchoHTTPErrorHanderFunc,
		// Set your custom error handler func here, for example:
		// func(e error, c echo.Context) (bool, error) {
		// 	return false, nil
		// },
	)

	b.BeforeServe = func() {
		// Init DB dependency.
		b.InitDB()

		// Init different routes.
		routers.Init(b)

		// You can also replace the default error handler:
		// b.Echo.HTTPErrorHandler = YourErrorHandler()

		// Or default validator:
		// b.Echo.Validator = &CustomValidator{}
	}

	if startQueue && startWeb {
		// Start both web and worker pool
		go func() {
			// TODO: start the queue
		}()
		b.ServeAt(host, port)
	} else if startQueue && !startWeb {
		// Only start worker pool
		// TODO: start the queue
	} else {
		// Only start the web
		b.ServeAt(host, port)
	}
}
