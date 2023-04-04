import { APIList, Field, Template } from "./api-types"
import axios, { AxiosResponse } from "axios"
const api = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
})

export const templates: Template[] = [
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "1",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "2",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "3",
  },
  {
    name: "Go Microservices",
    description: "spetial microservices",
    language: "go",
    id: "4",
  },
]

export const getAllTemplates = async (): Promise<
  AxiosResponse<APIList<Template>>
> => api.get("/templates/all")

export const getTemplateById = (id: string): Template | undefined =>
  templates.find((template) => template.id === id)

export const getFieldsForTemplate = async (
  id: string
): Promise<AxiosResponse<Field[]>> => api.get(`/fields/${id}`)

export const generateConfig = async (
  id: string,
  inputs: Record<string, any>
): Promise<AxiosResponse> =>
  api.post(
    "/config/" + id,
    Object.keys(inputs).map((key) => ({
      accessor: key,
      value: inputs[key],
    }))
  )
