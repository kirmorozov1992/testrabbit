package main

import (
	"git.ipc/kirillmorozov/testrabbit/internal/testrabbit"
)

func main() {
	tr := testrabbit.NewTestRabbit("amqp://login:password@localhost:15672/", "test", "tx")
	conn := tr.NewConn()
	defer conn.Close()

	tr.DeleteTransactionFromQueue()
}
