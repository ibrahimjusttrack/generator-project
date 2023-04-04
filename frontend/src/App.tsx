import { styled, Typography } from "@mui/material"
import { BrowserRouter, Link, Route, Routes } from "react-router-dom"
import Home from "./pages/Home/Home"
import Template from "./pages/Template/Template"
import logo from "./assets/imgs/logo.svg"
import { grey } from "@mui/material/colors"
import AddTemplate from "./pages/AddTemplate/AddTemplate"
const Navbar = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(3, 5),
  background: grey[50],
}))

function App() {
  return (
    <BrowserRouter>
      <Navbar>
        <Link to={"/"}>
          <img src={logo} alt="justtrack logo" />
        </Link>
        <Typography
          variant="h5"
          sx={{ flex: 1, textAlign: "center", fontWeight: "bold" }}
        >
          Welcome to the generator
        </Typography>
      </Navbar>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/template/add" element={<AddTemplate />} />
        <Route path="/template/:id" element={<Template />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
