<template>
    <div class="container">
        <div class="information-list">
            <ul>
                <li v-for="announcement in announcements" :key="announcement.id" class="announcement-item">
                    <span class="announcement-title">{{ announcement.title }}</span>
                    <p class="announcement-content">{{ announcement.content }}</p>
                </li>
            </ul>
        </div>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            announcements: [
                {
                    id: 1,
                    title: '题目选择',
                    content: '每天解决一个来自 LeetCode、AcWing、Codeforces、牛客网、洛谷或其他在线平台上的题目。或者您也可以选择算协的题目。'
                },
                {
                    id: 2,
                    title: '难度选择',
                    content: '尽量选择一些稍微超出您当前能力范围的题目，这样才能有更好的提高。'
                },
                {
                    id: 3,
                    title: '题目提交',
                    content: '提交完整的AC代码和题目链接，排行榜即可更新。后台会自动记录打卡，供相互学习。'
                },
                {
                    id: 4,
                    title: '持续学习',
                    content: '定期学习算法知识，参加算法训练班、阅读相关书籍和文章，参与编程竞赛等。'
                },
                {
                    id: 5,
                    title: '排名奖励',
                    content: '每月统计一次排名，前x(x为人数/5下取整)名有神秘小奖励。'
                },

            ]
        };
    }
};
</script>
  
<style scoped>
/* 样式可以根据需要调整 */
.information-list {
    padding: 20px;
    background-color: #f5f5f500;
}

.container {
    display: flex;
    justify-content: left;
    align-items: left;
    margin: 2% auto;
}

.section-title {
    font-size: 24px;
    margin-bottom: 10px;
}

.announcement-item {
    background-color: white;
    border: 1px solid #dddddd00;
    padding: 10px;
    margin-bottom: 10px;
}

.announcement-title {
    font-size: 18px;
    font-weight: bold;
}

.announcement-content {
    font-size: 16px;
    color: #666;
}
</style>
  