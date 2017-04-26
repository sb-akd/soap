package soap

type BadResponse struct {
	RawBody []byte
}

func (e *BadResponse) Error() string {
	return "This is not a SOAP-Message: \n" + string(e.RawBody)
}
