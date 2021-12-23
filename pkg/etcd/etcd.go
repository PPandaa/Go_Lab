package etcd

import (
	"context"
	"time"

	"GoLab/guard"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/client/v3/namespace"
)

type EtcdConn struct {
	Enable   bool
	Host     string
	Port     string
	Username string
	Password string
}

func (c *EtcdConn) new(enable, host, port, username, password string) {

	var e bool
	if enable == "true" {
		e = true
	} else if enable == "false" {
		e = false
	}
	c.Enable = e
	c.Host = host
	c.Port = port
	c.Username = username
	c.Password = password

}

type EtcdCli struct {
	client         *clientv3.Client
	leaseID        *clientv3.LeaseID
	session        *concurrency.Session
	election       *concurrency.Election
	namespace      string
	serviceName    string
	serviceID      string
	serviceVersion string

	rootKV  clientv3.KV
	watcher clientv3.Watcher
}

func (cli *EtcdCli) connect(ec *EtcdConn) *clientv3.Client {

	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{ec.Host + ":" + ec.Port},
			Username:    ec.Username,
			Password:    ec.Password,
			DialTimeout: 5 * time.Second,
		},
	)
	if err != nil {
		guard.Logger.Error("ETCD Connect Failed: " + err.Error())
	}
	cli.client = client
	guard.Logger.Info("ETCD Connect Success")
	return client

}

func (cli *EtcdCli) do() {

	cli.rootKV = cli.client.KV
	cli.watcher = clientv3.NewWatcher(cli.client)
	cli.client.KV = namespace.NewKV(cli.client.KV, cli.namespace)

	lease := clientv3.NewLease(cli.client) //new a lease
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	leaseGrantRes, err := lease.Grant(ctx, 60) //grant create a new lease
	if err != nil {
		panic(err)
	}

	cli.client.Lease = lease
	// cli.Client.KeepAlive(context.Background(), leaseGrantRes.ID)
	cli.leaseID = &leaseGrantRes.ID
	// fmt.Println("leaseID:", *cli.leaseID)

	// election
	s, err := concurrency.NewSession(cli.client, concurrency.WithLease(*cli.leaseID))
	if err != nil {
		panic(err)
	}
	// fmt.Printf("session: %+v\n", s)
	cli.session = s
	// cli.Session2 = *s

}

func (cli *EtcdCli) start(etcdConn *EtcdConn, isApp bool, serviceName string, serviceVersion string) {

	cli.connect(etcdConn)
	if isApp {
		cli.namespace = "iapps/"
	} else {
		cli.namespace = "services/"
	}
	cli.serviceName = serviceName
	cli.serviceVersion = serviceVersion
	cli.serviceID = time.Now().Format(time.RFC3339)
	cli.do()

}

func (cli *EtcdCli) startElection() {

	cli.election = concurrency.NewElection(cli.session, "election/"+cli.serviceName)
	// fmt.Printf("Session:%+v\n", cli.session)
	// fmt.Printf("ServiceName:%+v\n", cli.serviceName)
	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			err := cli.election.Campaign(ctx, cli.serviceVersion+"/"+cli.serviceID)
			if err != nil {
				guard.Logger.Error("ETCD Failed To Campaign: " + err.Error())
			} else {
				guard.Logger.Info("ETCD Become The Leader!")
				break
			}
		}
	}()

}

func (cli *EtcdCli) PutMetadata(key string, val string) (*clientv3.PutResponse, error) {

	return cli.client.Put(
		context.Background(),
		cli.serviceName+"/"+cli.serviceVersion+"/"+cli.serviceID+"/"+key,
		val,
		clientv3.WithLease(*cli.leaseID),
	)

}

func (cli *EtcdCli) Close() {

	if cli.election != nil {
		err := cli.election.Resign(context.Background())
		if err != nil {
			println(err)
		}
	}
	if cli.session != nil {
		err := cli.session.Close()
		if err != nil {
			println(err)
		}
	}
	if cli.client != nil {
		// e.Client.Lease.Revoke(context.Background(), *e.leaseID)
		// e.Client.Lease.Close()
		err := cli.client.Close()
		if err != nil {
			println(err)
		}
	}

}
