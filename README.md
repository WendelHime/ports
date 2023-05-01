# ports

Application for managing in-memory ports

## Testing
You can run tests by executing:
```bash
make test
```

## Building container
```
make build
```

## Running container
```bash
make run # you can press <ctrl-c> whenever you want to finish the app
```

## Routes available

| Endpoint | HTTP method | Description |
| :-- | :-- | :-- |
| `/port` | POST | Sync/upsert port data based on provided input request body |
| `/port/{unloc}` | GET | Retrieve port information |

## Some useful requests

Run the following command for executing a POST request for syncing/upserting ports
```bash
curl -X POST -H "Content-Type: application/json" -d @ports.json http://127.0.0.1:8080/ports
```

Run the following command for retrieving a port after syncing:
```bash
curl -X GET http://127.0.0.1:8080/ports/ANPHI
```

Remember to replace the unloc parameter
