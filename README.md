### Go Choreo sample

#### Go build & run

```shell
go build main.go && go run main.go
```

#### Generate API definitions

Generated using Go annotations https://github.com/swaggo/swag

```shell
# .git/hooks/pre-commit
swag init
```

#### Load initial data ( optional )

1. Set env in Choreo DevOps portal`INIT_DATA_PATH=<some_path>`
2. Mount the file contents of `configs/initial_data.json` in the path specified in step 1.