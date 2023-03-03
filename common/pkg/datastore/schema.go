package datastore

type Schema[T any] interface {
	Database() string
	Collection() string
}
