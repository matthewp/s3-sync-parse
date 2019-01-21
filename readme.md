# s3-sync-parse

Parse the output from `aws s3 sync`.

```go
import "github.com/matthewp/s3-sync-parse"

func main() {
  uploads := ParseAll("upload: public/index.html s3://example.com/index.html")

  ...
}
```

## License

BSD-2-Clause