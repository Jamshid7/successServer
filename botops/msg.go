package botops

type Message struct {
	Text string
	Priority string
	PriorityKey int
}

func NewMessage(txt string, p string, pk int) *Message{
	return &Message{txt, p, pk}
}

type Messages struct {
	Msgs []Message
}

func (msgs *Messages) Delete(msg Message) error{
	for key, val := range msgs.Msgs {
		if val.Text == msg.Text && val.Priority == msg.Priority {
			msgs.Msgs = append(msgs.Msgs[:key], msgs.Msgs[key+1:]...)
			return nil
		}
	}
	return nil
}