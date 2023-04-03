import { getAllTemplates } from "../../utils/api-calls"
import { APIList, Template } from "../../utils/api-types"
import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import { alpha, Box, styled, Typography } from "@mui/material"
import { grey } from "@mui/material/colors"
const Container = styled("div")(({ theme }) => ({
  padding: theme.spacing(4),
}))

const StyledLink = styled(Link)({
  color: "black",
  textTransform: "uppercase",
  textDecoration: "none",
})

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
      <Typography sx={{ textTransform: "uppercase" }} variant="h6">
        Please select a template first
      </Typography>
      <Box
        sx={{
          display: "grid",
          gridTemplateColumns: `repeat(3, minmax(200px,1fr))`,
          columnGap: 3,
          rowGap: 3,
          py: 4,
        }}
        className="p-10"
      >
        {templates?.results.map((template) => (
          <StyledLink key={template.id} to={`/template/${template.id}`}>
            <Box
              sx={{
                border: `solid 1px ${alpha(grey[800], 0.3)}`,
                borderRadius: 1,
                p: 2,
                textAlign: "center",
                transition: "all 0.2s ease",
                "&:hover": {
                  boxShadow: (theme) =>
                    `0 0 4px ${alpha(theme.palette.common.black, 0.4)}`,
                },
              }}
              className="rounded"
            >
              <strong>{template.name}</strong>
            </Box>
          </StyledLink>
        ))}
      </Box>
    </Container>
  )
}

export default Home
