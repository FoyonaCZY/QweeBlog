import axios from 'axios';
import config from './siteconfig';

// Post 类型结构
class Post {
    constructor(ID, title, content, avatar, author, authoravatar, category, created_at, updated_at) {
        this.ID = ID;
        this.title = title;
        this.content = content;
        this.avatar = avatar;
        this.author = author;
        this.authoravatar = authoravatar;
        this.category = category;
        this.created_at = created_at;
        this.updated_at = updated_at;
    }
}

// 这是一个普通的异步函数，获取分类数据
export default async function GetPostList(categoryID, page) {

    let url = categoryID === 0 ? `${config.apiDomain}posts/list/${page}` : `${config.apiDomain}posts/listbycategory/${categoryID}/${page}`;

    try {
        const response = await axios.get(url); // 请求数据

        // 检查返回数据的结构
        const { posts } = response.data;

        // 如果 posts 存在且是一个数组
        if (Array.isArray(posts)) {
            const postsArray = posts.map(item => {
                return new Post(item.ID, item.title, item.content, item.avatar, item.User.nickname, item.User.avatar, item.category_id, item.CreatedAt, item.UpdatedAt); // 转换为 Post 数组
            });

            return postsArray; // 返回数组
        }

        // 如果 posts 不是数组，返回空数组
        return [];
    } catch (error) {
        console.error("Error fetching posts:", error.message);
        return []; // 如果出错，返回空数组
    }
}