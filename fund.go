package funding

// Fund represents an amount stored in a Fund. 
type Fund struct {
    // balance is unexported (private), because it's lowercase
    balance int
}

// NewFund is a regular function returning a pointer to a new fund.
func NewFund(initialBalance int) *Fund {
    // We can return a pointer to a new struct without worrying about
    // whether it's on the stack or heap: Go figures that out for us.
    return &Fund{
        balance: initialBalance,
    }
}

// Balance returns the balance of the fund.
func (f *Fund) Balance() int {
    return f.balance
}

// Withdraw decreases the amount given from the fund.
func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}