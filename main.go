package main

import (
	"fmt"
	"os"

	"github.com/driver-eks/cluster"
	"github.com/driver-eks/nodegroup"
)

func main() {
	Credenthials()

	for {
		fmt.Println("-----=------DRIVER EKS-----=------")
		Menu()
		comando := insereComando()
		switch comando {
		case 1:
			fmt.Println("=== Creating Cluster Selected - Please wait..")
			//statecluster := cluster.StateCluster()
			cluster.CreateCluster()
			os.Exit(1)
		case 2:
			fmt.Println("=== Updating Cluster Selected - Please wait..")
			cluster.UpdateCluster()
			os.Exit(1)
		case 3:
			fmt.Println("=== Listing Cluster Selected - Please wait..")
			cluster.ListClusters()
			os.Exit(1)

		case 4:
			fmt.Println("=== Creating NodeGroup Selected - Please wait..")
			nodegroup.CreateNodeGroup()
			os.Exit(1)
		case 5:
			fmt.Println("=== Updating NodeGroup Selected - Please wait..")
			nodegroup.UpdateNodeGroup()
			os.Exit(1)
		case 6:
			fmt.Println("=== Listing NodeGroups Selected - Please wait..")
			nodegroup.ListNodegroups()
			os.Exit(1)
		case 7:
			fmt.Println("exiing..")
			os.Exit(1)

		default:
			fmt.Println("Invalid option")
			os.Exit(-1)
		}

	}

}

func Menu() {
	fmt.Println("[1] - Create Cluster")
	fmt.Println("[2] - Update Cluster")
	fmt.Println("[3] - List Clusters")
	fmt.Println("-------------------")
	fmt.Println("[4] - Create  NodeGroup")
	fmt.Println("[5] - Update  NodeGroup")
	fmt.Println("[6] - List NodeGroups")
}

func insereComando() int {
	var valor int
	fmt.Scan(&valor)
	return valor
}

func Credenthials() {
	os.Getenv("AWS_ACCESS_KEY_ID")
	os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := "us-west-2"

	os.Setenv("AWS_REGION", region)

}
