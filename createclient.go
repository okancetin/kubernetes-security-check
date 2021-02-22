func createK8sClient(endpoint, kubeCfgPath string) (*k8s.Clientset, error) {
	var (
		config *k8sRest.Config
		err    error
	)
	if kubeCfgPath != "" {
		config, err = k8sClientCmd.BuildConfigFromFlags("", kubeCfgPath)
	} else {
		config = &k8sRest.Config{Host: endpoint}
		err = k8sRest.SetKubernetesDefaults(config)
	}
	if err != nil {
		return nil, err
	}
	return k8s.NewForConfig(config)
}

func NewInClusterClientWithEndpoint(endpoint string) (Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	config.Host = endpoint

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		clientset: clientset,
	}, nil
}