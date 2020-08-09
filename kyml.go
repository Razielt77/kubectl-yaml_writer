package main
// use the command below to build it before packaging it in a docker container
// CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)


func main() {

	fmt.Println("Hello Kyml!")
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	txtKind := updateCommand.String("kind","","Kind of resource to update. (Required)")
	txtName := updateCommand.String("name","","Name of the resource to update. (Required)")
	txtAtt := updateCommand.String("att","","Name of the attribute to update. (Required)")
	txtVal := updateCommand.String("value","","Desired value of the attribute to update. (Required)")
	intIndex := updateCommand.Int("index",0,"In case attribute is in array, use index to specify the array index. (Optional)")

	if len(os.Args) < 2{
		fmt.Print("No command specified.\nCurrently supported commands:\nUpdate - Search for kubernetes entity and update its attributes\n")
		os.Exit(1)
	}
	switch os.Args[1]{
	case "update":
		updateCommand.Parse(os.Args[2:])
		Tail := updateCommand.Args()
		if *txtName == "" || *txtKind ==""{
			updateCommand.PrintDefaults()
			os.Exit(1)
		}
		if len(Tail)!=1{
			updateCommand.PrintDefaults()
			os.Exit(1)
		}
		txtContext:= Tail[0]
		update(txtContext,*txtKind,*txtName,*txtAtt,*txtVal,*intIndex)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func update(txtContext, txtKind,txtName,txtAtt,txtVal string, intIndex int) error {

	err := filepath.Walk(txtContext, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		matched, _ := regexp.MatchString(`\.yaml$`, path)
		if matched{
			yamlFile, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("yamlFile.Get err   #%v ", err)
			}
			if txtKind == "deployment"{
				var deployment Deployment
				err = yaml.Unmarshal([]byte(yamlFile), &deployment)
				if err != nil {
					log.Fatalf("Unmarshal: %v", err)
				}
				if deployment.Meta.Name == txtName{

					if txtAtt == "image"{
						fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n",deployment.Kind,deployment.Meta.Name,deployment.Spec.Template.Spec.Containers[intIndex].Image,txtVal)
						deployment.Spec.Template.Spec.Containers[intIndex].Image = txtVal
						data, err := yaml.Marshal(&deployment)
						if err != nil {
							log.Fatalf("error: %v", err)
						}
						err = ioutil.WriteFile(path, data, 0644)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

			}
			if txtKind == "rollout"{
				var rollout Rollout
				err = yaml.Unmarshal([]byte(yamlFile), &rollout)
				if err != nil {
					log.Fatalf("Unmarshal: %v", err)
				}
				if rollout.Meta.Name == txtName{

					if txtAtt == "image"{
						fmt.Printf("Updating resource of kind: %s\tNamed: %s\tImage:%s ==> %s\n",rollout.Kind,rollout.Meta.Name,rollout.Spec.Template.Spec.Containers[intIndex].Image,txtVal)
						rollout.Spec.Template.Spec.Containers[intIndex].Image = txtVal
						data, err := yaml.Marshal(&rollout)
						if err != nil {
							log.Fatalf("error: %v", err)
						}
						err = ioutil.WriteFile(path, data, 0644)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

			}

		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil

	}

