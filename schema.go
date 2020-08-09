package main

type Deployment struct {
	ApiVersion 		string `yaml:"apiVersion"`
	Kind			string `yaml:"kind"`
	Meta 			struct{
		Name 		string `yaml:"name"`
	} `yaml:"metadata"`
	Spec			struct{
		Template 	struct{
			Spec	struct{
				Containers []struct{
					Image			string `yaml:"image"`
				}`yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

type Rollout struct {
	ApiVersion 		string `yaml:"apiVersion"`
	Kind			string `yaml:"kind"`
	Meta 			struct{
		Name 		string `yaml:"name"`
	} `yaml:"metadata"`
	Spec			struct{
		Template 	struct{
			Spec	struct{
				Containers []struct{
					Image			string `yaml:"image"`
				}`yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}