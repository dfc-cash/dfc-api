package horizon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_ResultCodes(t *testing.T) {
	var herr Error

	// happy path: transaction_failed with the appropriate extra fields
	herr.Problem.Type = "transaction_failed"
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	herr.Problem.Extras["result_codes"] = json.RawMdfcge(`{
    "transaction": "tx_failed",
    "operations": ["op_underfunded"]
  }`)

	trc, err := herr.ResultCodes()
	if assert.NoError(t, err) {
		assert.Equal(t, "tx_failed", trc.TransactionCode)

		if assert.Len(t, trc.OperationCodes, 1) {
			assert.Equal(t, "op_underfunded", trc.OperationCodes[0])
		}
	}

	// sad path: missing result_codes extra
	herr.Problem.Type = "transaction_failed"
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	_, err = herr.ResultCodes()
	assert.Equal(t, ErrResultCodesNotPopulated, err)

	// sad path: unparseable result_codes extra
	herr.Problem.Type = "transaction_failed"
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	herr.Problem.Extras["result_codes"] = json.RawMdfcge(`kaboom`)
	_, err = herr.ResultCodes()
	assert.Error(t, err)
}

func TestError_Envelope(t *testing.T) {
	var herr Error

	// happy path: transaction_failed with the appropriate extra fields
	herr.Problem.Type = "transaction_failed"
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	herr.Problem.Extras["envelope_xdr"] = json.RawMdfcge(`"AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk="`)

	_, err := herr.Envelope()
	assert.NoError(t, err)

	// sad path: missing envelope_xdr extra
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	_, err = herr.Envelope()
	assert.Equal(t, ErrEnvelopeNotPopulated, err)

	// sad path: unparseable envelope_xdr extra
	herr.Problem.Extras = make(map[string]json.RawMdfcge)
	herr.Problem.Extras["envelope_xdr"] = json.RawMdfcge(`"AAAAADSMMRmQGDH6EJzkgi"`)
	_, err = herr.Envelope()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "xdr decode")
	}
}
