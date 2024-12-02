import React from 'react';
import { Container, CssBaseline, Box } from '@mui/material';
import AppAppBar from './components/AppBarHomePage';
import MainContentHomePage from './components/MainContentHomePage.js';

function HomePage() {
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 4 }}>
            <CssBaseline enableColorScheme />
            <Container
                maxWidth="lg"
                component="main"
                sx={{ display: 'flex', flexDirection: 'column', my: 16, gap: 4 }}
            >
                <AppAppBar />
                <MainContentHomePage />
            </Container>
        </Box>
    );
}

export default HomePage;