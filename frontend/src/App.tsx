import { BrowserRouter, Route, Routes } from "react-router-dom"
import Home from "./pages/Home/Home"
import Template from "./pages/Template/Template"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/template/:id" element={<Template />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
