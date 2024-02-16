import { Button } from '@mui/material'
import { useRouter } from 'next/router'

const NavBar = () => {
  const router = useRouter()
  return (
    <>
      <Button onClick={() => router.push('/')}>Home</Button>
      <Button onClick={() => router.push('/login')}>Login</Button>
    </>
  )
}

export default NavBar
