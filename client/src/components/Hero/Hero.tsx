import { Box, Button, Typography } from '@mui/material'

const Hero = () => {
  return (
    <Box
      sx={{
        display: 'flex',
        justifyContent: 'center',
        flexDirection: 'column',
        alignItems: 'center',
        paddingTop: '64px',
        paddingBottom: '64px',
      }}
      maxWidth="sm"
      gap={4}
    >
      <Typography variant="h1">BugTracker</Typography>
      <Typography variant="h6" style={{ opacity: '80%', textAlign: 'center' }}>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua.
      </Typography>
      <Button style={{ width: '160px' }} variant="contained">
        Pricing
      </Button>
    </Box>
  )
}

export default Hero
