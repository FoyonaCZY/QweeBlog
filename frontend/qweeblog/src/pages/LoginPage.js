import * as React from 'react';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import FormLabel from '@mui/material/FormLabel';
import FormControl from '@mui/material/FormControl';
import Link from '@mui/material/Link';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import MuiCard from '@mui/material/Card';
import { styled } from '@mui/material/styles';
import { StyledTextButtonSignIn } from './components/AppBarHomePage';
import axios from 'axios';
import config from '../data/siteconfig';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import { useState } from 'react';

const Card = styled(MuiCard)(({ theme }) => ({
    display: 'flex',
    flexDirection: 'column',
    alignSelf: 'center',
    width: '100%',
    padding: theme.spacing(4),
    gap: theme.spacing(2),
    margin: 'auto',
    [theme.breakpoints.up('sm')]: {
        maxWidth: '450px',
    },
    boxShadow:
        'hsla(220, 30%, 5%, 0.05) 0px 5px 15px 0px, hsla(220, 25%, 10%, 0.05) 0px 15px 35px -5px',
    ...theme.applyStyles('dark', {
        boxShadow:
            'hsla(220, 30%, 5%, 0.5) 0px 5px 15px 0px, hsla(220, 25%, 10%, 0.08) 0px 15px 35px -5px',
    }),
}));

const SignInContainer = styled(Stack)(({ theme }) => ({
    height: 'calc((1 - var(--template-frame-height, 0)) * 100dvh)',
    minHeight: '100%',
    padding: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
        padding: theme.spacing(4),
    },
    '&::before': {
        content: '""',
        display: 'block',
        position: 'absolute',
        zIndex: -1,
        inset: 0,
        backgroundImage:
            'radial-gradient(ellipse at 50% 50%, hsl(210, 100%, 97%), hsl(0, 0%, 100%))',
        backgroundRepeat: 'no-repeat',
        ...theme.applyStyles('dark', {
            backgroundImage:
                'radial-gradient(at 50% 50%, hsla(210, 100%, 16%, 0.5), hsl(220, 30%, 5%))',
        }),
    },
}));

// 登录并获取 JWT Token
const loginUser = async (email, password) => {
    try {
        // 向后端发送 POST 请求进行登录
        const response = await axios.post(`${config.apiDomain}user/login`, {
            email,
            password,
        });

        // 获取返回的 JWT Token
        const { token ,user} = response.data; // 假设返回的数据中包含 token 字段

        // 将 Token 保存到 localStorage 或其他存储中
        localStorage.setItem('jwtToken', token);
        localStorage.setItem('nickname', user.nickname);
        localStorage.setItem('avatar', user.avatar);
        localStorage.setItem('email', user.email);
        localStorage.setItem('groupID', user.group_id);

        return token; // 返回 JWT Token
    } catch (error) {
        console.error('登录失败:', error);
        throw error; // 抛出错误以便处理
    }
};

export default function LoginPage() {
    const [emailError, setEmailError] = React.useState(false);
    const [emailErrorMessage, setEmailErrorMessage] = React.useState('');
    const [passwordError, setPasswordError] = React.useState(false);
    const [passwordErrorMessage, setPasswordErrorMessage] = React.useState('');
    const [openSnackbar, setOpenSnackbar] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('success'); // success or error

    const handleSubmit = async (event) => {
        event.preventDefault(); // 确保阻止表单提交
    
        // 如果存在错误，阻止表单提交
        if (emailError || passwordError) {
            console.log('表单验证错误');
            return;
        }
    
        const data = new FormData(event.currentTarget);
        const email = data.get('email');
        const password = data.get('password');
    
        console.log('尝试登录，邮箱:', email, '密码:', password);
    
        try {
            // 阻塞等待 loginUser 完成
            const token = await loginUser(email, password);
            console.log('登录成功，Token:', token);
    
            // 显示成功通知并跳转到首页
            setSnackbarMessage('登录成功');
            setSnackbarSeverity('success');
            setOpenSnackbar(true);
    
            // 跳转到首页
            window.location.href = '/'; // 或者使用 react-router 的 history.push('/') 重定向到首页
    
        } catch (error) {
            // 显示失败通知，留在当前页面
            console.error('登录失败:', error);
            setSnackbarMessage('登录失败，请检查用户名或密码');
            setSnackbarSeverity('error');
            setOpenSnackbar(true);
        }
    };

    const validateInputs = () => {
        const email = document.getElementById('email');
        const password = document.getElementById('password');

        let isValid = true;

        if (!email.value || !/\S+@\S+\.\S+/.test(email.value)) {
            setEmailError(true);
            setEmailErrorMessage('请输入有效的邮箱地址。');
            isValid = false;
        } else {
            setEmailError(false);
            setEmailErrorMessage('');
        }

        if (!password.value || password.value.length < 8 || password.value.length > 20) {
            setPasswordError(true);
            setPasswordErrorMessage('密码长度必须在 8 到 20 个字符之间。');
            isValid = false;
        } else {
            setPasswordError(false);
            setPasswordErrorMessage('');
        }

        return isValid;
    };

    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
            <CssBaseline enableColorScheme />
            <SignInContainer direction="column" justifyContent="space-between">
                <Card variant="outlined">
                    <Typography
                        component="h1"
                        variant="h4"
                        sx={{ width: '100%', fontSize: 'clamp(2rem, 10vw, 2.15rem)' }}
                    >
                        登录
                    </Typography>
                    <Box
                        component="form"
                        onSubmit={handleSubmit}
                        noValidate
                        sx={{
                            display: 'flex',
                            flexDirection: 'column',
                            width: '100%',
                            gap: 2,
                        }}
                    >
                        <FormControl>
                            <FormLabel htmlFor="email">邮箱</FormLabel>
                            <TextField
                                error={emailError}
                                helperText={emailErrorMessage}
                                id="email"
                                type="email"
                                name="email"
                                placeholder="your@email.com"
                                autoComplete="email"
                                autoFocus
                                required
                                fullWidth
                                variant="outlined"
                                color={emailError ? 'error' : 'primary'}
                            />
                        </FormControl>
                        <FormControl>
                            <FormLabel htmlFor="password">密码</FormLabel>
                            <TextField
                                error={passwordError}
                                helperText={passwordErrorMessage}
                                name="password"
                                placeholder="••••••••"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                                autoFocus
                                required
                                fullWidth
                                variant="outlined"
                                color={passwordError ? 'error' : 'primary'}
                            />
                        </FormControl>
                        <StyledTextButtonSignIn
                            type="submit"
                            onClick={validateInputs}
                        >
                            登录
                        </StyledTextButtonSignIn>
                    </Box>
                    <Divider>or</Divider>
                    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                        <Typography sx={{ textAlign: 'center' }}>
                            还没有账号？{' '}
                            <Link
                                href="/material-ui/getting-started/templates/sign-in/"
                                variant="body2"
                                sx={{ alignSelf: 'center' }}
                            >
                                注册
                            </Link>
                        </Typography>
                    </Box>
                </Card>
            </SignInContainer>
            {/* Snackbar 用于显示通知 */}
            <Snackbar
                open={openSnackbar}
                autoHideDuration={6000}
                onClose={() => setOpenSnackbar(false)}
            >
                <Alert
                    onClose={() => setOpenSnackbar(false)}
                    severity={snackbarSeverity}
                    sx={{ width: '100%' }}
                >
                    {snackbarMessage}
                </Alert>
            </Snackbar>
        </Box>
    );
}