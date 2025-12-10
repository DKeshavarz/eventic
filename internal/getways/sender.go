package getways

type Sender interface {
	Send(to string, msg *Message) error
}

type Message struct {
	Title string
	Text  string
}
