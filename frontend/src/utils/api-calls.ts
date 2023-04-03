import { APIList, Template } from "@/utils/api-types"

export const templates: Template[] = [
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "if of template",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "if of template",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "if of template",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "if of template",
  },
]

export const getAllTemplates = (): APIList<Template> => ({
  results: templates,
  total: templates.length,
})

export const getTemplateById = (id: string): Template | undefined =>
  templates.find((template) => template.id === id)
