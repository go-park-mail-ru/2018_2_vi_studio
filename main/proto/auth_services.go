package proto

type AuthServices struct {
	Sessions SessionServiceClient
	Users UserServiceClient
}
