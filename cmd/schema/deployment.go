package schema

import (
	"fmt"
)

type MetaData  struct {
	Name string `yaml:"name"`
	Labels map[string]string `yaml:"labels,omitempty"`
}

func (m *MetaData)Init(name, app string){
	m.Name = name
	m.Labels = make(map[string]string)
	m.Labels["app"]=app
}

type Container struct {
	Image string `yaml:"image"`
	Name  string `yaml:"name"`
	Ports []ContainerPort `yaml:"ports"`
}

type ContainerPort struct {
	ContPort int `yaml:"containerPort"`
}

func (c *Container) Init (image,name string, port int) {
	c.Image = image
	c.Name = name
	c.Ports = append(c.Ports,ContainerPort{port})
}

type Selector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
}

func (s *Selector)Init(app string){
	s.MatchLabels = make(map[string]string)
	s.MatchLabels["app"] = app
}

type TemplateMetadata struct {
	Labels map[string]string `yaml:"labels"`
}

func (m *TemplateMetadata)Init(app string){
	m.Labels = make(map[string]string)
	m.Labels["app"] = app
}

type BaseInfo struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Meta       MetaData `yaml:"metadata,omitempty"`
}


type Deployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Meta       MetaData `yaml:"metadata,omitempty"`
	Spec struct {
		Replicas             int `yaml:"replicas"`
		RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
		SelectorObj    *Selector `yaml:"selector,omitempty"`
		Template struct {
			MetadataObj *TemplateMetadata `yaml:"metadata,omitempty"`
			Spec struct {
				Containers *[]Container `yaml:"containers,omitempty"`
			} `yaml:"spec"`
		} `yaml:"template"`
		MinReadySeconds int `yaml:"minReadySeconds,omitempty"`
	} `yaml:"spec"`
}

type Rollout struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Meta       *MetaData `yaml:"metadata",omitempty`
	Spec struct {
		Replicas             int `yaml:"replicas"`
		RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
		SelectorObj    *Selector `yaml:"selector,omitempty"`
		Template struct {
			MetadataObj *TemplateMetadata `yaml:"metadata,omitempty"`
			Spec struct {
				Containers *[]Container `yaml:"containers,omitempty"`
			} `yaml:"spec"`
		} `yaml:"template"`
		MinReadySeconds int `yaml:"minReadySeconds,omitempty"`
		Strategy        struct {
			CanarySteps *Canary `yaml:"canary,omitempty"`
		} `yaml:"strategy"`
	} `yaml:"spec"`
}

type Canary struct {
	Steps []CanaryStep `yaml:"steps,omitempty"`
}

type CanaryStep struct {
	// SetWeight sets what percentage of the newRS should receive
	SetWeight *int32 `yaml:"setWeight,omitempty"`
	// Pause freezes the rollout by setting spec.Paused to true.
	// A Rollout will resume when spec.Paused is reset to false.
	// +optional
	Pause *RolloutPause `yaml:"pause,omitempty"`
}

type RolloutPause struct {
	// Duration the amount of time to wait before moving to the next step.
	// +optional
	Duration *int `yaml:"duration,omitempty"`
}

func (dp *Deployment) Init(name,app,image string,replica,port int){
	dp.ApiVersion = "apps/v1"
	dp.Kind = "Deployment"
	dp.Meta.Init(name,app)
	dp.Spec.Replicas = replica
	dp.Spec.RevisionHistoryLimit = 3
	dp.Spec.SelectorObj = new(Selector)
	dp.Spec.SelectorObj.Init(app)
	dp.Spec.Template.MetadataObj = new(TemplateMetadata)
	dp.Spec.Template.MetadataObj.Init(app)

	dp.Spec.Template.Spec.Containers = new([]Container)
	*dp.Spec.Template.Spec.Containers = append(*dp.Spec.Template.Spec.Containers,*new(Container))
	(*dp.Spec.Template.Spec.Containers)[0].Init(image,name,port)
}

func (dp *Deployment) Update(att,value string,index int) error{
	var err error = nil
	switch att{
	case "image":
		if (*dp.Spec.Template.Spec.Containers)[index].Image != value{
			(*dp.Spec.Template.Spec.Containers)[index].Image = value
		}else{
			fmt.Printf("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	return err
}

func (rl *Rollout) Update(att,value string,index int) error{
	var err error = nil
	switch att{
	case "image":
		if (*rl.Spec.Template.Spec.Containers)[index].Image != value{
			(*rl.Spec.Template.Spec.Containers)[index].Image = value
		}else{
			fmt.Printf("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	return err
}