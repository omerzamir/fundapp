package funding

import (
    "sync"
    "testing"
)

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {

    // Skip N = 1
    if b.N < WORKERS {
        return
    }

    // Add as many dollars as we have iterations this run
    server := NewFundServer(b.N)

    // Casually assume b.N divides cleanly
    dollarsPerFounder := b.N / WORKERS

    // WaitGroup structs don't need to be initialized
    // (their "zero value" is ready to use).
    // So, we just declare one and then use it.
    var wg sync.WaitGroup

    // Spawn off the workers
    for i := 0; i < WORKERS; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < dollarsPerFounder; i++ {
                server.Withdraw(1)
            }
        }()
    }

	// Wait for all the workers to finish
	wg.Wait()

	balance := server.Balance()

    if balance != 0 {
        b.Error("Balance wasn't zero:", balance)
    }
}