package redis

import (
	"sync"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	r,err:=NewRedis("192.168.203.43:31031")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	totalReq := 24
	wg.Add(totalReq)
	for i:=0; i< totalReq; i++ {
		go func(i int){
			api:="getUserInfo"
			rate := 1
			bucketCap :=20
			now:= time.Now()
			info,err := r.EvalSha(token_bucket.scriptName,api, bucketCap,rate,time.Now().UnixNano() / 1e6)
			if err != nil {
				t.Error(err)
			}else{
				t.Logf("[%s] %t",now.Format(time.StampMilli),info==1)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
}

