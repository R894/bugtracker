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
import { useContext, useState } from 'react'
const RegisterForm = () => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState(false)
  const [errMessage, setErrMessage] = useState('')
  const { registerUser } = useContext(UserContext) as UserContextType
  const router = useRouter()

  const clearForm = () => {
    setUsername('')
    setEmail('')
    setPassword('')
  }

  const handleRegister = async () => {
    setError(false)
    const response = await registerUser({ email, username, password })
    clearForm()
    if (response.error !== undefined) {
      setError(true)
      setErrMessage(JSON.stringify(response.message))
      return
    }
    router.push('/')
  }

  return (
    <Container maxWidth="sm" sx={{ minWidth: '320px', padding: '12px' }}>
      <Stack spacing={1}>
        <Typography variant="h5" sx={{ textAlign: 'center' }}>
          Register
        </Typography>
        <TextField
          id="standard-basic"
          label="Username"
          variant="standard"
          onChange={(e) => {
            setUsername(e.target.value)
          }}
          value={username}
        />
        <TextField
          id="standard-basic"
          label="Email"
          variant="standard"
          onChange={(e) => {
            setEmail(e.target.value)
          }}
          value={email}
        />
        <TextField
          id="standard-basic"
          type="password"
          label="Password"
          variant="standard"
          onChange={(e) => {
            setPassword(e.target.value)
          }}
          value={password}
        />
        <Button
          variant="contained"
          onClick={() => {
            handleRegister()
          }}
        >
          Register
        </Button>
        <Alert severity="error" sx={{ display: error ? 'flex' : 'none' }}>
          Error: {errMessage}
        </Alert>
      </Stack>
    </Container>
  )
}

export default RegisterForm
