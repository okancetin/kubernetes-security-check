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