import { UserContext, UserContextType } from '@/context/UserContext'
import {
  Alert,
  Button,
  Container,
  Stack,
  TextField,
  Typography,
} from '@mui/material'
import { useRouter } from 'next/router'
import { FormEvent, useContext, useState } from 'react'

const LoginForm = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(false)
  const [errMessage, setErrMessage] = useState('')
  const { loginUser } = useContext(UserContext) as UserContextType
  const router = useRouter()

  const clearForm = () => {
    setEmail('')
    setPassword('')
  }

  const handleLogin = async (e: FormEvent) => {
    e.preventDefault()
    setError(false)
    const response = await loginUser({ email, password })
    console.log(response)
    clearForm()
    if (response.error !== undefined) {
      setError(true)
      setErrMessage(response.message.toString())
      return
    }
    router.push('/')
  }
  return (
    <Container maxWidth="sm" sx={{ minWidth: '320px', padding: '12px' }}>
      <form
        onSubmit={(e) => {
          handleLogin(e)
        }}
      >
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
            type="password"
            label="Password"
            value={password}
            onChange={(e) => {
              setPassword(e.target.value)
            }}
            variant="standard"
          />
          <Button variant="contained" type="submit">
            Login
          </Button>
          <Alert severity="error" sx={{ display: error ? 'flex' : 'none' }}>
            Error: {errMessage}
          </Alert>
        </Stack>
      </form>
    </Container>
  )
}

export default LoginForm
