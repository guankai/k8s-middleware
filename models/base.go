package models

import (
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	k8s_client "bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"
	"crypto/tls"
	"crypto/x509"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

var Client k8s_client.Client

func init() {
	k8s_url := beego.AppConfig.String("kubernetes_url")
	k8s_schema := beego.AppConfig.String("kubernetes_schema")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	logs.Debug("Kubernetes server: %s, schema: %s, username: %s", k8s_url, k8s_schema, username)

	if k8s_schema == "http" {
		Client = k8s_client.NewClient(k8s_url, username, password)
	} else if k8s_schema == "https" {
		// Load client cert
		cert, err := tls.LoadX509KeyPair("/Users/Allan/.minikube/apiserver.crt", "/Users/Allan/.minikube/apiserver.key")
		if err != nil {
			panic(err)
		}

		// Load CA cert
		caCert, err := ioutil.ReadFile("/Users/Allan/.minikube/ca.crt")
		if err != nil {
			panic(err.Error())
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// set HTTPS client
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}

		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}

		http_client := &http.Client{Transport: transport}

		Client = k8s_client.NewClientWithHttpClient(http_client, k8s_url, username, password)
	} else {
		panic("unknown kubernetes_schema, only `http` or `https` allowed")
	}
}

// these ***copy struct just for shut up the bee gendoc, no other use
type ReplicationControllerCopy struct {
	v1.ReplicationController
}

type ServiceCopy struct {
	v1.Service
}
