import { getAllTemplates } from "../../utils/api-calls"
import { APIList, Template } from "../../utils/api-types"
import React, { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import styles from "./home.module.css"
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
    <div className={styles.container}>
      <h1>Templates</h1>
      <div className={styles.templatesContainer}>
        {templates?.results.map((template) => (
          <div className={styles.template} key={template.id}>
            <Link to={`/template/${template.id}`}>
              <strong className={styles.templateName}>{template.name}</strong>
              <div className={styles.templateLanguage}>
                <strong>Language: </strong>
                <span>{template.language}</span>
              </div>
              <p className={styles.description}>{template.description}</p>
            </Link>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Home
