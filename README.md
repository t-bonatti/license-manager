# license-manager
[![CircleCI](https://circleci.com/gh/t-bonatti/license-manager.svg?style=shield)](https://circleci.com/gh/t-bonatti/license-manager)

Manage application licenses with versions

# Dependences

license-manager expects a PostgreSQL instance

# Environment Variables

- DATABASE_URL
- PORT

# Build

```
go build main.go
```
# Examples

Add license

```
curl -i -H "Content-Type:application/json" -X POST localhost:3000/license -d '{"id": "abcde12345", "version" : "v2", "info": {"user" : "blah", "company" : "xxyyxx"}}'
```

Get license

```
curl -i localhost:3000/license/abcde12345/versions/v2
```
