# mongonitor
mongodb basic monitor tool for golang package.

outputs;

- query
- db
- cluster
- request id
- miiliseconds
- filter
- limit
- sort by
- pipeline for aggregate

# mongo package
https://github.com/mongodb/mongo-go-driver

# usage
```go
// use with newrelic
nrMon := nrmongo.NewCommandMonitor(mongonitor.NewMongonitor())
clientOptions := options.Client().ApplyURI(uri).SetMonitor(nrMon)

// standart usage
clientOptions := options.Client().ApplyURI(uri).SetMonitor(mongonitor.NewMongonitor())
```

# output
<a href="url"><img src="https://raw.githubusercontent.com/cemkiy/mongonitor/main/mongonitor.png"></a>
