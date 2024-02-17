import { Button, Container, Stack, TextField } from '@mui/material'
const LoginForm = () => {
  return (
    <Container maxWidth="sm">
      <Stack spacing={1}>
        <TextField id="standard-basic" label="Email" variant="standard" />
        <TextField id="standard-basic" label="Password" variant="standard" />
        <Button variant="contained">Login</Button>
      </Stack>
    </Container>
  )
}

export default LoginForm
