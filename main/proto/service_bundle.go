package proto


type ServiceBundle struct {
	Sessions SessionServiceClient
	Users UserServiceClient
}
