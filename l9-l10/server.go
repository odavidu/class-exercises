package main

import (
	"fmt"
	"log"
	"maps"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

type PeerServer struct {
	Address string
	Client  *rpc.Client
}

type Args struct {
	GossipLive map[string]int
	Round      int
	Sender     string
}

type Server struct {
	live    map[string]int
	lock    sync.Mutex
	Round   int
	Address string
	peers   []PeerServer
}

var lock sync.Mutex

func (t *Server) Heartbeat(args *Args, reply *int) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if args.Round > t.Round {
		t.Round = args.Round
	}
	t.live[args.Sender] = t.Round

	for server, newRound := range args.GossipLive {
		if round, ok := t.live[server]; ok {
			if newRound > round {
				t.live[server] = newRound
			}
		} else {
			continue
		}
	}
	return nil
}

func (t *Server) sendHeartbeat(to PeerServer) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.Round++
	args := Args{
		GossipLive: maps.Clone(t.live),
		Round:      t.Round,
		Sender:     t.Address,
	}
	reply := 0
	err := to.Client.Call("Server.Heartbeat", args, &reply)
	if err != nil {
		log.Println("RPC error:", err)
	}
}

func (t *Server) GenerateReport() {
	fmt.Printf("Report [Round %d]\n", t.Round)
	fmt.Println(t.live)
	fmt.Println("Live Servers: ")
	for server, round := range t.live {
		if math.Abs(float64(t.Round-round)) <= 10 {
			fmt.Println(server)
		}
	}
}

func main() {

	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	my_address := "10.239.187.225"
	server.Address = my_address
	server.Round = 0
	server.peers = make([]PeerServer, 0)
	server.live = make(map[string]int)
	peer_addresses := []string{
		"10.239.23.111:1234",
		//"10.239.50.138:1234",
	}

	time.Sleep(10 * time.Second) // WAIT to start other servers

	for _, addr := range peer_addresses {
		if addr == my_address {
			continue
		}
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		server.peers = append(server.peers, PeerServer{addr, client})
	}

	heartbeat := time.After(time.Duration(1) * time.Second)
	report := time.After(time.Duration(5) * time.Second)

	for {
		select {
		case <-heartbeat:
			randomServer := server.peers[rand.Intn(len(server.peers))]
			server.sendHeartbeat(randomServer)

			fmt.Println("Sending heartbeat to", randomServer.Address)

			heartbeat = time.After(time.Duration(1) * time.Second)
		case <-report:
			server.GenerateReport()

			report = time.After(time.Duration(5) * time.Second)
		default:
			continue
		}
	}

}
