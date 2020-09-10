package cmd

import (
	"fmt"
)

type metaData struct {
	Name   string            `yaml:"name"`
	Labels map[string]string `yaml:"labels,omitempty"`
}

func (m *metaData) Init(name, app string) {
	m.Name = name
	m.Labels = make(map[string]string)
	m.Labels["app"] = app
}

type container struct {
	Image string          `yaml:"image"`
	Name  string          `yaml:"name"`
	Ports []containerPort `yaml:"ports"`
}

type containerPort struct {
	ContPort int `yaml:"containerPort"`
}

func (c *container) Init(image, name string, port int) {
	c.Image = image
	c.Name = name
	c.Ports = append(c.Ports, containerPort{port})
}

type selector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
}

func (s *selector) Init(app string) {
	s.MatchLabels = make(map[string]string)
	s.MatchLabels["app"] = app
}

type templateMetadata struct {
	Labels map[string]string `yaml:"labels"`
}

func (m *templateMetadata) Init(app string) {
	m.Labels = make(map[string]string)
	m.Labels["app"] = app
}

type baseInfo struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Meta       metaData `yaml:"metadata,omitempty"`
}

type deployment struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Meta       metaData `yaml:"metadata,omitempty"`
	Spec       struct {
		Replicas             int       `yaml:"replicas"`
		RevisionHistoryLimit int       `yaml:"revisionHistoryLimit"`
		SelectorObj          *selector `yaml:"selector,omitempty"`
		Template             struct {
			MetadataObj *templateMetadata `yaml:"metadata,omitempty"`
			Spec        struct {
				Containers *[]container `yaml:"containers,omitempty"`
			} `yaml:"spec"`
		} `yaml:"template"`
		MinReadySeconds int `yaml:"minReadySeconds,omitempty"`
	} `yaml:"spec"`
}

type rollout struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Meta       metaData `yaml:"metadata",omitempty`
	Spec       struct {
		Replicas             int       `yaml:"replicas"`
		RevisionHistoryLimit int       `yaml:"revisionHistoryLimit"`
		SelectorObj          *selector `yaml:"selector,omitempty"`
		Template             struct {
			MetadataObj *templateMetadata `yaml:"metadata,omitempty"`
			Spec        struct {
				Containers *[]container `yaml:"containers,omitempty"`
			} `yaml:"spec"`
		} `yaml:"template"`
		MinReadySeconds int `yaml:"minReadySeconds,omitempty"`
		Strategy        struct {
			CanarySteps *Canary `yaml:"canary,omitempty"`
		} `yaml:"strategy"`
	} `yaml:"spec"`
}

func (r *rollout) Init(name, app, image string, replica, port int) {
	r.ApiVersion = "apps/v1"
	r.Kind = "Rollout"
	r.Meta.Init(name, app)
	r.Spec.Replicas = replica
	r.Spec.RevisionHistoryLimit = 3
	r.Spec.SelectorObj = new(selector)
	r.Spec.SelectorObj.Init(app)
	r.Spec.Template.MetadataObj = new(templateMetadata)
	r.Spec.Template.MetadataObj.Init(app)

	r.Spec.Template.Spec.Containers = new([]container)
	*r.Spec.Template.Spec.Containers = append(*r.Spec.Template.Spec.Containers, *new(container))
	(*r.Spec.Template.Spec.Containers)[0].Init(image, app, port)
	r.Spec.MinReadySeconds = 30
	r.Spec.Strategy.CanarySteps = new(Canary)
	r.Spec.Strategy.CanarySteps.Steps = append(r.Spec.Strategy.CanarySteps.Steps, CanaryStep{})
	r.Spec.Strategy.CanarySteps.Steps[0].SetWeight = new(int32)
	*r.Spec.Strategy.CanarySteps.Steps[0].SetWeight = 50
	r.Spec.Strategy.CanarySteps.Steps = append(r.Spec.Strategy.CanarySteps.Steps, CanaryStep{})
	(*r.Spec.Strategy.CanarySteps).Steps[1].Pause = new(RolloutPause)
	//r.Spec.Strategy.CanarySteps.Steps
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

func (dp *deployment) Init(name, app, image string, replica, port int) {
	dp.ApiVersion = "apps/v1"
	dp.Kind = "Deployment"
	dp.Meta.Init(name, app)
	dp.Spec.Replicas = replica
	dp.Spec.RevisionHistoryLimit = 3
	dp.Spec.SelectorObj = new(selector)
	dp.Spec.SelectorObj.Init(app)
	dp.Spec.Template.MetadataObj = new(templateMetadata)
	dp.Spec.Template.MetadataObj.Init(app)

	dp.Spec.Template.Spec.Containers = new([]container)
	*dp.Spec.Template.Spec.Containers = append(*dp.Spec.Template.Spec.Containers, *new(container))
	(*dp.Spec.Template.Spec.Containers)[0].Init(image, app, port)
}

func (dp *deployment) Update(att, value string, index int) error {
	var err error = nil
	switch att {
	case "image":
		if (*dp.Spec.Template.Spec.Containers)[index].Image != value {
			(*dp.Spec.Template.Spec.Containers)[index].Image = value
		} else {
			fmt.Printf("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	return err
}

func (rl *rollout) Update(att, value string, index int) error {
	var err error = nil
	switch att {
	case "image":
		if (*rl.Spec.Template.Spec.Containers)[index].Image != value {
			(*rl.Spec.Template.Spec.Containers)[index].Image = value
		} else {
			fmt.Printf("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	return err
}
