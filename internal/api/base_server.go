package api

type BaseServer interface {
	ListenAndServe() error
	Shutdown() error
}
