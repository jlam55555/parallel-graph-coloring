package main

import (
	"bufio"
	"flag"
	"graphnet"
	"net"
	"os"
	"proj2/common"
	"sync"
)

// main is the driver to be built into the executable for the server
func main() {
	// create logger
	logger, logFile := common.CreateLogger("server")
	defer func() {
		err := logFile.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Takes in command line flag(s)
	configFile := flag.String("config", "",
		"File containing the node configurations")
	flag.Parse()
	if *configFile == "" {
		logger.Fatal("No configuration file.")
	}

	// read node configuration file
	// https://stackoverflow.com/questions/8757389
	file, err := os.Open(*configFile)
	if err != nil {
		logger.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	addresses := make([]string, 0)
	for fileScanner.Scan() {
		logger.Printf("Reading config: %s\n", fileScanner.Text())
		addresses = append(addresses, fileScanner.Text())
	}

	// create node connection pool
	ncp := graphnet.NewNodeConnPool()

	// create server dispatch table
	// TODO: fill with actual handlers
	dispatchTab := make(map[byte]graphnet.Dispatch)

	var wg sync.WaitGroup
	nWorkers := len(addresses)
	wg.Add(nWorkers)
	dispatchTab[graphnet.MSG_VERTEX_INFO] = func(vertexInfo []byte) {
		logger.Printf("Indexes %d, %d, %d, %d have been updated to %d, %d, %d, %d.",
			vertexInfo[4], vertexInfo[5], vertexInfo[6], vertexInfo[7],
			vertexInfo[0], vertexInfo[1], vertexInfo[2], vertexInfo[3])
	}
	dispatchTab[graphnet.MSG_NODE_FINISHED] = func(nodeIndex []byte) {
		wg.Done()
		logger.Printf("Node %d has finished processing.\n",
			nodeIndex[0])
	}
	dispatchTab[graphnet.MSG_NODE_ROUND_FINISHED] = func(nodeIndex []byte) {
		logger.Printf("Node %d has finished a round.\n",
			nodeIndex[0])
	}
	dispatchTab[graphnet.MSG_NODE_INDEX_COUNT] = func(indexCount []byte) {
		logger.Printf("Node %d has %d total nodes.",
			indexCount[0], indexCount[1])
	}
	dispatchTab[graphnet.MSG_NODE_IP] = func(ip []byte) {
		logger.Printf("Node %d has IP of %d.%d.%d.%d and port of %d",
			ip[1], ip[2], ip[3], ip[4], ip[5:])
	}

	// establish a connection with each node from configuration file
	for i, address := range addresses {
		logger.Printf("Establishing connection with %s...\n", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("Connection established with %s.\n", address)
		nodeConn := graphnet.NewNodeConn(conn, logger, dispatchTab)
		nodeConn.Index = i + 1
		ncp.AddUnregistered(nodeConn)
	}

	// send information about all nodes to each node
	for i, nodeConn := range ncp {
		// TODO: send node index to node

		// TODO: send total nodes count to node

		// TODO: send addresses of higher indexed nodes to node
		for j := i + 1; j < len(addresses); j++ {
		}

		// TODO: end packet
	}

	// this shouldn't have any effect for the server, since all nodes were
	// added in order
	ncp.Register()

	// TODO: send subgraphs

	// TODO: start coloring
	// distributed.ColorDistributedServer()
}
