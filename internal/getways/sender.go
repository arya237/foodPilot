package getways

type Sender interface{
	Send(to, message string) error
}