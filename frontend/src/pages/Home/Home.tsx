import { getAllTemplates } from "../../utils/api-calls"
import { APIList, Template } from "../../utils/api-types"
import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import styled from "styled-components"
const Container = styled.div`
  padding: 12px;
`

const Home = () => {
  const [templates, setTemplates] = useState<APIList<Template>>()
  const fetchData = async () => {
    const result = await getAllTemplates()
    setTemplates(result.data)
  }
  useEffect(() => {
    fetchData()
  }, [])
  return (
    <Container>
      <h1>Templates</h1>
      <div className="p-10">
        {templates?.results.map((template) => (
          <div key={template.id} className="rounded">
            <Link to={`/template/${template.id}`}>
              <strong>{template.name}</strong>
              <div>
                <strong>Language: </strong>
                <span>{template.language}</span>
              </div>
              <p>{template.description}</p>
            </Link>
          </div>
        ))}
      </div>
    </Container>
  )
}

export default Home
