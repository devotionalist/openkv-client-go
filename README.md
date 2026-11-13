# openkv-client-go

OpenKV official **Go** client. Talks to Gateway `kv.v1.KVService` / `DSService` over gRPC.

```bash
go get github.com/devotionalist/openkv-client-go@latest
```

```go
import "github.com/devotionalist/openkv-client-go/kv"

cli, err := kv.Dial(ctx, "127.0.0.1:19000")
rev, err := cli.PutTTL(ctx, []byte("k"), []byte("v"), 0)
```

## API planes

| plane | 含义 | 方法 |
|-------|------|------|
| **txn** | MVCC 事务面（TSO / lock / ONEPC） | `Get` `Put` `Delete` `BatchGet` `BatchPut` `Scan` `TTL` `Watch` `Begin` … |
| **raw** | 无事务 | **仅** `RawGet` `RawPut` `RawPutTTL` `RawPutCAS` `RawDelete` `RawDeleteCAS` |
| **ds** | 数据结构（底层仍走 MVCC） | 包 `ds`：`HSet`/`SAdd`/`LPush`/`ZAdd` … |

两平面**勿混用同一 key**。`BatchPut` / `BatchGet` 属事务面：Gateway **扇出**单 key，不是多 key 一次 Raft Propose；尚无 `RawBatch*`。

## API (v0.1)

**KV** (`kv`): `Dial` · `Get` / `GetOpt` · `Put` / `PutTTL` · `PutCAS` · `Delete` / `DeleteCAS` · `RawGet` / `RawPut` / `RawPutTTL` / `RawPutCAS` / `RawDelete` / `RawDeleteCAS` · `BatchGet` / `BatchGetOpt` · `BatchPut` · `Scan` / `ListPrefix` · `TTL` / `Expire` / `Persist` · `Watch` · `Begin` (txn)

**DS** (`ds`): Hash / Set / List / ZSet (`HSet`/`HGet`/… · `SAdd`/… · `LPush`/… · `ZAdd`/…)

## Proto

`proto/kv/v1/{kv,ds}.proto` synced from [openkv](https://github.com/devotionalist/openkv):

```bash
./scripts/sync-proto-from-openkv.sh
./scripts/gen.sh
```

## License

Apache-2.0
