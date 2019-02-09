package funding

import (
	"fmt"
)

// FundServer is a server of Fund.
type FundServer struct {
	commands chan interface{}
	fund     Fund
}

type withdrawCommand struct {
	Amount int
}

type balanceCommand struct {
	Response chan int
}

// Balance returnes the balance of the server's fund.
func (s *FundServer) Balance() int {
	responseChan := make(chan int)
	s.commands <- balanceCommand{Response: responseChan}
	return <-responseChan
}

// Withdraw is decreasing amount from the balance of the server.
func (s *FundServer) Withdraw(amount int) {
	s.commands <- withdrawCommand{Amount: amount}
}

// NewFundServer returning Fund Server
func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make() creates builtins like channels, maps, and slices
		commands: make(chan interface{}),
		fund:     *NewFund(initialBalance),
	}
	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	for command := range s.commands {

		// command is just an interface{}, but we can check its real type
		switch command.(type) {

		case withdrawCommand:
			// And then use a "type assertion" to convert it
			withdrawal := command.(withdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case balanceCommand:
			getBalance := command.(balanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
