import { Button, Container, Stack, TextField, Typography } from '@mui/material'
import { useState } from 'react'
const RegisterForm = () => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const clearForm = () => {
    setUsername('')
    setEmail('')
    setPassword('')
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
            clearForm()
          }}
        >
          Register
        </Button>
      </Stack>
    </Container>
  )
}

export default RegisterForm
