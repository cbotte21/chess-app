package schema

type Schema[T any] interface {
	Database() string
	Collection() string
}
