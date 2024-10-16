package brotli_encoder

import (
	"bytes"
	"io"

	"github.com/andybalholm/brotli"
	i_core "github.com/pefish/go-interface/i-core"
	t_core "github.com/pefish/go-interface/t-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type BrotliEncoder struct {
}

func NewBrotliEncoder() *BrotliEncoder {
	return &BrotliEncoder{}
}

var BrotliEncoderInstance = NewBrotliEncoder()

func (be *BrotliEncoder) Encode(apiSession i_core.IApiSession, apiResult *t_core.ApiResult) (interface{}, *t_error.ErrorInfo) {
	toEncodeData := ""
	if apiResult.Code != 0 {
		toEncodeData = apiResult.Msg
	} else {
		toEncodeData = apiResult.Data.(string)
	}

	var b bytes.Buffer
	bw := brotli.NewWriter(nil)
	bw.Reset(&b)
	if _, err := io.WriteString(bw, toEncodeData); err != nil {
		apiSession.Logger().Error(err)
		return nil, t_error.INTERNAL_ERROR
	}
	if err := bw.Close(); err != nil {
		apiSession.Logger().Error(err)
		return nil, t_error.INTERNAL_ERROR
	}

	apiSession.ResponseWriter().Header().Set("Content-Type", "text/html; charset=UTF-8")
	apiSession.ResponseWriter().Header().Set("Vary", "Accept-Encoding")
	apiSession.ResponseWriter().Header().Set("Content-Encoding", "br")
	apiSession.ResponseWriter().WriteHeader(200)
	apiSession.ResponseWriter().Write(b.Bytes())
	return nil, nil
}
