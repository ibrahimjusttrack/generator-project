import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"
import "./index.css"
import { ThemeProvider, createTheme } from "@mui/material"
const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement)
root.render(
  <React.StrictMode>
    <ThemeProvider
      theme={createTheme({
        typography: {
          fontFamily: "Nunito, sans-serif",
        },
      })}
    >
      <App />
    </ThemeProvider>
  </React.StrictMode>
)
