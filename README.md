# mongonitor
Mongo db basic monitor tool for golang package.

outputs;

- query
- db
- cluster info
- request id
- milliseconds
- filter
- limit
- sort by
- pipeline for aggregate

### mongo package
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

### additional info for mongo go driver
[event#CommandMonitor in mongo-go-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.4.4/event#CommandMonitor)

## usage
```go
// Set client options
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(mongonitor.NewMongonitor())

// use with newrelic
// nrMon := nrmongo.NewCommandMonitor(mongonitor.NewMongonitor())
// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(nrMon)

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
```

### output
<a href="url"><img src="https://raw.githubusercontent.com/cemkiy/mongonitor/main/mongonitor.png"></a>
