import { useUserContext } from '@/context/UserContext'
import { Box, Button, Typography, useMediaQuery, useTheme } from '@mui/material'
import { Variant } from '@mui/material/styles/createTypography'
import { useRouter } from 'next/router'

const Hero = () => {
  const router = useRouter()
  const { isLoggedIn } = useUserContext()
  const theme = useTheme()
  const isSmallScreen = useMediaQuery(theme.breakpoints.down('sm'))
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
      <Typography variant={isSmallScreen ? 'h3' as Variant : 'h1' as Variant}>BugTracker</Typography>
      <Typography
        variant={isSmallScreen ? 'h7' as Variant : 'h6' as Variant } 
        style={{ opacity: '80%', textAlign: 'center' }}
      >
        BugTracker centralizes communication, allowing teams to collaborate
        effortlessly and keep everyone on the same page.
      </Typography>
      <Button
        style={{ width: '160px' }}
        variant="contained"
        onClick={() => {
          isLoggedIn ? router.push('/dashboard') : router.push('/register')
        }}
      >
        Start Tracking
      </Button>
    </Box>
  )
}

export default Hero
