package internal

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

type environment struct {
	name string
}

func newEnv(name string) *environment {
	name = strings.ReplaceAll(name, "-", " ")
	name = normalize(name)

	return &environment{
		name: name,
	}
}

func (e *environment) setup() error {
	err := e.createActorFolder()
	if err != nil {
		return err
	}

	err = e.createActorJson()
	if err != nil {
		return err
	}

	err = e.createDockerFile()
	if err != nil {
		return err
	}

	err = e.createActorGoFile()
	if err != nil {
		return err
	}

	err = e.createInputSchema()
	if err != nil {
		return err
	}

	doesGoExist := checkFileExists("go.mod")

	if doesGoExist {
		log.Println("go.mod found, installing latest goapify")

		err = e.installGoApify()
		if err != nil {
			return err
		}
	}

	log.Printf("created %s actor environment\n", e.name)

	return nil
}

func (e *environment) createActorJson() error {
	actorJson := actorJsonTemplate
	actorJson = strings.ReplaceAll(actorJson, "${name}", e.name)

	return os.WriteFile(".actor/actor.json", []byte(actorJson), 0666)
}

func (e *environment) createInputSchema() error {
	inputSchema := inputSchemaTemplate
	inputSchema = strings.ReplaceAll(inputSchema, "${name}", e.name)

	return os.WriteFile(".actor/input_schema.json", []byte(inputSchema), 0666)
}

func (e *environment) createActorGoFile() error {
	actorFile := actorTemplate
	actorFile = strings.ReplaceAll(actorFile, "'", "`")

	return os.WriteFile("actor.go", []byte(actorFile), 0666)
}

func (e *environment) installGoApify() error {
	//github.com/data-harvesters/goapify
	cmd := exec.Command("go get github.com/data-harvesters/goapify@main")

	return cmd.Run()
}

func (e *environment) createActorFolder() error {
	return os.Mkdir(".actor", 0777)
}

func (e *environment) createDockerFile() error {
	dockerFile := dockerFileTemplate
	dockerFile = strings.ReplaceAll(dockerFile, "${name}", e.name)

	return os.WriteFile("Dockerfile", []byte(dockerFile), 0666)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func normalize(input string) string {
	split := strings.Split(input, " ")

	for i, s := range split {
		s = toTitle(s)
		split[i] = s
	}

	return strings.Join(split, " ")
}

func toTitle(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
