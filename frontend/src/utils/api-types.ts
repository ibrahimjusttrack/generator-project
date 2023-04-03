export type APIList<T> = {
  results: T[]
  total: number
}

export type Template = {
  id: string
  name: string
  description: string
  language: string
}

export type Field = {
  id: string
  name: string
  type: FieldTypes
}

export enum FieldTypes {
  number = "number",
  string = "string",
  bool = "bool",
  array = "array",
  fromArray = "from_array",
}
