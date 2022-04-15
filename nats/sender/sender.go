package sender

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"time"
)

var counter int

func Send() {
	st, _ := stan.Connect("prod", "simple-sender")

	start := make(chan struct{})

	Sender(st, start)
	for {
		start <- struct{}{}
		time.Sleep(500 * time.Millisecond)
	}
}

func Sender(st stan.Conn, start <-chan struct{}) {
	go func() {
		for {
			select {
			case <-start:
				for i := 1; i <= 10; i++ {
					if i == 7 {
						a, err := ioutil.ReadFile("./sender/jsons/test7.xml")
						if err != nil {
							fmt.Println(err)
						}
						err = st.Publish("channel", a)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						data, err := ioutil.ReadFile(fmt.Sprintf("./sender/jsons/test%d.json", i))
						if err != nil {
							fmt.Println(err)
						}

						newData := ChangeId(data)
						log.Printf("sendindg %d data\n", counter)
						counter++

						err = st.Publish("channel", newData)
						if err != nil {
							fmt.Println(err)
						}
					}
					time.Sleep(3 * time.Second)
				}
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}
