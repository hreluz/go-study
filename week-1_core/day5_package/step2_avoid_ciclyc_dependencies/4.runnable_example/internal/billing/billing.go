package billing

import "fmt"

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) Charge(userID string, amount float64) {
	fmt.Printf("billing: charged %.2f to user %s\n", amount, userID)
}
