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
  type: FieldTypes
  accessor: string
  default?: string | number | string[] | null
  options?: string[]
  description: string
}

export enum FieldTypes {
  number = "number",
  string = "string",
  bool = "bool",
  array = "array",
  fromArray = "from_array",
}
