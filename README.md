# Cointracker

## Setup

```
git clone https://github.com/dawson1096/cointracker-project.git
cd cointracker-project/
go run .
```

## Endpoints

### Add Address

Method: POST
Endpoint: `localhost:8080/api/:address/add`
`:address` must be a valid BTC address

Example

```
curl localhost:8080/api/bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h/add --request "POST"
```

Response:

```
{
    "address": "bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h",
    "success": true
}
```

### Remove Address

Method: DELETE
Endpoint: `localhost:8080/api/:address/remove`
`:address` must be a valid BTC address

Example

```
curl localhost:8080/api/bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h/remove --request "DELETE"
```

Response:

```
{
    "address": "bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h",
    "success": true
}
```

### Sync Address

Method: POST
Endpoint: `localhost:8080/api/:address/sync`
`:address` must be a valid BTC address

Example

```
curl localhost:8080/api/bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h/sync --request "POST"
```

Response:

```
{
    "success": true
}
```

### List Addresses

Method: GET
Endpoint: `localhost:8080/api/listAddresses`

Example

```
curl localhost:8080/api/listAddresses --request "GET"
```

Response:

```
{
    "result" : [
      "bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h"
    ],
    "success": true
}
```

### Get Address

Method: GET
Endpoint: `localhost:8080/api/:address`
`:address` must be a valid BTC address

Example

```
curl localhost:8080/api/bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h --request "GET"
```

Response:

```
{
    "result" : {
        "address": "bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h",
        "balance": 5349007183772,
        "txns": [
            {
                "hash": "71e4d51181150305a3a6d65ed6bac66d3cda1982ccbb50198208bebe4b24dd22",
                "ver": 1,
                "vin_sz": 2,
                "vout_sz": 12,
                "size": 695,
                "weight": 2132,
                "fee": 22134,
                "relayed_by": "0.0.0.0",
                "lock_time": 0,
                "tx_index": 1226662145939393,
                "double_spend": false,
                "time": 1663373562,
                "block_index": 0,
                "block_height": 0,
                "inputs": [
                    {
                        "sequence": 4294967295,
                        "witness": "02473044022029d5e110848b5de4b1cc144586e284512b5a84bc521881452f56049dc2c65cb502205f024ae617ab9ed0e7d1e6de3bbd45b798e0aa4fb8705d027ab860a98c8960fc012102174ee672429ff94304321cdae1fc1e487edf658b34bd1d36da03761658a2bb09",
                        "script": "",
                        "index": 0,
                        "prev_out": {
                            "type": 0,
                            "spent": true,
                            "value": 328600684,
                            "spending_outpoints": [
                                {
                                    "tx_index": 1226662145939393,
                                    "n": 0
                                }
                            ],
                            "n": 0,
                            "tx_index": 4099503009220989,
                            "script": "0014dc6bf86354105de2fcd9868a2b0376d6731cb92f",
                            "addr": "bc1qm34lsc65zpw79lxes69zkqmk6ee3ewf0j77s3h"
                        }
                    }
                ],
                "out": [
                    {
                        "type": 0,
                        "spent": false,
                        "value": 81571,
                        "spending_outpoints": [],
                        "n": 0,
                        "tx_index": 1226662145939393,
                        "script": "a914f43bdcf565129126375f9d24841ed591dc7dfa5f87",
                        "addr": "3PxQWiJDquddQfHtr2qCJtW4tREvYq53hf"
                    }
                ],
                "result": -319295855,
                "balance": 5349007183772
            }
    },
    "success": true
}
```
