package say_hello

import (
	"encoding/json"
	"fmt"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/serializers"
	"io"
	"net/http"
	"net/http/httptest"
)

func (ctx *HelloFeatureContext) iSendAGETRequestTo(path string) error {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	ctx.responseRecorder = httptest.NewRecorder()
	ctx.router.ServeHTTP(ctx.responseRecorder, req)

	body, err := io.ReadAll(ctx.responseRecorder.Body)
	if err != nil {
		return err
	}

	ctx.response = string(body)
	return nil
}

func (ctx *HelloFeatureContext) iShouldGetStatusCode(expectedCode int) error {
	if ctx.responseRecorder.Code != expectedCode {
		return fmt.Errorf("expected status code %d but got %d", expectedCode, ctx.responseRecorder.Code)
	}
	return nil
}

func (ctx *HelloFeatureContext) theResponseShouldContain(expectedMessage string) error {
	var response serializers.HelloSerializer
	if err := json.Unmarshal([]byte(ctx.response), &response); err != nil {
		return err
	}

	if response.Message != expectedMessage {
		return fmt.Errorf("expected message %q but got %q", expectedMessage, response.Message)
	}

	return nil
}
