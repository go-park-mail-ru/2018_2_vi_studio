package game

type ClientQueue struct {
	Clients []*Client
}

func (cq *ClientQueue)pop() *Client {
	result := cq.Clients[0]
	cq.Clients = cq.Clients[1:]
	return result
}

func (cq *ClientQueue)push(client *Client) {
	cq.Clients = append(cq.Clients, client)
}

func (cq *ClientQueue)len() int {
	return len(cq.Clients)
}