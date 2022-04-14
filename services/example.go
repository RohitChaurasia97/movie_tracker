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

	"movie_tracker/repositories"

	"github.com/getsentry/sentry-go"
	"github.com/retail-ai-inc/bean/helpers"
)

type ExampleService interface {
	GetMasterSQLTableList(ctx context.Context) (string, error)
}

type exampleService struct {
	exampleRepository repositories.ExampleRepository
}

func NewExampleService(exampleRepo repositories.ExampleRepository) *exampleService {
	return &exampleService{exampleRepo}
}

func (service *exampleService) GetMasterSQLTableList(ctx context.Context) (string, error) {
	span := sentry.StartSpan(ctx, "http.service")
	span.Description = helpers.CurrFuncName()
	defer span.Finish()
	return service.exampleRepository.GetMasterSQLTableName(span.Context())
}
