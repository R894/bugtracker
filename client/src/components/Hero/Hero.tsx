import { Box, Button, Typography } from '@mui/material'
import { useRouter } from 'next/router'

const Hero = () => {
  const router = useRouter()
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
        BugTracker centralizes communication, allowing teams to collaborate
        effortlessly and keep everyone on the same page.
      </Typography>
      <Button
        style={{ width: '160px' }}
        variant="contained"
        onClick={() => {
          router.push('/register')
        }}
      >
        Start Tracking
      </Button>
    </Box>
  )
}

export default Hero
