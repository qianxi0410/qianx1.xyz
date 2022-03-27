package r

// we don's need status code
// cause we follow the http status code
// 200 represents success
// 400 represents bad request
// 500 represents internal server error
// etc...
type R[T any] struct {
	Data T      `json:"data"`
	Err  string `json:"err"`
}

func Ok[T any](data T) R[T] {
	return R[T]{
		Data: data,
		Err:  "",
	}
}

func Error[T any](s string) R[T] {
	var t T
	return R[T]{
		Data: t,
		Err:  s,
	}
}
