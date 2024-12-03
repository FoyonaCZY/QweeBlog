import React from 'react';
import { Container, CssBaseline, Box, Typography } from '@mui/material';
import config from '../data/siteconfig';
import AppAppBar from './components/AppBarHomePage';
import { useParams } from 'react-router-dom';
import OnePostDetail from './components/onePostDetail';

export default function PostDetailPage() {

    const {id} = useParams();
    console.log(id);

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
            <Box sx={{
                display: 'flex',
                flexDirection: 'column',
                minHeight: '100vh',
            }}>
                <CssBaseline enableColorScheme />
                <Container
                    maxWidth="lg"
                    component="main"
                    sx={{ display: 'flex', flexDirection: 'column', my: 16, gap: 4, position: 'relative', zIndex: 1 }}
                >
                    <AppAppBar />
                    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 4 }}>

                        {/* 网站标题和描述 */}
                        <div>
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
                                }
                            }>
                                {config.siteTitle}
                            </Typography>
                        </div>
                        
                        <OnePostDetail ID={id}/>
                    </Box>
                </Container>
            </Box>
        </>
    );
}
