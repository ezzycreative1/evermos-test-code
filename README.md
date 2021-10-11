# evermos-test-code
# Create config
```cp .env-example .env```

# Set Config
```set HTTP_PORT 3001```

And adjust your settings accordingly.

# Run application
`gin -p 4040 main.go`

And you are ready to Go!

# Migration

**to Run migration**

`DBEVENT=migrate go run main.go`

**to rollback**

`DBEVENT=rollback go run main.go`

**to rollback and migration**

`DBEVENT=rollback_migrate go run main.go`

**for dockerfile** 

add `ENV DBEVENT=rollback_migrate`

if you only want to run code, you just have to execute the command `go run main.go`
