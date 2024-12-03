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

const registerUser = async (email, password,nickname) => {
    try {
        const response = await axios.post(`${config.apiDomain}user/register`, {
            email,
            password,
            nickname,
            "receive_email":true
        });
        const { id } = response.data;

        return id;
    } catch (error) {
        console.error('注册失败:', error);
        throw error; // 抛出错误以便处理
    }
};

export default function RegisterPage() {
    const [emailError, setEmailError] = React.useState(false);
    const [emailErrorMessage, setEmailErrorMessage] = React.useState('');
    const [passwordError, setPasswordError] = React.useState(false);
    const [passwordErrorMessage, setPasswordErrorMessage] = React.useState('');
    const [nickError, setNickError] = React.useState(false);
    const [nickErrorMessage, setNickErrorMessage] = React.useState('');
    const [openSnackbar, setOpenSnackbar] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('success'); // success or error

    const handleSubmit = async (event) => {
        event.preventDefault(); // 确保阻止表单提交
    
        // 如果存在错误，阻止表单提交
        if (emailError || passwordError || nickError) {
            return;
        }
    
        const data = new FormData(event.currentTarget);
        const email = data.get('email');
        const password = data.get('password');
        const nickname = data.get('nickname');
    
        try {
            // 阻塞等待 loginUser 完成
            const id = await registerUser(email, password,nickname);
            console.log('注册成功，用户 ID:', id);
    
            // 显示成功通知并跳转到首页
            alert('注册成功，激活邮件已发送至您的邮箱，请查收。');
    
            // 跳转到首页
            window.location.href = '/'; // 或者使用 react-router 的 history.push('/') 重定向到首页
    
        } catch (error) {
            // 显示失败通知，留在当前页面
            console.error('注册失败:', error);
            setSnackbarMessage('注册失败，请检查您的输入。');
            setSnackbarSeverity('error');
            setOpenSnackbar(true);
        }
    };

    const validateInputs = () => {
        const email = document.getElementById('email');
        const password = document.getElementById('password');
        const nickname = document.getElementById('nickname');

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

        if(!nickname.value){
            setNickError(true);
            setNickErrorMessage('请输入昵称');
            isValid = false;
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
                        注册
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
                            <FormLabel htmlFor="nickname">昵称</FormLabel>
                            <TextField
                                error={nickError}
                                helperText={nickErrorMessage}
                                id="nickname"
                                type="nickname"
                                name="nickname"
                                placeholder="请输入您的昵称"
                                autoComplete="nickname"
                                autoFocus
                                required
                                fullWidth
                                variant="outlined"
                                color={emailError ? 'error' : 'primary'}
                            />
                        </FormControl>
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
                            注册
                        </StyledTextButtonSignIn>
                    </Box>
                    <Divider>or</Divider>
                    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                        <Typography sx={{ textAlign: 'center' }}>
                            已经有账号了？{' '}
                            <Link
                                href="/login"
                                variant="body2"
                                sx={{ alignSelf: 'center' }}
                            >
                                登录
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