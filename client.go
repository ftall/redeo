package redeo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type clientSlice []Client

func (p clientSlice) Len() int           { return len(p) }
func (p clientSlice) Less(i, j int) bool { return p[i].ID < p[j].ID }
func (p clientSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var clientInc = uint64(0)

// A client is the origin of a request
type Client struct {
	baseInfo
	ID         uint64      `json:"id,omitempty"`
	RemoteAddr string      `json:"remote_addr,omitempty"`
	Ctx        interface{} `json:"ctx,omitempty"`

	cmd   string
	mutex sync.Mutex
}

// NewClient creates a new client info container
func NewClient(addr string) *Client {

	return &Client{
		ID:         atomic.AddUint64(&clientInc, 1),
		RemoteAddr: addr,
		baseInfo:   baseInfo{StartTime: time.Now()},
	}
}

// OnCommand callback to track user command
func (i *Client) OnCommand(cmd string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.cmd = cmd
}

// Command returns the last user command
func (i *Client) LastCommand() string {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	return i.cmd
}

// String generates an info string
func (i *Client) String() string {
	return fmt.Sprintf("id=%d addr=%s age=%d cmd=%s", i.ID, i.RemoteAddr, i.Uptime()/time.Second, i.LastCommand())
}
