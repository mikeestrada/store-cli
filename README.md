# store-cli

## Install Redis
`brew install redis`
`brew services start redis`
Run `redis-cli PING` to confirm cache is running

##Usage
### Add
- Add a k,v pair with `go run main.go ADD key val`
- Get a value with `go run main.go MEMBERS key`