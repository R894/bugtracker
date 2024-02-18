import { Box, Container, Typography } from '@mui/material'

const Footer = () => {
  return (
    <Box sx={{ paddingBottom: '4px' }}>
      <Container maxWidth="sm">
        <Typography>Copyright © BugTracker</Typography>
      </Container>
    </Box>
  )
}

export default Footer
