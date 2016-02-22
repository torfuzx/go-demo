package logicclient

import (
	"flag"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"google.golang.org/grpc"
	"yto.net.cn/kefu/logic/Godeps/_workspace/src/google.golang.org/grpc"
	"yto.net.cn/kefu/router/rpc/logicclient"
)

func newPool(server) *logicclient.Pool {
	return &logicclient.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (*grpc.ClientConn, error) {
			c, err := grpc.Dial(addr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c *grpc.ClientConn, t time.Time) error {
			// Ping
			return err
		},
	}
}

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", "kefu-dev.hotpu.cn:6379", "")
)

func main() {
	flag.Parse()
	pool = newPool(*redisServer)
}

// A request handler gets a connection from the pool and closes the connection
// when the handler is done:

func serveHome(w http.ResponseWriter, r *http.Request) {
	conn := pool.Get()
	defer conn.Close()
}
