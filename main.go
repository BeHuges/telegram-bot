package main

import (
	"flag"
	"log"
)

func main() {

	// token = flags.Get(token)

	token := mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("tocken is not specified")
	}

	return *token

}
