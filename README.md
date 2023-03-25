# Go Bots

A package for managing trading bots in the Go programming language. This package provides a simple and easy-to-use interface for interacting with trading bots, managing their settings, and monitoring their performance.

## Features

- Create, list, update, and delete bots
- Manage bot settings such as currency pairs, strategies, and safety orders
- Control bot actions such as pausing, unpausing, and fetching stats
- Support for multiple strategies and order types

## Installation

To install the package, simply run:

```bash
go get -u github.com/EwanValentine/go-3commas
```

## Usage

```go
package main

import (
    "github.com/EwanValentine/go-3commas"
)
```

### Create a Bot
```go
botConfig := &bots.Bot{
	Name:          "MyBot",
	AccountID:     1,
	Pairs:         bots.Pairs{"BTC/USD"},
	Strategy:      bots.Long,
	BaseOrderVolume: "0.01",
	TakeProfit: "1",
	SafetyOrderVolume: "0.01",
	MartingaleVolumeCoefficient: "1",
	MartingaleStepCoefficient: "1",
	MaxSafetyOrders: 3,
	ActiveSafetyOrdersCount: 1,
	SafetyOrderStepPercentage: "2",
	TakeProfitType: bots.Total,
	StrategyList: bots.StrategyList{
		{"strategy": "nonstop"},
	},
	StartOrderType: bots.Limit,
}

newBot, err := botManager.Create(botConfig)
if err != nil {
	log.Fatalf("Failed to create a new bot: %v", err)
}
```

### List Bots
```go
bots, err := botManager.List()
if err != nil {
	log.Fatalf("Failed to list bots: %v", err)
}

for _, bot := range bots {
	fmt.Printf("Bot ID: %d, Name: %s\n", bot.ID, bot.Name)
}
```

### Fetch Bot Stats

```go
botID := 1
stats, err := botManager.Stats(botID)
if err != nil {
	log.Fatalf("Failed to get bot stats: %v", err)
}

fmt.Printf("Stats for bot ID %d: %v\n", botID, stats)
```

### Pause/Unpause Bot

```go
botID := 1

// Pause the bot
err := botManager.Pause(botID)
if err != nil {
	log.Fatalf("Failed to pause bot: %v", err)
}

// Unpause the bot
err = botManager.Unpause(botID)
if err != nil {
	log.Fatalf("Failed to unpause bot: %v", err)
}
```

## Contributing
Contributions are welcome! If you have any suggestions or improvements, please feel free to submit a pull request or open an issue.

## License
This package is released under the MIT License.
