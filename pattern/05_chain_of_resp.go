package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type IHandler interface {
	SendRequest(m int) string
}

type HandlerOne struct {
	next IHandler
}

func (h *HandlerOne) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerTwo struct {
	next IHandler
}

func (h *HandlerTwo) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type HandlerThree struct {
	next IHandler
}

func (h *HandlerThree) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

/*
	Создание и вызов выполняются так:
	
	handlers := &HandlerOne {
		next: &HandlerTwo{
			next: &HandlerThree{},
		},
	}
	result := handlers.SendRequest(2)
*/
