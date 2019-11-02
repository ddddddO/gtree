package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	//"gopkg.in/yaml.v2"
	"github.com/ghodss/yaml"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

/*
	サンプル
	http://ks888.hatenablog.com/entry/2017/12/25/001259

	網羅的：
	https://pizi.netlify.com/posts/kubernetes-client-go-example/

	公式
	https://godoc.org/k8s.io/client-go

	on GCP
	https://github.com/kubernetes/client-go/tree/master/examples

	ビルドの際にハマった
	https://github.com/kubernetes/client-go/issues/584
*/
func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	/*
		if err := getPods(client); err != nil {
			log.Fatal(err)
		}
	*/

	if err := createDeployment(client); err != nil {
		log.Fatalf("%+v", err)
	}

}

func newClient() (*kubernetes.Clientset, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return kubernetes.NewForConfig(kubeConfig)
}

func getPods(client *kubernetes.Clientset) error {
	pods, err := client.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	return nil
}

// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
// クラスタにDeploymentをデプロイする
func createDeployment(client *kubernetes.Clientset) error {
	// https://godoc.org/k8s.io/client-go/kubernetes/typed/apps/v1
	deploymentClient := client.AppsV1().Deployments(apiv1.NamespaceDefault)

	// subのdeployment
	// https://godoc.org/k8s.io/api/apps/v1#Deployment
	deployment, err := convertYAMLtoStruct()
	if err != nil {
		return err
	}

	_, err = deploymentClient.Create(deployment)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func convertYAMLtoStruct() (*appsv1.Deployment, error) {
	deploymentFile, err := os.Open("../../deployments/02_k8s_gke/deployment/deploy-sub.yml")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer deploymentFile.Close()

	fi, err := deploymentFile.Stat()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := make([]byte, fi.Size())
	_, err = deploymentFile.Read(buf)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	deployment := &appsv1.Deployment{}
	if err := yaml.Unmarshal(buf, deployment); err != nil {
		return nil, errors.WithStack(err)
	}

	return deployment, nil
}
