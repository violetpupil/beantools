package main

import (
	"github.com/beanstalkd/go-beanstalk"
)

func NewConn() *beanstalk.Conn {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	return c
}

func NewTube(name string) *beanstalk.Tube {
	conn := NewConn()
	t := beanstalk.NewTube(conn, name)
	return t
}

func NewTubeSet(name string) *beanstalk.TubeSet {
	conn := NewConn()
	t := beanstalk.NewTubeSet(conn, name)
	return t
}
