package internal

import (
	"os"
	"strings"
)

type environment struct {
	name      string
	actorName string
}

func newEnv(actorName, name string) *environment {
	name = strings.ReplaceAll(name, "-", " ")

	return &environment{
		name:      name,
		actorName: actorName,
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

func (e *environment) createActorFolder() error {
	return os.Mkdir(".actor", 0777)
}

func (e *environment) createDockerFile() error {
	dockerFile := dockerFileTemplate
	dockerFile = strings.ReplaceAll(dockerFile, "${name}", e.actorName)

	return os.WriteFile("Dockerfile", []byte(dockerFile), 0666)
}
