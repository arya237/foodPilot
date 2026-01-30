package messaging

type Config struct {
	From string
	Key  string
}

type Sender interface{
	Send(to, message string) error
}