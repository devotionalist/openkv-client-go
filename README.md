# openkv-client-go

OpenKV official **Go** client. Talks to Compute Gateway `kv.v1.KVService` over gRPC.

```bash
go get github.com/devotionalist/openkv-client-go@latest
```

```go
import "github.com/devotionalist/openkv-client-go/kv"

cli, err := kv.Dial(ctx, "127.0.0.1:19000")
rev, err := cli.PutTTL(ctx, []byte("k"), []byte("v"), 0)
```

## API (v0.1)

`Dial` · `Get` / `GetOpt` · `Put` / `PutTTL` · `PutCAS` · `Delete` / `DeleteCAS` · `Scan` / `ListPrefix` · `TTL` / `Expire` / `Persist` · `Watch` · `Begin` (txn)

## Proto

`proto/kv/v1/kv.proto` is synced from [openkv](https://github.com/devotionalist/openkv). Regenerate:

```bash
./scripts/gen.sh
```

## License

Apache-2.0
