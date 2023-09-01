<template>
  <div>
    <!-- 在这里嵌入外部网页 -->
    <object :data="externalUrl" width="100%" height="580"></object>
  </div>
</template>

<script>
export default {
  data() {
    return {
      externalUrl: 'https://docs.qq.com/sheet/DSk9MYlhaZUpJdHZm?tab=BB08J2' // 替换为你想嵌入的外部网页的 URL
    };
  }
};
</script>