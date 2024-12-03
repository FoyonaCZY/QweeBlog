import React from "react";
import { Box, Typography, CircularProgress } from "@mui/material";
import GetPostDetail from "../../data/postdetail";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

export default function OnePostDetail(ID) {

    const [loading, setLoading] = React.useState(true);      // 加载状态
    const [error, setError] = React.useState(null);          // 错误状态
    const [post, setPost] = React.useState([]);              // 文章数据

    // 获取文章数据
    React.useEffect(() => {
        const fetchPost = async () => {
            try {
                const data = await GetPostDetail(ID); // 异步获取数据
                setPost(data);  // 设置文章数据
            } catch (err) {
                setError('Failed to load post');
            } finally {
                setLoading(false);  // 结束加载状态
            }
        };

        fetchPost();  // 调用获取数据的函数
    }, [ID]);  // 依赖数组，表示只在 ID 变化时调用

    if (loading) {
        return (
            <Box sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '500px',
            }}>
                <CircularProgress />
            </Box>
        );
    }

    if (error) {
        return (
            <Box sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '500px',
            }}>
                <Typography variant="h4" color="error">
                    {error}
                </Typography>
            </Box>
        );
    }

    const isoDate = new Date(post.CreatedAt);

    const date = isoDate.toLocaleString('zh-CN', {
        weekday: 'long',    // 星期几（例如：星期二）
        year: 'numeric',    // 年份（例如：2024）
        month: 'long',      // 月份（例如：十二月）
        day: 'numeric',     // 日（例如：3）
        hour: '2-digit',    // 小时（例如：00）
        minute: '2-digit',  // 分钟（例如：43）
        second: '2-digit',  // 秒（例如：34）
        hour12: false,      // 使用 24 小时制
    });

    return (
        <Box key={ID}
            sx={{
                borderRadius: '8px',
                height: '100%',
                boxShadow: 3,
                overflow: 'hidden',
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'background.paper',
                alignItems: 'center',
            }}
        >
            <Box
                sx={{
                    height: '350px',
                    width: '100%',
                    flexDirection: 'column',
                    backgroundImage: `url(${post.avatar})`,
                    backgroundSize: 'cover',
                    backgroundPosition: 'center',
                    alignItems: 'center',
                    display: 'flex',
                    justifyContent: 'center',
                }}
            >
                <Typography variant="h5" gutterBottom sx={
                    {
                        fontFamily: 'Microsoft YaHei',
                        fontSize: '2.0rem',
                        color: 'white',
                        textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                    }
                }>
                    {post.title}
                </Typography>

                <Box sx={
                    {
                        display: 'flex',
                        flexDirection: 'row',
                        marginTop: 1,
                        gap: 1,
                        alignItems: 'center',
                        justifyContent: 'center',
                    }
                }>
                    <img src={post.User.avatar} alt={post.User.nickname} style={{
                        width: '30px',
                        height: '30px',
                        borderRadius: '50%',
                        boxShadow: '4px 4px 8px rgba(0, 0, 0, 0.2)', // 水平偏移、垂直偏移、模糊半径、阴影颜色
                    }} />
                    <Typography variant="body2" color="text.secondary" sx={
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
                    }>
                        {post.User.nickname}
                    </Typography>
                    <Typography variant="body2" color="text.secondary" sx={
                        {
                            fontFamily: 'Microsoft YaHei',
                            fontSize: '1.0rem',
                            color: 'white',
                            textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                        }
                    }>
                        |
                    </Typography>
                    <Typography variant="body2" color="text.secondary" sx={
                        {
                            fontFamily: 'Microsoft YaHei',
                            fontSize: '1.0rem',
                            color: 'white',
                            textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                        }
                    }>
                        {date}
                    </Typography>
                </Box>
            </Box>

            <Box sx={{
                height: '100%',
                width: '95%',
                display: 'flex',
                marginTop: 2,
                flexDirection: 'column',
                gap: 2,
                padding: 2,
            }}>
                <ReactMarkdown
                    children={post.content}
                    remarkPlugins={[remarkGfm]}
                    components={{
                        // 自定义Markdown元素的渲染方式
                        h1: ({ node, ...props }) => <Typography variant="h4" gutterBottom sx={{lineHeight:1.8,}} {...props} />,
                        h2: ({ node, ...props }) => <Typography variant="h5" gutterBottom sx={{lineHeight:1.8,}} {...props} />,
                        p: ({ node, ...props }) => <Typography variant="body1" color="text.secondary" sx={{lineHeight:1.8,}} {...props} />,
                        a: ({ node, ...props }) => <Typography variant="body1" color="text.secondary" sx={{lineHeight:1.8,}} {...props} />,
                        // 可以继续为其他Markdown元素添加自定义样式
                    }}
                />
            </Box>

        </Box>
    );
}