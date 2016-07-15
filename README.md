# fireside

### DEVELOPMENT

**Regenerating gRPC/Protobufs**

`fireside-backend` uses [protocol buffers](https://github.com/golang/protobuf) with [gRPC](http://www.grpc.io/).

If you change anything in `/proto`, regenerate the `.pb.go` files with the `./rebuildProto.sh` script.

#### Sample `.env`

You need to create a `.env` in the project's root directory containing your local settings

```bash
REST_PORT=:9000
PORT=:9001
POST_ENDPOINT=:9001

DB_HOST=localhost
DB_PORT=5432
DB_USER=nhan
DB_PASSWORD=
DB_DBNAME=fireside
```