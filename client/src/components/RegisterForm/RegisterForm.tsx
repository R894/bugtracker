import { Button, Container, Stack, TextField, Typography } from '@mui/material'
const RegisterForm = () => {
  return (
    <Container maxWidth="sm" sx={{minWidth:'320px', padding:'12px'}}>
      <Stack spacing={1}>
        <Typography variant='h5' sx={{textAlign:'center'}}>Register</Typography>
        <TextField id="standard-basic" label="Username" variant="standard" />
        <TextField id="standard-basic" label="Email" variant="standard" />
        <TextField id="standard-basic" label="Password" variant="standard" />
        <Button variant="contained">Register</Button>
      </Stack>
    </Container>
  )
}

export default RegisterForm
