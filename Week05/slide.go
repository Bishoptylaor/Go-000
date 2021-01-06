package Week05

import (
	"fmt"
	"sync"
	"time"
)

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2021/1/6
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

var BucketsNum float64 = 10
var BucketsTime int64 = 10

type BNumber struct {
	Buckets map[int64]*numberBucket
	*sync.RWMutex
}

type numberBucket struct {
	Value float64
}

// 初始化
func NewBNumber() *BNumber {
	r := &BNumber{
		Buckets: make(map[int64]*numberBucket),
	}
	return r
}

func (r *BNumber) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	return bucket
}

func (r *BNumber) removeOldBuckets() {
	now := time.Now().Unix()

	for timestamp := range r.Buckets {
		if timestamp <= now - BucketsTime {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *BNumber) Increment(i float64) {
	if i == 0 {
		return
	}
	r.Lock()
	defer r.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	fmt.Println(i)
	r.removeOldBuckets()

}

func (r *BNumber) UpdateMax(n float64) {
	r.Lock()
	defer r.Unlock()

	b := r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.removeOldBuckets()
}

func (r *BNumber) Sum(now time.Time) float64 {
	sum := float64(0)

	r.RLock()
	defer r.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()- BucketsTime {
			sum += bucket.Value
		}
	}
	fmt.Println("sum:", sum)
	return sum
}

func (r *BNumber) Max(now time.Time) float64 {
	var max float64

	r.RLock()
	defer r.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()- BucketsTime {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	fmt.Println("max:", max)
	return max
}

func (r *BNumber) Avg(now time.Time) float64 {
	avg := r.Sum(now) / BucketsNum
	fmt.Println("avg is:", avg)
	return avg
}
