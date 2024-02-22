import { UserContext, UserContextType } from '@/context/UserContext'
import { Button, Container, Stack, TextField, Typography } from '@mui/material'
import { useContext, useState } from 'react'

const LoginForm = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const { loginUser } = useContext(UserContext) as UserContextType

  const clearForm = () => {
    setEmail('')
    setPassword('')
  }
  const handleLogin = async () => {
    const response = await loginUser({ email, password })
    clearForm()
    console.log(response)
  }
  return (
    <Container maxWidth="sm" sx={{ minWidth: '320px', padding: '12px' }}>
      <Stack spacing={1}>
        <Typography variant="h5" sx={{ textAlign: 'center' }}>
          Login
        </Typography>
        <TextField
          id="standard-basic"
          label="Email"
          value={email}
          onChange={(e) => {
            setEmail(e.target.value)
          }}
          variant="standard"
        />
        <TextField
          id="standard-basic"
          label="Password"
          value={password}
          onChange={(e) => {
            setPassword(e.target.value)
          }}
          variant="standard"
        />
        <Button
          variant="contained"
          onClick={() => {
            handleLogin()
          }}
        >
          Login
        </Button>
      </Stack>
    </Container>
  )
}

export default LoginForm
