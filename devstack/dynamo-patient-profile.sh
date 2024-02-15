#!/bin/bash -e
set -eo pipefail

echo "-----------configure dynamodb patient profile----------------------"
# enable below step to delete table locally if needed
#aws dynamodb delete-table --endpoint-url=http://localhost:4566 \
#    --table-name local-subscription-referrals-Referral
#

AWS_PAGER="" aws dynamodb --endpoint-url=http://localhost:4566 create-table \
  --table-name patient-profile \
  --attribute-definitions \
  AttributeName=PatientId,AttributeType=S \
  AttributeName=FirstName,AttributeType=S \
  --key-schema \
  AttributeName=PatientId,KeyType=HASH \
  AttributeName=FirstName,KeyType=String \
  --provisioned-throughput \
  ReadCapacityUnits=10,WriteCapacityUnits=5
