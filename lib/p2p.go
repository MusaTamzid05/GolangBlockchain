package lib

import (
    "net"
    "fmt"
    "encoding/gob"
)

type P2P struct {
    addr string
    Server net.Listener 
    peers []net.Conn // clients
}

func MakeP2P(addr string) (P2P, error) {
    p2p := P2P{addr: addr}
    listener, err := net.Listen("tcp", addr)

    if err != nil {
        return p2p, err

    }

    p2p.Server = listener

    return p2p, nil
}

func (p *P2P) StartServer() {
    fmt.Println("Running the P2P Server at ", p.addr)
    defer p.Server.Close()

    serverRunning := true

    for serverRunning {
        peer, err := p.Server.Accept()

        if err != nil {
            fmt.Println(err)
            continue
        }

        fmt.Println("Peer connected ", peer.RemoteAddr()) // @TODO Fix the addr value
        p.peers = append(p.peers, peer)
        go p.ListenToPeer(peer, true)

    }
}



func (p *P2P) ListenToPeer(peer net.Conn, isServer bool) {
    peerRunning := true

    defer peer.Close()

    if isServer {
        p.Send(peer, CurrentBlockchain)

    }

    for peerRunning {
        var peerBlockChain BlockChain
        decoder := gob.NewDecoder(peer)
        decoder.Decode(&peerBlockChain)

        peerBlockChain.Show()

        flag := CurrentBlockchain.Replace(peerBlockChain)

        if flag {
            fmt.Println("Blockchain replace with peer ", peer.RemoteAddr())
        }

    }
}

func (p *P2P) Send(peer net.Conn, blockchain BlockChain) {
    encoder := gob.NewEncoder(peer)
    err := encoder.Encode(blockchain)

    if err != nil {
        fmt.Println("Error sending blockchain ", err.Error())

    }
}

func (p *P2P) AddPeers(peerAddrs []string) {
    for _, peerAddr  := range peerAddrs {
        //  normal client creation in golang
        peer, err := net.Dial("tcp", peerAddr)

        if err != nil {
            fmt.Println("Peer Add error ", peer.LocalAddr(), " ", err.Error())
            continue

        }

        p.peers = append(p.peers, peer)
        p.ListenToPeer(peer, false)
    }

}





