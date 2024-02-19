import { BugReport } from '@mui/icons-material'
import {
  AppBar,
  Button,
  IconButton,
  Stack,
  Toolbar,
  Typography,
} from '@mui/material'
import { useRouter } from 'next/router'

const NavBar = () => {
  const router = useRouter()
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="logo"
          >
            <BugReport />
          </IconButton>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            BugTracker
          </Typography>
          <Stack direction="row" spacing={2}>
            <Button color="inherit" onClick={() => router.push('/')}>
              Home
            </Button>
            <Button color="inherit" onClick={() => router.push('/dashboard')}>
              Dashboard
            </Button>
            <Button color="inherit" onClick={() => router.push('/login')}>
              Login
            </Button>
          </Stack>
        </Toolbar>
      </AppBar>
    </>
  )
}

export default NavBar
