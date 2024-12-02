import axios from 'axios';
import config from './siteconfig';

// Category 类型结构
class Category {
    constructor(ID, Name) {
        this.ID = ID;
        this.Name = Name;
    }
}

// 这是一个普通的异步函数，获取分类数据
export default async function GetCategories() {
    try {
        const response = await axios.get(`${config.apiDomain}categories/list`); // 请求数据

        // 检查返回数据的结构
        const { categories } = response.data;

        // 如果 categories 存在且是一个数组
        if (Array.isArray(categories)) {
            const categoriesArray = categories.map(item => {
                return new Category(item.ID, item.name); // 转换为 Category 数组
            });
            
            return categoriesArray; // 返回数组
        }

        // 如果 categories 不是数组，返回空数组
        return [];
    } catch (error) {
        console.error("Error fetching categories:", error.message);
        return []; // 如果出错，返回空数组
    }
}
