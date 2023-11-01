# Holocron commands

```sh

holocron proxy --natsURL="nats://localhost:4222" --port="9999"
./holocron insert --addr localhost:9999 --collection articles -f ./samples/sample-0.yaml
./holocron query --addr localhost:9999 --collection articles -f ./samples/sample-query.yaml
holocron lookup run --bucket="holocron-wodgy"
holocron ingestor run --natsURL="nats://localhost:4222" --bucket my-bucket
./holocron ingestor run --bucket="holocron-wodgy"

```