import { alpha, Box, Button, CircularProgress } from "@mui/material"
import { grey } from "@mui/material/colors"
import { useCallback, useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { toast } from "react-toastify"
import { generateConfig, getFieldsForTemplate } from "../../utils/api-calls"
import { Field as TypeField } from "../../utils/api-types"
import Field from "./Field/Field"

const TemplatePage = () => {
  const params = useParams<{ id: string }>()
  const [fields, setFields] = useState<TypeField[]>([])
  const [loading, setLoading] = useState<boolean>(false)
  const [values, setValues] = useState<Record<string, any>>({})
  const fetchTemplates = useCallback(async () => {
    if (params?.id) {
      const result = await getFieldsForTemplate(params.id)
      setFields(result.data)
    }
  }, [params.id])

  const onSubmit = async () => {
    setLoading(true)
    if (!params.id) return
    try {
      await generateConfig(params.id, values)
      toast.success("Your service has been created!")
    } catch (error) {
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTemplates()
  }, [fetchTemplates, params])

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
        {loading ? <CircularProgress color="inherit" size={20} /> : "Generate"}
      </Button>
    </Box>
  )
}

export default TemplatePage
