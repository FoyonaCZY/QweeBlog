import React, { useEffect } from 'react';
import { Container, CssBaseline, Box, Typography } from '@mui/material';
import config from '../data/siteconfig';
import { useLocation } from 'react-router-dom';
import axios from 'axios';

export default function ActivatePage() {

    const location = useLocation();

    const id = new URLSearchParams(location.search).get('id');
    const token = new URLSearchParams(location.search).get('token');

    useEffect(() => {
        const activateUser = async () => {
            let url = `${config.apiDomain}activate?id=${id}&token=${token}`;

            const response = await axios.get(url); // 请求数据

            console.log(response.data); // 输出返回数据
        };

        activateUser();
    }, [id, token]);

    return (
        <>
            {/* 背景图的设置 */}
            <Box
                sx={{
                    display: 'flex',
                    backgroundImage: `url(${config.backgroundImage})`, // 从 config 获取背景图 URL
                    backgroundSize: 'cover', // 背景图覆盖整个区域
                    backgroundPosition: 'center', // 背景图居中
                    backgroundAttachment: 'fixed', // 背景固定，不随页面滚动
                    position: 'fixed', // 背景图固定
                    top: 0,
                    left: 0,
                    right: 0,
                    bottom: 0, // 保证背景图覆盖整个视口
                    zIndex: -1, // 背景图在内容后面
                }}
            />

            {/* 内容区域 */}
            <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
                <CssBaseline enableColorScheme />
                <Container
                    maxWidth="lg"
                    component="main"
                    sx={{ display: 'flex', flexDirection: 'column', my: 16, gap: 4, position: 'relative', zIndex: 1 }}
                >
                    <Typography variant="h4" sx={
                        {
                            fontFamily: 'Microsoft YaHei',
                            fontWeight: 'bold',
                            fontSize: '3.0rem',
                            color: 'white',
                            textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                            alignItems: 'center',
                            display: 'flex',
                            justifyContent: 'center',
                            '&:hover': {
                                cursor: 'default',
                            },
                        }
                    }>
                        激活成功，请返回首页重新登录
                    </Typography>
                </Container>
            </Box>
        </>
    );
}
