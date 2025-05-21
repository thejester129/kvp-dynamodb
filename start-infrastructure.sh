#!/bin/bash

cd "${0%/*}" # ensure cwd is script dir

echo -e "\n starting mock infrastructure...\n"

docker compose -f mock-infrastructure.yaml up -d --build 

sleep 10 # wait for localstack

region=us-east-1
table_name=kvp-table
endpoint_url=http://localhost:4566

aws dynamodb create-table \
--endpoint-url $endpoint_url \
--table-name $table_name \
--region $region \
--key-schema AttributeName=key,KeyType=HASH \
--attribute-definitions AttributeName=key,AttributeType=S \
--provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=5 \
--no-cli-pager


# --some test commands--

# aws dynamodb scan --endpoint-url http://localhost:4566 --region us-east-1 --table-name kvp-table

# aws dynamodb put-item --region us-east-1 --endpoint-url http://localhost:4566 --table-name linn-api-user-persistence-table --item '{ "accountId": {"S": "123"}, "domain": {"S": "manage-systems"}, "seen_notifications" : { "NS" : ["1"] } }'

# aws dynamodb get-item --region us-east-1 --endpoint-url http://localhost:4566  --table-name linn-api-user-persistence-table --key '{"accountId": {"S": "123"}, "domain":{"S": "manage-systems"} }'