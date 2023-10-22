package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/migration"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/usecase/create_balance"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/usecase/get_balance"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/web"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/web/webserver"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/pkg/kafka"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-wallet-balance-api", "3307", "wallet_balance_db"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration.RunMigrationBalance(db)

	balanceDB := database.NewBalanceDB(db)

	createBalanceUsecase := create_balance.NewCreateBalanceUsecase(balanceDB)
	getBalanceUsecase := get_balance.NewGetBalanceUsecase(balanceDB)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	consumer := kafka.NewConsumer(&configMap, []string{"balances"})

	msgChan := make(chan *ckafka.Message)

	go consumer.Consume(msgChan)

	go func() {
		var eventDTO create_balance.CreateBalanceEventDTO

		for msg := range msgChan {
			err := json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&eventDTO)

			if err != nil {
				panic(err)
			}

			inputAccountFrom := create_balance.CreateBalanceInputDTO{
				AccountID: eventDTO.Payload.AccountIDFrom,
				Balance:   eventDTO.Payload.BalanceAccountIDFrom,
			}

			inputAccountTo := create_balance.CreateBalanceInputDTO{
				AccountID: eventDTO.Payload.AccountIDTo,
				Balance:   eventDTO.Payload.BalanceAccountIDTo,
			}

			createBalanceUsecase.Execute(inputAccountFrom)
			createBalanceUsecase.Execute(inputAccountTo)
		}
	}()

	balanceHandler := web.NewWebBalanceHandler(*getBalanceUsecase)

	webserver := webserver.NewWebServer(":3003")

	webserver.Router.Use(middleware.Logger)

	webserver.Router.Get("/balances/{account_id}", balanceHandler.GetBalance)

	fmt.Println("Server is running")
	webserver.Start()
}
