### Patient Profile

`/patientprofile` is an endpoint to post patient profile which goes to kafka. For the ease of running and testing locally, using dockerized kafka and localstack for DynamoDB

### Problem Statement

How to guarantee the process of updating patient profile

### Solution

There could be other solutions, one of the approach taken in this project is `Kafka Message Keys`. The reason for this approach is kafka doesn't guarantee order across the partition,
although it guarantees order within the partition. Here, the key for the kafka message to be produced is chosen as `PatientID` so that it goes to same partition

Partial Updates are allowed

## Steps to execute

## Run Locally

Boot up kafka UI and DB locally

```shell
make local-dev-stack
```

Open 

```
http://localhost:8080/
```

Looks like 

![img.png](img.png)

### Steps to execute the app

```shell
make mod-tidy
```


```shell
make mod-dl
```
### Tests
The below make command will run all the unit tests

```shell
make test
```

One one terminal run below

``` shell
make run-producer

```

One other terminal run below

``` shell
make run-consumer

```

Run the following command

```shell
curl --location --request POST '0.0.0.0:3000/patientprofile' \
--header 'Content-Type: application/json' \
--data-raw '{ "PatientId": "1", "FirstName": "John", "LastName": "Doe", "IsPregnant": "False", "Sex": "Female"}' 
```

To check if its inserted in DB properly, run the following command

```shell
aws dynamodb scan --endpoint-url=http://localhost:4566 \      
   --table-name patient-profile
```

The result will be as follows

```shell

{
    "Items": [
        {
            "LastName": {
                "S": "Doe"
            },
            "IsPregnant": {
                "S": "False"
            },
            "UpdatedAt": {
                "S": "2023-08-14T01:11:45.105211+10:00"
            },
            "FirstName": {
                "S": "John"
            },
            "Sex": {
                "S": "Female"
            },
            "PatientId": {
                "S": "1"
            }
        }
    ],
    "Count": 1,
    "ScannedCount": 1,
    "ConsumedCapacity": null
}

```

Kafka will look like this 

![img_1.png](img_1.png)

``
brew install mockery
``