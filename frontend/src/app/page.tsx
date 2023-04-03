"use client" // this is a client component 👈🏽

import { getAllTemplates } from "@/utils/api-calls"
import { APIList, Template } from "@/utils/api-types"
import React, { useEffect, useState } from "react"
import styles from "./page.module.css"
const Page = () => {
  const [templates, setTemplates] = useState<APIList<Template>>()

  useEffect(() => {
    setTemplates(getAllTemplates())
  }, [])
  return (
    <div className={styles.container}>
      <h1>Templates</h1>
      <div className={styles.templatesContainer}>
        {templates?.results.map((template) => (
          <div className={styles.template} key={template.id}>
            <strong className={styles.templateName}>{template.name}</strong>
            <div className={styles.templateLanguage}>
              <strong>Language: </strong>
              <span>{template.language}</span>
            </div>
            <p className={styles.description}>{template.description}</p>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Page
