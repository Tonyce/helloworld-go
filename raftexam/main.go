package main

import (
	"flag"
	"strings"

	"go.etcd.io/etcd/raft/v3/raftpb"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	proposeC := make(chan string)
	defer close(proposeC)

	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	var kvs *kvstore
	getSnapshot := func() ([]byte, error) {
		return kvs.getSnapshot()
	}

	commitC, errorC, snapshotterReady := newRaftNode(
		*id,
		strings.Split(*cluster, ","),
		*join,
		getSnapshot,
		proposeC,
		confChangeC,
	)

	_, _, _ = commitC, errorC, snapshotterReady
	_ = kvport
	kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	serveHttpKVAPI(kvs, *kvport, confChangeC, errorC)
	// select {}
}
