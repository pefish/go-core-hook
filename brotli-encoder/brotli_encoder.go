package brotli_encoder

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
	go_core_type_api "github.com/pefish/go-core-type/api"
	api_session "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
)

type BrotliEncoder struct {
}

func NewBrotliEncoder() *BrotliEncoder {
	return &BrotliEncoder{}
}

var BrotliEncoderInstance = NewBrotliEncoder()

func (be *BrotliEncoder) Encode(apiSession api_session.IApiSession, apiResult *go_core_type_api.ApiResult) (interface{}, *go_error.ErrorInfo) {
	var b bytes.Buffer
	bw := brotli.NewWriter(nil)
	bw.Reset(&b)
	if _, err := io.WriteString(bw, apiResult.Data.(string)); err != nil {
		apiSession.Logger().Error(err)
		return nil, go_error.INTERNAL_ERROR
	}
	if err := bw.Close(); err != nil {
		apiSession.Logger().Error(err)
		return nil, go_error.INTERNAL_ERROR
	}

	apiSession.ResponseWriter().Header().Set("Content-Type", "text/html; charset=UTF-8")
	apiSession.ResponseWriter().Header().Set("Vary", "Accept-Encoding")
	apiSession.ResponseWriter().Header().Set("Content-Encoding", "br")
	apiSession.ResponseWriter().WriteHeader(200)
	apiSession.ResponseWriter().Write(b.Bytes())
	return nil, nil
}
