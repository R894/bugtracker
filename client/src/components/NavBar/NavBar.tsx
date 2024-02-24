import { UserContext, UserContextType } from '@/context/UserContext'
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
import { useContext } from 'react'

const NavBar = () => {
  const router = useRouter()
  const { logoutUser, isLoggedIn } = useContext(UserContext) as UserContextType
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
            {isLoggedIn ? (
              <Button color="inherit" onClick={() => logoutUser()}>
                Logout
              </Button>
            ) : (
              <Button color="inherit" onClick={() => router.push('/login')}>
                Login
              </Button>
            )}
          </Stack>
        </Toolbar>
      </AppBar>
    </>
  )
}

export default NavBar
