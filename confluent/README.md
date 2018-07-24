# Kafka Go Client examples

## Confluent Go
Confluent's Go client for Kafka is dependent on the `librdkafka` library which must be installed first to run the examples locally. See instructions [here](https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka). 

```sh
# Start Kafka and Zookeeper
docker-compose up 

# Start the producer, which publishes messages to a given topic every 10 seconds
go run producer/main.go

# Consume the messages
go run consumer/main.go
```

### Deploying to Kubernetes or Openshift
The `deployment.yaml` file in this repo assumes you're using a [Strimzi](http://strimzi.io/) cluster for Kafka and Zookeeper. 

```sh
# Build and push the consumer image to your Dockerhub
export DOCKER_ORG=<your-dockerhub-username>

cd consumer
make docker_release

# Build and push the producer image
cd ../producer
make docker_release
```

Update `deployment.yaml` with your dockerhub username, and create the `Deployment`s:

```sh
# Deploy to an Openshift cluster
oc create -f confluent/deployment.yaml

# Deploy to a k8s cluster
kubectl create -f confluent/deployment.yaml
```