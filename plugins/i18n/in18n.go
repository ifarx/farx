package i18n

type Message string

type MessageSource map[string]Message

func (s MessageSource) GetMessage(k string) Message {
	return s[k]
}

func (s MessageSource) PushMessage(k string, m Message) {
	s[k] = m
}

type MessageManager map[string]MessageSource

func (m MessageManager) GetMessage(s, k string) Message {
	if m[k] == nil {
		return ""
	}
	return m[k].GetMessage(k)
}

func (m MessageManager) PushMessage(s, k string, _m Message) {
	if m[k] == nil {
		m[k] = MessageSource{}
	}
	m[k].PushMessage(k, _m)
}

var Instance = MessageManager{}
