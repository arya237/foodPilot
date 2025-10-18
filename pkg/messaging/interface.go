package messaging

type Sender interface{
	Send(to, message string) error
}