package advanced

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}
	return orders
}

func updateOrderStatus(order *Order) {

	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf("Updated order %d status: %s\n", order.ID, status)

}

func processOrders(orders []*Order) {
	for _, order := range orders {
		fmt.Printf("Processing order %d\n", order.ID)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		order.Status = "Processed"
	}
}

func reportOrderStatus(orders []*Order) {

	time.Sleep(1 * time.Second)
	fmt.Println("\n --- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}
	fmt.Println("-----------------------------")

}

func go_routines() {
	var wg sync.WaitGroup
	wg.Add(2)

	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
	}()

	reportOrderStatus(orders)

	wg.Wait()
	fmt.Println("All operations completed. Exiting.")
}
