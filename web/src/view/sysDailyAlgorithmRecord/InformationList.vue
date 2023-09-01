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
                    title: '重要通知',
                    content: '系统将于明天凌晨进行维护，届时将暂时不可用。'
                },
                {
                    id: 2,
                    title: '系统维护',
                    content: '系统将于明天凌晨进行维护，届时将暂时不可用。'
                }
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
  