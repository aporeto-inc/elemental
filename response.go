package elemental

// A Response contains the response from a Request.
type Response struct {
	StatusCode int
	Data       []byte
	Count      int
	Total      int
	Messages   []string
	Redirect   string
	RequestID  string
	Request    *Request
}

// NewResponse returns a new Response
func NewResponse(req *Request) *Response {

	return &Response{
		RequestID: req.RequestID,
		Request:   req,
	}
}

// Encode encodes the given oject into the response.
func (r *Response) Encode(obj interface{}) (err error) {

	r.Data, err = Encode(r.Request.Accept, obj)
	return err
}
