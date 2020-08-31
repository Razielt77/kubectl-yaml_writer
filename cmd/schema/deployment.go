package schema

import (
	"errors"
	"fmt"
	"github.com/razielt77/kyml/cmd/utils"
)

type MetaData  struct {
	Name string `yaml:"name"`
	Labels map[string]string `yaml:"labels,omitempty"`
}

type Container struct {
	Image string `yaml:"image"`
	Name  string `yaml:"name"`
	Ports []struct {
		ContainerPort int `yaml:"containerPort"`
	} `yaml:"ports"`
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
		Selector             struct {
			MatchLabels map[string]string `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels map[string]string `yaml:"labels"`
			} `yaml:"metadata"`
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
	Meta       MetaData `yaml:"metadata",omitempty`
	Spec struct {
		Replicas             int `yaml:"replicas"`
		RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
		Selector             struct {
			MatchLabels map[string]string `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels map[string]string `yaml:"labels"`
			} `yaml:"metadata"`
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
func (dp *Deployment) Update(att,value string,index int) error{
	var err error = nil
	switch att{
	case "image":
		if (*dp.Spec.Template.Spec.Containers)[index].Image != value{
			(*dp.Spec.Template.Spec.Containers)[index].Image = value
		}else{
			err = errors.New("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	utils.DieOnError(err)
	return err
}

func (rl *Rollout) Update(att,value string,index int) error{
	var err error = nil
	switch att{
	case "image":
		if (*rl.Spec.Template.Spec.Containers)[index].Image != value{
			(*rl.Spec.Template.Spec.Containers)[index].Image = value
		}else{
			err = errors.New("value was already set")
		}
	default:
		err = fmt.Errorf("attribute: %s is not supported", att)
	}
	utils.DieOnError(err)
	return err
}