import { Field as FieldProps, FieldTypes } from "../../../utils/api-types"

const Field = (
  props: FieldProps & { value: any; setValue: (value: any) => void }
) => {
  const { setValue } = props
  return (
    <div>
      <label htmlFor={props.id}>{props.accessor}: </label>
      {props.type === FieldTypes.string ? (
        <input
          type={"text"}
          id={props.id}
          onChange={(e) => setValue(e.target.value)}
        />
      ) : props.type === FieldTypes.number ? (
        <input type={"number"} id={props.id} />
      ) : props.type === FieldTypes.fromArray ? (
        <select id={props.id}>
          {props.options?.map((opt) => (
            <option key={opt} value={opt}>
              {opt}
            </option>
          ))}
        </select>
      ) : (
        <input type={"text"} id={props.id} />
      )}
    </div>
  )
}

export default Field
