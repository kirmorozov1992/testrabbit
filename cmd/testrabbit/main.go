package main

import (
	"git.ipc/kirillmorozov/testrabbit/internal/testrabbit"
)

func main() {
	tr := testrabbit.NewTestRabbit("amqp://guest:guest@192.168.1.5:5672", "transactions", 12345)
	conn := tr.NewConn()
	defer conn.Close()

	tr.DeleteTransactionFromQueue()
}
