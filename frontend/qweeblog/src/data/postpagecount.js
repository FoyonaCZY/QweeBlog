import axios from "axios";
import config from "./siteconfig";

export default async function GetPageCount(categoryID) {
    let url = categoryID === 0 ? `${config.apiDomain}posts/count` : `${config.apiDomain}posts/countbycategory/${categoryID}`;

    try {
        const response = await axios.get(url); // 请求数据

        // 检查返回数据的结构
        const { count } = response.data;

        return count;
    } catch (error) {
        console.error("Error fetching page counts:", error.message);
        return 0;
    }
}