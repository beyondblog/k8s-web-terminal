package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type K8sClient struct {
	Host string

	dockerClient map[string]*DockerClient
}

type NodeList struct {
	ApiVersion string `json:"apiVersion,omitempty" yaml:"api_version,omitempty"`

	Items []Node `json:"items,omitempty" yaml:"items,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Metadata *ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type ListMeta struct {
	ResourceVersion string `json:"resourceVersion,omitempty" yaml:"resource_version,omitempty"`

	SelfLink string `json:"selfLink,omitempty" yaml:"self_link,omitempty"`
}

type Node struct {
	ApiVersion string `json:"apiVersion,omitempty" yaml:"api_version,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Status *NodeStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ObjectMeta struct {
	Annotations map[string]interface{} `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	CreationTimestamp string `json:"creationTimestamp,omitempty" yaml:"creation_timestamp,omitempty"`

	DeletionGracePeriodSeconds int64 `json:"deletionGracePeriodSeconds,omitempty" yaml:"deletion_grace_period_seconds,omitempty"`

	DeletionTimestamp string `json:"deletionTimestamp,omitempty" yaml:"deletion_timestamp,omitempty"`

	GenerateName string `json:"generateName,omitempty" yaml:"generate_name,omitempty"`

	Generation int64 `json:"generation,omitempty" yaml:"generation,omitempty"`

	Labels map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	ResourceVersion string `json:"resourceVersion,omitempty" yaml:"resource_version,omitempty"`

	SelfLink string `json:"selfLink,omitempty" yaml:"self_link,omitempty"`

	Uid string `json:"uid,omitempty" yaml:"uid,omitempty"`
}

type NodeStatus struct {
	Addresses []NodeAddress `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	Capacity map[string]interface{} `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	NodeInfo *NodeSystemInfo `json:"nodeInfo,omitempty" yaml:"node_info,omitempty"`
}

type NodeAddress struct {
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type NodeSystemInfo struct {
	BootID string `json:"bootID,omitempty" yaml:"boot_id,omitempty"`

	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty" yaml:"container_runtime_version,omitempty"`

	KernelVersion string `json:"kernelVersion,omitempty" yaml:"kernel_version,omitempty"`

	KubeProxyVersion string `json:"kubeProxyVersion,omitempty" yaml:"kube_proxy_version,omitempty"`

	KubeletVersion string `json:"kubeletVersion,omitempty" yaml:"kubelet_version,omitempty"`

	MachineID string `json:"machineID,omitempty" yaml:"machine_id,omitempty"`

	OsImage string `json:"osImage,omitempty" yaml:"os_image,omitempty"`

	SystemUUID string `json:"systemUUID,omitempty" yaml:"system_uuid,omitempty"`
}

func (client *K8sClient) Nodes() []Node {
	resp, err := http.Get(client.Host + "/api/v1/nodes")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var nodeList NodeList
	json.Unmarshal(body, &nodeList)
	return nodeList.Items

}

func (client *K8sClient) GetContainer(node string) []Container {
	if _, ok := client.dockerClient[node]; !ok {
		//using default port
		client.dockerClient[node] = &DockerClient{"http://" + node + ":2375"}
	}

	return client.dockerClient[node].ListContainers()
}
