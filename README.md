# eth-parser
An Ethereum transaction parser built in Go to monitor blockchain activity and notify users of incoming and outgoing transactions for subscribed addresses.


## Features

* Monitors the Ethereum blockchain in real-time using the JSON-RPC interface.

* Allows users to subscribe to specific addresses for transaction updates.

* Exposes HTTP APIs to interact with the parser:

* Subscribe to an address.

* Retrieve transactions for a subscribed address.

* Get the last processed block.

* Designed with in-memory storage (easily extendable to persistent storage).

* Lightweight, efficient, and easily extensible.



## Prerequisites

* Go: Ensure you have Go installed (>= v1.20).

* Ethereum Node: Uses https://ethereum-rpc.publicnode.com for JSON-RPC interactions.


## Setup
1. Clone the Repository
    
        git clone https://github.com/your-username/eth-tx-parser.git
        cd eth-tx-parser
2. Install Dependencies

        go mod tidy

3. Set Environment Variables

use env.local.sh for reference and run the command to load the values.

        source env.local.sh

## HTTP API Endpoints

1. Subscribe to an Address

Subscribe to an Ethereum address for transaction monitoring.

* URL: api/v1/subscribe
* Method: POST
* query: address=0xYourEthereumAddress

Response 

    200OK

2. Get Transactions

Retrieve a list of inbound and outbound transactions for a subscribed address.

* URL: api/v1/transactions
* Method: GET
* query: address=0xYourEthereumAddress

Response 

    200OK

```
{   
    "data": [
        {
            "from": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5",
            "to": "0x2194331af2cf9de9adb36cf09654fad65cafb58b",
            "value": "0xa42403b9763e19",
            "hash": "0xb75368ccb88fee22131a7cdc5ead0391d3618bcd07ef6b5fc6ec82dbf10fed5f",
            "blockNumber": 21207340
        }
    ]
}
```

3. Get Last Processed Block

Retrieve the last block number processed by the parser.

* URL: api/v1/current-block
* Method: GET

Response 

    200OK

```
{
    "data": {
        "blockNumber": 21208999
    }
}
```


## Logging

The application uses logrus for structured logging. You can set the log level through 
the LOG_LEVEL environment variable.


## Testing

Use Postman or curl to test the HTTP APIs.


## Running the server

source the env file and the run the command.

    source env.local.sh
    go run main.go


## License

This project is licensed under the MIT License.


#### Thank you!