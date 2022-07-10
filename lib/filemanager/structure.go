package filemanager

import (
	"embed"
	"fmt"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
	"os"
	"text/template"
)

//go:embed template/*
var templateFileFS embed.FS

type EntityTemplateVariable struct {
	Name string
}

func WriteEntityByTemplate(entity metadata.StructureEntity, placeholders ...any) {
	if len(entity.Template) > 0 {
		name := placeholders[0].(string)
		variables := EntityTemplateVariable{
			Name: name,
		}
		path := fmt.Sprintf(entity.Pattern, placeholders...)
		tmpl := template.New(entity.Template)

		entityFile, _ := os.OpenFile(path, 1, 0655)
		tmpl, _ = tmpl.ParseFS(templateFileFS, fmt.Sprintf("template/%s", entity.Template))

		tmpl.ExecuteTemplate(entityFile, entity.Template, variables)

		logger.LogSuccess(path, "Wrote by template")
	}
}

func BuildFromStructure(entities []metadata.StructureEntity, placeholders ...any) {
	for _, entity := range entities {
		switch entity.EntityType {
		case metadata.ENTITY_TYPE_DIR:
			CreateDir(entity.Pattern, placeholders...)
		case metadata.ENTITY_TYPE_FILE:
			CreateFile(entity.Pattern, placeholders...)
			WriteEntityByTemplate(entity, placeholders...)
		}
	}
}
