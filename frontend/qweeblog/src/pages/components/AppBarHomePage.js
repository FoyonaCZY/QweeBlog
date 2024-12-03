import * as React from 'react';
import { alpha, styled } from '@mui/material/styles';
import Box from '@mui/material/Box';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Container from '@mui/material/Container';
import Divider from '@mui/material/Divider';
import MenuItem from '@mui/material/MenuItem';
import Drawer from '@mui/material/Drawer';
import MenuIcon from '@mui/icons-material/Menu';
import CloseRoundedIcon from '@mui/icons-material/CloseRounded';
import { useNavigate } from 'react-router-dom';

const StyledToolbar = styled(Toolbar)(({ theme }) => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    flexShrink: 0,
    borderRadius: `18px`,
    backdropFilter: 'blur(40px)',
    border: '1px solid',
    borderColor: (theme.vars || theme).palette.divider,
    backgroundColor: alpha(theme.palette.background.default, 0.05),
    boxShadow: (theme.vars || theme).shadows[5],
    padding: '10px 20px',
}));

const StyledTextButton = styled(Button)(({ theme }) => ({
    fontFamily: 'Microsoft YaHei', // 设置字体
    fontWeight: 'bold',               // 设置字体加粗
    fontSize: '0.8rem',               // 设置字体大小
    color: 'white',                    // 设置字体颜色
    textTransform: 'none',            // 防止自动转换为大写
    ':hover': {
        color: 'darkgray',
    },
}));

const StyledTextButtonSignUp = styled(Button)(({ theme }) => ({
    fontFamily: 'Microsoft YaHei', // 设置字体
    fontWeight: 'bold',               // 设置字体加粗
    fontSize: '0.8rem',               // 设置字体大小
    color: 'black',                    // 设置字体颜色
    textTransform: 'none',            // 防止自动转换为大写
    ':hover': {
        color: 'darkgray',
    },
}));

const StyledTextButtonSignIn = styled(Button)(({ theme }) => ({
    fontFamily: 'Microsoft YaHei', // 设置字体
    fontWeight: 'bold',               // 设置字体加粗
    fontSize: '0.8rem',               // 设置字体大小
    color: 'white',                    // 设置字体颜色
    textTransform: 'none',            // 防止自动转换为大写
    borderRadius: '10px',
    backgroundColor: '#333333',
    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
    ':hover': {
        color: 'darkgray',
    },
}));

export default function AppAppBar() {
    const navigate = useNavigate();

    const goToHome = () => {
        navigate('/');
    };

    const goToCommentBoard = () => {
        navigate('/留言板');
    };

    const [open, setOpen] = React.useState(false);

    const toggleDrawer = (newOpen) => () => {
        setOpen(newOpen);
    };

    return (
        <AppBar
            position="fixed"
            enableColorOnDark
            sx={{
                boxShadow: 0,
                bgcolor: 'transparent',
                backgroundImage: 'none',
                mt: 'calc(var(--template-frame-height, 0px) + 28px)',
            }}
        >
            <Container maxWidth="lg">
                <StyledToolbar variant="dense" disableGutters>
                    <Box sx={{ flexGrow: 1, display: 'flex', alignItems: 'center', px: 0 }}>
                        <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
                            <StyledTextButton variant="text" size="small" onClick={goToHome}>
                                首页
                            </StyledTextButton>
                            <StyledTextButton variant="text" size="small" onClick={goToCommentBoard}>
                                留言板
                            </StyledTextButton>
                            <StyledTextButton variant="text" size="small">
                                归档
                            </StyledTextButton>
                            <StyledTextButton variant="text" size="small">
                                朋友们
                            </StyledTextButton>
                        </Box>
                    </Box>
                    <Box
                        sx={{
                            display: { xs: 'none', md: 'flex' },
                            gap: 1,
                            alignItems: 'center',
                        }}
                    >
                        <StyledTextButton variant="text" size="small">
                            注册
                        </StyledTextButton>
                        <StyledTextButtonSignIn variant="contained" size="small">
                            登录
                        </StyledTextButtonSignIn>
                    </Box>
                    <Box sx={{ display: { xs: 'flex', md: 'none' }, gap: 1 }}>
                        <IconButton aria-label="Menu button" onClick={toggleDrawer(true)}>
                            <MenuIcon />
                        </IconButton>
                        <Drawer
                            anchor="top"
                            open={open}
                            onClose={toggleDrawer(false)}
                            PaperProps={{
                                sx: {
                                    top: 'var(--template-frame-height, 0px)',
                                },
                            }}
                        >
                            <Box sx={{ p: 2, backgroundColor: 'background.default' }}>
                                <Box
                                    sx={{
                                        display: 'flex',
                                        justifyContent: 'flex-end',
                                    }}
                                >
                                    <IconButton onClick={toggleDrawer(false)}>
                                        <CloseRoundedIcon />
                                    </IconButton>
                                </Box>
                                <MenuItem onClick={goToHome}>首页</MenuItem>
                                <MenuItem onClick={goToCommentBoard}>留言板</MenuItem>
                                <MenuItem>归档</MenuItem>
                                <MenuItem>朋友们</MenuItem>
                                <Divider sx={{ my: 3 }} />
                                <MenuItem>
                                    <StyledTextButtonSignUp variant="text" size="small" fullWidth>
                                        注册
                                    </StyledTextButtonSignUp>
                                </MenuItem>
                                <MenuItem>
                                    <StyledTextButtonSignIn variant="outlined" fullWidth>
                                        登录
                                    </StyledTextButtonSignIn>
                                </MenuItem>
                            </Box>
                        </Drawer>
                    </Box>
                </StyledToolbar>
            </Container>
        </AppBar>
    );
}