import { Button, TextField } from "@mui/material";
import styles from './loginform.module.scss'
const LoginForm = () => {
    return ( 
        <div className={styles['login-form']}>
            <TextField id="standard-basic" label="Email" variant="standard" />
            <TextField id="standard-basic" label="Password" variant="standard" />
            <Button variant="contained">Login</Button>
        </div>
     );
}
 
export default LoginForm;