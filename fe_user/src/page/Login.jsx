import Button from "@mui/material/Button";
import AlternateEmailSharpIcon from "@mui/icons-material/AlternateEmailSharp";
function Login() {
  return (
    <div>
      <Button variant="contained" startIcon={<AlternateEmailSharpIcon />}>
        Google
      </Button>
    </div>
  );
}

export default Login;
