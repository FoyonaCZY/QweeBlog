import React from 'react';
import { Container, CssBaseline, Box } from '@mui/material';
import AppAppBar from './components/AppBarHomePage';
import config from '../data/siteconfig';

export default function CommentBoardPage() {

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
                    <AppAppBar />
                </Container>
            </Box>
        </>
    );
}
