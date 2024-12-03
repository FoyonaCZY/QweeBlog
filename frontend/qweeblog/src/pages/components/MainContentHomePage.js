import * as React from 'react';
import { Typography, Box, Chip, Container, CircularProgress } from '@mui/material';
import config from '../../data/siteconfig';
import GetCategories from '../../data/category';
import GetPageCount from '../../data/postpagecount';
import GetPostList from '../../data/postlist';
import PostDetail from './postdetail';

export default function MainContentHomePage() {
    const [selectedCategory, setSelectedCategory] = React.useState(0);
    const [currentPage, setPage] = React.useState(1);
    const [categories, setCategories] = React.useState([]);  // 初始化为一个空数组
    const [loading, setLoading] = React.useState(true);      // 加载状态
    const [error, setError] = React.useState(null);          // 错误状态
    const [pageCount, setPageCount] = React.useState(1);     // 总页数
    const [posts, setPosts] = React.useState([]);            // 文章列表

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

    //获取总页数
    React.useEffect(() => {
        const fetchPageCount = async () => {
            try {
                const data = await GetPageCount(selectedCategory); // 异步获取数据
                setPageCount(data === 0 ? 1 : data);  // 设置分类数据
            } catch (err) {
                setError('Failed to load categories');
            } finally {
                setLoading(false);  // 结束加载状态
            }
        };

        fetchPageCount();  // 调用获取数据的函数
    }, [selectedCategory]);

    // 获取文章列表
    React.useEffect(() => {
        const fetchPosts = async () => {
            try {
                const data = await GetPostList(selectedCategory, currentPage); // 异步获取数据
                setPosts(data);  // 设置文章列表
            } catch (err) {
                setError('Failed to load posts');
            } finally {
                setLoading(false);  // 结束加载状态
            }
        };

        fetchPosts();  // 调用获取数据的函数
    }, [selectedCategory, currentPage]);  // 依赖于 selectedCategory 和 currentPage

    // 处理点击某个分类的事件
    function handleCategoryClick(categoryID) {
        setPage(1);
        setSelectedCategory(categoryID);
    }

    // 处理点击“全部分类”的事件
    function handleAllCategoryClick() {
        setPage(1);
        setSelectedCategory(0);
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    // 渲染组件
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 4}}>

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
                        '&:hover': {
                            cursor: 'default',
                        }
                    }
                }>
                    {config.siteTitle}
                </Typography>
                {/* <Typography>
                    {config.siteDescription}
                </Typography> */}
            </div>

            {/* 分类列表 */}
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
                        textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                        color: 'white',
                        fontSize: '1.0rem',
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
                                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                                    color: 'white',
                                    fontSize: '1.0rem',
                                    '&:hover': {
                                        backgroundColor: 'lightgray',
                                        cursor: 'pointer',
                                    }
                                }}
                            />
                        ))}
                    </Box>}
            </Container>

            {/* 文章列表 */}
            <Container maxWidth="lg" sx={{
                display: 'flex',
                flexDirection: 'column',
                gap: 5,
            }}>
                {loading === true ? <CircularProgress /> :
                    posts.map((post) => (
                        PostDetail(post.ID, post.title, post.content, post.avatar, post.author, post.authoravatar, post.category, post.created_at, post.updated_at)
                    ))}
            </Container>

            {/*换页器*/}
            <Container maxWidth="lg" sx={{
                display: 'flex',
                gap: 2,
                alighItems: 'center',
                justifyContent: 'center',
            }}>
                {/* 首页按钮 */}
                {currentPage === 1 ? null : <Chip variant='text' label='首页' onClick={() => setPage(1)} sx={{
                    backgroundColor: 'transparent',
                    color: 'white',
                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    fontWeight: 'bold',
                    '&:hover': {
                        backgroundColor: 'lightgray',
                        cursor: 'pointer',
                    }
                }} />}

                {/* 当前页码-1 */}
                {currentPage === 1 ? null : <Chip variant='text' label={currentPage - 1} onClick={() => setPage(currentPage - 1)} sx={{
                    backgroundColor: 'transparent',
                    fontWeight: 'bold',
                    color: 'white',
                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    '&:hover': {
                        backgroundColor: 'lightgray',
                        cursor: 'pointer',
                    }
                }} />}

                {/* 当前页码 */}
                <Chip variant='text' label={currentPage} sx={{
                    backgroundColor: 'lightgray',
                    fontWeight: 'bold',
                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    '&:hover': {
                        backgroundColor: 'lightgray',
                        cursor: 'pointer',
                    }
                }} />

                {/* 当前页码+1 */}
                {currentPage === pageCount ? null : <Chip variant='text' label={currentPage + 1} onClick={() => setPage(currentPage + 1)} sx={{
                    backgroundColor: 'transparent',
                    fontWeight: 'bold',
                    color: 'white',
                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    '&:hover': {
                        backgroundColor: 'lightgray',
                        cursor: 'pointer',
                    }
                }} />}

                {/* 尾页按钮 */}
                {currentPage === pageCount ? null : <Chip variant='text' label='尾页' onClick={() => setPage(pageCount)} sx={{
                    backgroundColor: 'transparent',
                    fontWeight: 'bold',
                    color: 'white',
                    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    '&:hover': {
                        backgroundColor: 'lightgray',
                        cursor: 'pointer',
                    }
                }} />}
            </Container>
        </Box>
    );
}
