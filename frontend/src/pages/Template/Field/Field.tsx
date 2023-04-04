import { Box, Select, styled, TextField, Tooltip } from "@mui/material"
import { Help } from "@mui/icons-material"
import { Field as FieldProps, FieldTypes } from "../../../utils/api-types"

const Label = styled("label")({
  textTransform: "capitalize",
  fontWeight: "600",
  marginRight: 12,
})
const HelpIcon = styled(Help)(({ theme }) => ({
  height: "0.75em",
  verticalAlign: "middle",
  lineHeight: 1,
}))

const Field = (
  props: FieldProps & { value: any; setValue: (value: any) => void }
) => {
  const { setValue } = props
  return (
    <Box sx={{ display: "flex", alignItems: "center", mb: 2 }}>
      <Label htmlFor={props.id}>{props.accessor}: </Label>
      {props.type === FieldTypes.string ? (
        <TextField
          type={"text"}
          id={props.id}
          onChange={(e) => setValue(e.target.value)}
          sx={{ flex: 1 }}
          variant="outlined"
          size="small"
        />
      ) : props.type === FieldTypes.number ? (
        <TextField
          type={"number"}
          id={props.id}
          onChange={(e) => setValue(e.target.value)}
          sx={{ flex: 1 }}
          variant="outlined"
          size="small"
        />
      ) : props.type === FieldTypes.fromArray ? (
        <Select
          id={props.id}
          onChange={(e) => setValue(e.target.value)}
          sx={{ flex: 1 }}
          variant="outlined"
          size="small"
        >
          {props.options?.map((opt) => (
            <option key={opt} value={opt}>
              {opt}
            </option>
          ))}
        </Select>
      ) : (
        <input type={"text"} id={props.id} />
      )}
      <Tooltip title={props.description}>
        <HelpIcon />
      </Tooltip>
    </Box>
  )
}

export default Field
