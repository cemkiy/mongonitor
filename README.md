# mongonitor
mongodb basic monitor tool.

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
<a href="url"><img src="https://raw.githubusercontent.com/cemkiy/mongonitor/main/mongonitor.png" align="left" height="350" width="280" ></a>
