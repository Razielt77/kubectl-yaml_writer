package schema

/*
apiVersion: v1
kind: Service
metadata:
  name: gitopsapp
  spec:
    ports:
     - port: 80
       targetPort: 3000
    selector:
      app: gitopsapp
    type: LoadBalancer
*/

type Service struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Meta       MetaData `yaml:"metadata,omitempty"`
	Spec       struct {
		Selector *map[string]string `yaml:"selector,omitempty"`
		Ports    *[]Port            `yaml:"ports,omitempty"`
		Type     *string            `yaml:"type,omitempty"`
	} `yaml:"spec,omitempty"`
}

type Port struct {
	Protocol   *string `yaml:"protocol,omitempty"`
	Port       *string `yaml:"port,omitempty"`
	TargetPort *string `yaml:"targetport,omitempty"`
}

func (p *Port) Init(targetPort, externalPort string) {
	p.Protocol = new(string)
	*p.Protocol = "TCP"
	p.Port = new(string)
	*p.Port = externalPort
	p.TargetPort = new(string)
	*p.TargetPort = targetPort
}

func (s *Service) Init(app, targetPort, externalPort string) {
	s.ApiVersion = "apps/v1"
	s.Kind = "Service"
	s.Meta.Init(app+"_service", app)
	s.Spec.Selector = new(map[string]string)
	*s.Spec.Selector = make(map[string]string)
	(*s.Spec.Selector)["app"] = app
	s.Spec.Ports = new([]Port)
	*s.Spec.Ports = append(*s.Spec.Ports, *new(Port))
	(*s.Spec.Ports)[0].Init(targetPort, externalPort)
	s.Spec.Type = new(string)
	*s.Spec.Type = "LoadBalancer"
}
