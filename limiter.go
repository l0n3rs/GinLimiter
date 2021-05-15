package limiter

import (
	"errors"
)

type limiter struct{
	currentCon int
	bucket chan int
}


func newLimiter(maxConn int) *limiter {
	return &limiter{
		currentCon: maxConn,
		bucket: make(chan int,maxConn),
	}
}

func (l *limiter)GetCon() error{
	if len(l.bucket)>=l.currentCon {
		return errors.New("Full of the limiter")
	}
	l.currentCon-=1
	l.bucket<-1
	return nil
}

func (l *limiter)Release()error{
	if len(l.bucket)==0 {
		return errors.New("Can't release,because the bucket is empty!")
	}
	l.currentCon+=1
	<-l.bucket
	return nil
}


