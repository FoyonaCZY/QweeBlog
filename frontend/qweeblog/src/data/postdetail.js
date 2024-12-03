import axios from 'axios';
import config from './siteconfig';

// 这是一个普通的异步函数，获取分类数据
export default async function GetPostDetail(PostID) {

    let url = `${config.apiDomain}posts/detail/${PostID.ID}`;

    try {
        const response = await axios.get(url); // 请求数据

        return response.data;
    } catch (error) {
        console.error("Error fetching posts:", error.message);
        return []; // 如果出错，返回空数组
    }
}