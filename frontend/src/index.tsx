import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"
import { ThemeProvider, createTheme } from "@mui/material"
import { ToastContainer } from "react-toastify"
import "./index.css"
import "react-toastify/dist/ReactToastify.css"

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
      <ToastContainer position="bottom-center" />
    </ThemeProvider>
  </React.StrictMode>
)
