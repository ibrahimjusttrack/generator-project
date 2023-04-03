import { alpha, Box, Button } from "@mui/material"
import { grey } from "@mui/material/colors"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { generateConfig, getFieldsForTemplate } from "../../utils/api-calls"
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

  const onSubmit = async () => {
    if (!params.id) return
    await generateConfig(params.id, values)
    alert("Successfully created!")
  }

  useEffect(() => {
    fetchTemplates()
  }, [params])

  return (
    <Box
      sx={{
        width: "100%",
        maxWidth: "720px",
        margin: "20px auto",
        border: `solid 1px ${alpha(grey[800], 0.3)}`,
        borderRadius: 1,
        p: 2,
        textAlign: "center",
      }}
    >
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
      <Button variant="contained" onClick={() => onSubmit()}>
        Generate
      </Button>
    </Box>
  )
}

export default TemplatePage
