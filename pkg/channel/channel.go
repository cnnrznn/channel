package channel

import (
    "fmt"
    "net"
)

type Channel struct {
    Id int
    Peers []string
}

func (c Channel) String() string {
    return fmt.Sprintf("{%v, %v}", c.Id, c.Peers)
}

func (c Channel) PingAll() {
    for index, _ := range c.Peers {
        go c.Send("ping", index)
    }
}

func (c Channel) Send(msg string, index int) {
    conn, _ := net.Dial("udp", c.Peers[index])
    defer conn.Close()

    conn.Write([]byte(msg))
}

func (c Channel) Broadcast(msg string) {
    for index, _ := range c.Peers {
        go c.Send(msg, index)
    }
}

func (c Channel) Serve(ch chan string) {
    pc, _ := net.ListenPacket("udp", c.Peers[c.Id])
    defer pc.Close()

    buffer := make([]byte, 1024)

    for {
        n, _, err := pc.ReadFrom(buffer)
        if err != nil {
            continue
        }
        ch <- string(buffer[:n])
    }
}
