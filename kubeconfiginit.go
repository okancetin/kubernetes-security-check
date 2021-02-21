import (
	"fmt"
	"flag"
	"os"
	"os/user"
	"path/filepath"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig, uri string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "",
		"kubeconfig path.")

	flag.StringVar(&uri, "master", "",
		"The uri of the Kubernetes API server. ")
}
func GetConfig() (*rest.Config, error) {
	if len(kubeconfig) > 0 {
		return clientcmd.BuildConfigFromFlags(uri, kubeconfig)
	}
	if len(os.Getenv("KUBECONFIG")) > 0 {
		return clientcmd.BuildConfigFromFlags(uri, os.Getenv("KUBECONFIG"))
	}
	if c, err := rest.InClusterConfig(); err == nil {
		return c, nil
	}
	if usr, err := user.Current(); err == nil {
		if c, err := clientcmd.BuildConfigFromFlags(
			"", filepath.Join(usr.HomeDir, ".kube", "config")); err == nil {
			return c, nil
		}
	}

	return nil, fmt.Errorf("could not find a kubeconfig")
}

func GetConfigFromFile(kubeconfig string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}