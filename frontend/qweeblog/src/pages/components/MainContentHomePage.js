import * as React from 'react';
import { Typography, Box, Chip, Container, CircularProgress } from '@mui/material';
import config from '../../data/siteconfig';
import GetCategories from '../../data/category';

export default function MainContentHomePage() {
    const [selectedCategory, setSelectedCategory] = React.useState(0);
    const [categories, setCategories] = React.useState([]);  // 初始化为一个空数组
    const [loading, setLoading] = React.useState(true);      // 加载状态
    const [error, setError] = React.useState(null);          // 错误状态

    // 获取分类数据
    React.useEffect(() => {
        const fetchCategories = async () => {
            try {
                const data = await GetCategories(); // 异步获取数据
                setCategories(data);  // 设置分类数据
            } catch (err) {
                setError('Failed to load categories');
            } finally {
                setLoading(false);  // 结束加载状态
            }
        };

        fetchCategories();  // 调用获取数据的函数
    }, []);  // 空依赖数组，表示只在组件挂载时调用一次

    // 处理点击某个分类的事件
    function handleCategoryClick(categoryID) {
        setSelectedCategory(categoryID);
    }

    // 处理点击“全部分类”的事件
    function handleAllCategoryClick() {
        setSelectedCategory(0);
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    // 渲染组件
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 4 }}>
            <div>
                <Typography variant="h4" gutterBottom>
                    {config.siteTitle}
                </Typography>
                <Typography>
                    {config.siteDescription}
                </Typography>
            </div>
            <Container maxWidth="lg" sx={{
                display: 'flex',
                gap: 2,
            }}>
                {/* 全部分类按钮 */}
                <Box sx={{
                    display: 'flex', // 使用 flexbox 横向排列
                    overflowX: 'auto', // 启用横向滚动条
                    whiteSpace: 'nowrap', // 防止换行
                    gap: 1, // 设置 Chip 之间的间距
                    padding: 1, // 设置一些内边距
                    maxWidth: '100%', // 最大宽度为容器宽度
                    scrollbarWidth: 'thin', // 设置滚动条宽度（适用于 Firefox）
                }}>
                    <Chip variant='text' label='全部分类' onClick={handleAllCategoryClick} sx={{
                        backgroundColor: 'transparent',
                        fontWeight: 'bold',
                        '&:hover': {
                            backgroundColor: 'lightgray',
                            cursor: 'pointer',
                        }
                    }} />
                </Box>

                {/* 分类列表 */}
                {loading === true ? <CircularProgress /> :
                    <Box sx={{
                        display: 'flex', // 使用 flexbox 横向排列
                        overflowX: 'auto', // 启用横向滚动条
                        whiteSpace: 'nowrap', // 防止换行
                        gap: 1, // 设置 Chip 之间的间距
                        padding: 1, // 设置一些内边距
                        maxWidth: '100%', // 最大宽度为容器宽度
                        scrollbarWidth: 'thin', // 设置滚动条宽度（适用于 Firefox）
                    }}>
                        {categories.map((category) => (
                            <Chip
                                key={category.ID}
                                label={category.Name}
                                variant="text"
                                onClick={() => handleCategoryClick(category.ID)}
                                sx={{
                                    backgroundColor: category.ID === selectedCategory ? 'lightgray' : 'transparent',
                                    fontWeight: 'bold',
                                    '&:hover': {
                                        backgroundColor: 'lightgray',
                                        cursor: 'pointer',
                                    }
                                }}
                            />
                        ))}
                    </Box>}
            </Container>
        </Box>
    );
}
