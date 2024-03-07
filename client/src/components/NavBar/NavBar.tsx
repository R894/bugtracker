import { UserContext, UserContextType } from '@/context/UserContext'
import { BugReport } from '@mui/icons-material'
import { useTheme } from '@mui/material/styles'
import {
  AppBar,
  Box,
  Button,
  Drawer,
  IconButton,
  Stack,
  Toolbar,
  Typography,
  useMediaQuery,
} from '@mui/material'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useContext, useState } from 'react'
import MenuIcon from '@mui/icons-material/Menu'

const NavBar = () => {
  const router = useRouter()
  const theme = useTheme()
  const { logoutUser, isLoggedIn } = useContext(UserContext) as UserContextType
  const matches = useMediaQuery(theme.breakpoints.up('sm'))
  const [open, setOpen] = useState(false)

  const toggleDrawer = (newOpen: boolean) => () => {
    setOpen(newOpen)
  }

  const HeaderOptions = ({ direction }: { direction: 'row' | 'column' }) => {
    return (
      <Stack direction={direction} spacing={2}>
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
    )
  }

  const DrawerButton = () => {
    return (
      <div>
        <IconButton color="inherit" onClick={toggleDrawer(true)}>
          <MenuIcon />
        </IconButton>
        <Drawer open={open} onClose={toggleDrawer(false)}>
          <HeaderOptions direction="column" />
        </Drawer>
      </div>
    )
  }

  return (
    <Box>
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
            <Link href="/">BugTracker</Link>
          </Typography>
          {matches ? <HeaderOptions direction="row" /> : <DrawerButton />}
        </Toolbar>
      </AppBar>
    </Box>
  )
}

export default NavBar
