import { Button, Container, TextField } from '@mui/material'
import styles from './loginform.module.scss'
const LoginForm = () => {
  return (
    <Container maxWidth="sm">
      <div className={styles['login-form']}>
        <TextField id="standard-basic" label="Email" variant="standard" />
        <TextField id="standard-basic" label="Password" variant="standard" />
        <Button variant="contained">Login</Button>
      </div>
    </Container>
  )
}

export default LoginForm
