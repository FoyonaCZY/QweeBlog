import React from "react";
import { Box, Typography } from "@mui/material";
import config from "../../data/siteconfig";

export default function PostDetail(ID, title, content, avatar, author, authoravatar, category, created_at, updated_at) {

    const isoDate = new Date(created_at);

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
                height: '500px',
                boxShadow: 3,
                overflow: 'hidden',
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'background.paper',
            }}
        >
            <Box
                sx={{
                    height: '60%',
                    flexDirection: 'column',
                    backgroundImage: `url(${avatar})`,
                    backgroundSize: 'cover',
                    backgroundPosition: 'center',
                    alignItems: 'center',
                    display: 'flex',
                    justifyContent: 'center',
                }}
            >
                <Typography variant="h5" component='a' href={`${config.siteDomain}post/${ID}`} gutterBottom sx={
                    {
                        maxWidth: '90%',
                        fontFamily: 'Microsoft YaHei',
                        fontSize: '2.0rem',
                        color: 'white',
                        textShadow: '2px 2px 4px rgba(0, 0, 0, 0.3)',  // 设置阴影
                        textDecoration: 'none',
                        textAlign: 'center',
                        '&:hover': {
                            color: 'lightgray',
                            cursor: 'pointer',
                        }
                    }
                }>
                    {title}
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
                    <img src={authoravatar} alt={author} style={{
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
                        {author}
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
                            '&:hover': {
                                cursor: 'default',
                            },
                        }
                    }>
                        {date}
                    </Typography>
                </Box>
            </Box>

            <Box
                sx={{
                    height: '40%',
                    padding: 2,
                    boxShadow: '0 -50px 100px rgba(0, 0, 0, 0.5)', // 水平偏移 0，垂直偏移 -4px，模糊半径
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Typography variant="body2" color="text.secondary" sx={
                    {
                        height: '60%',
                        maxHeight: '60%',                    // 高度限制
                        width: '95%',
                        maxWidth: '95%',                    // 宽度限制
                        display: '-webkit-box',           // 启用多行截断
                        gap: 1,
                        alignItems: 'center',
                        fontFamily: 'Microsoft YaHei',
                        fontSize: '0.9rem',
                        lineHeight: 1.8,
                        marginTop: 2,
                        whiteSpace: 'normal',             // 允许换行
                        overflow: 'hidden',               // 隐藏超出的文本
                        WebkitBoxOrient: 'vertical',      // 设置为垂直排列
                        WebkitLineClamp: 4,               // 限制显示三行，超出部分显示省略号
                        '&:hover': {
                            cursor: 'default',
                        },
                    }
                }>
                    {content}
                </Typography>
            </Box>
        </Box>
    );
}