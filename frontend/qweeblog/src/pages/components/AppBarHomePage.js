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
import { useState } from 'react';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import { ListItemText, ListItemIcon } from '@mui/material';
import LogoutIcon from '@mui/icons-material/Logout';
import PermIdentityIcon from '@mui/icons-material/PermIdentity';

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

    const avatar = localStorage.getItem('avatar');
    const nickname = localStorage.getItem('nickname');

    const [isTokenValid, setIsTokenValid] = useState(null);
    const [anchorEl, setAnchorEl] = useState(null); // 用于控制悬浮菜单

    const goToLogin = () => {
        navigate('/login');
    };

    const goToRegister = () => {
        navigate('/register');
    };

    const goToHome = () => {
        navigate('/');
    };

    const goToCommentBoard = () => {
        navigate('/留言板');
    };

    const handleMenuOpen = (event) => {
        setAnchorEl(event.currentTarget); // 设置菜单的锚点
    };

    const handleMenuClose = () => {
        setAnchorEl(null); // 关闭菜单
    };

    const [open, setOpen] = React.useState(false);

    const toggleDrawer = (newOpen) => () => {
        setOpen(newOpen);
    };

    React.useEffect(() => {
        const token = localStorage.getItem('jwtToken');
        if (!token) {
            setIsTokenValid(false);
        } else {
            setIsTokenValid(true);
        }
    }, []);

    return (
        <>
            <AppBar
                position="absolute"
                enableColorOnDark
                sx={{
                    position: 'fixed',
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
                            {!isTokenValid ? (
                                <>
                                    <StyledTextButton variant="text" size="small" onClick={goToRegister}>
                                        注册
                                    </StyledTextButton>
                                    <StyledTextButtonSignIn variant="contained" size="small" onClick={goToLogin}>
                                        登录
                                    </StyledTextButtonSignIn>
                                </>
                            ) : (
                                <>
                                    <img src={avatar} alt={nickname} style={{
                                        width: '30px',
                                        height: '30px',
                                        borderRadius: '50%',
                                        boxShadow: '4px 4px 8px rgba(0, 0, 0, 0.2)', // 水平偏移、垂直偏移、模糊半径、阴影颜色
                                    }} />
                                    <Typography variant="body2" color="text.secondary"
                                        sx={
                                            {
                                                fontFamily: 'Microsoft YaHei',
                                                fontSize: '1.0rem',
                                                color: 'white',
                                                textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                                                '&:hover': {
                                                    color: 'lightgray',
                                                    cursor: 'pointer',
                                                }
                                            }
                                        }
                                        onClick={handleMenuOpen}>
                                        {nickname}
                                    </Typography>

                                    {/* 这里是悬浮菜单 */}
                                    <Menu
                                        anchorEl={anchorEl}
                                        open={Boolean(anchorEl)}
                                        onClose={handleMenuClose}
                                        sx={{
                                            position: 'absolute', // 确保菜单不会影响布局
                                            top: '10px',          // 根据需要调整菜单位置
                                            left: 0,              // 确保菜单不影响其他组件
                                            zIndex: 9999,         // 确保菜单在其他组件上面
                                        }}
                                    >
                                        <MenuItem
                                            autoFocus={false}
                                            onClick={() => { handleMenuClose(); }}
                                            sx={{
                                                '&.Mui-selected': {
                                                    backgroundColor: 'transparent', // 禁止选中的背景颜色
                                                },
                                                '&.Mui-focusVisible': {
                                                    backgroundColor: 'transparent', // 禁止聚焦时的背景颜色
                                                },
                                            }}
                                        >
                                            <ListItemIcon>
                                                <PermIdentityIcon />
                                            </ListItemIcon>
                                            <ListItemText>个人资料</ListItemText>
                                        </MenuItem>
                                        <MenuItem
                                            autoFocus={false}
                                            onClick={() => {
                                                localStorage.removeItem('jwtToken');
                                                localStorage.removeItem('nickname');
                                                localStorage.removeItem('avatar');
                                                setIsTokenValid(false);
                                                handleMenuClose();
                                            }}
                                            sx={{
                                                '&.Mui-selected': {
                                                    backgroundColor: 'transparent', // 禁止选中的背景颜色
                                                },
                                                '&.Mui-focusVisible': {
                                                    backgroundColor: 'transparent', // 禁止聚焦时的背景颜色
                                                },
                                            }}>
                                            <ListItemIcon>
                                                <LogoutIcon />
                                            </ListItemIcon>
                                            <ListItemText>退出登录</ListItemText>
                                        </MenuItem>
                                    </Menu>
                                </>
                            )}
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
                                    {!isTokenValid ? (
                                        <>
                                            <MenuItem>
                                                <StyledTextButtonSignUp variant="text" size="small" fullWidth onClick={goToRegister}>
                                                    注册
                                                </StyledTextButtonSignUp>
                                            </MenuItem>
                                            <MenuItem>
                                                <StyledTextButtonSignIn variant="outlined" fullWidth onClick={goToLogin}>
                                                    登录
                                                </StyledTextButtonSignIn>
                                            </MenuItem>
                                        </>
                                    ) : null}
                                </Box>
                            </Drawer>
                        </Box>
                    </StyledToolbar>
                </Container>
            </AppBar>
        </>
    );
}

export { StyledTextButtonSignIn };