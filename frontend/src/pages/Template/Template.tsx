import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { getFieldsForTemplate } from "../../utils/api-calls"
import { Field as TypeField } from "../../utils/api-types"
import Field from "./Field/Field"
const TemplatePage = () => {
  const params = useParams<{ id: string }>()
  const [fields, setFields] = useState<TypeField[]>([])
  const [values, setValues] = useState<Record<string, any>>({})
  const fetchTemplates = async () => {
    if (params?.id) {
      const result = await getFieldsForTemplate(params.id)
      setFields(result.data)
    }
  }
  useEffect(() => {
    fetchTemplates()
  }, [params])

  return (
    <div>
      {fields.map((field) => (
        <Field
          key={field.id}
          {...field}
          value={values[field.accessor]}
          setValue={(value) =>
            setValues((props) => ({ ...props, [field.accessor]: value }))
          }
        />
      ))}
    </div>
  )
}

export default TemplatePage
