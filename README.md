# env

Package env implements a environment variable retrieval utility.

## Install

```
go get github.com/gomodrepo/env
```

## Usage

```go
import "github.com/gomodrepo/env"

value := env.Get("ENV_KEY", "defaultValue")
```

