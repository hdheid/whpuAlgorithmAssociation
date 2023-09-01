<template>
  <div class="page-container">
    <div class="left-section">
      <h2 class="section-title">你好，来提交每日一题吧！</h2>
      <form @submit.prevent="submitData" class="data-form">
        <div class="form-item">
          <label for="code" class="form-label" >提交代码：</label>
          <div class="textarea-wrapper">
            <textarea id="code" v-model="code" class="textarea" :placeholder="this.codeText"></textarea>
          </div>
        </div>

        <div class="form-item">
          <label for="link" class="form-label">链接：</label>
          <input type="text" id="link" v-model="link" class="input" placeholder="url链接">
        </div>

        <button type="submit" class="submit-button" v-on:click="showModal = true">提交</button>
        <!-- 弹窗内容 -->
        <div v-if="showModal" class="modal-wrapper">
          <div class="modal">
            <sysDailyAlgorithmRecordForm :link="link" :code = "code" @change-modal="handleBack"/>
          </div>
        </div>
      </form>
    </div>

    <div class="right-section">
      <h2 class="section-title">打卡规则</h2>
      <!-- <calendar /> -->
      <InformationList />
    </div>

  </div>
</template>

<script>
import calendar from "./calendar.vue";
import sysDailyAlgorithmRecordForm from "./sysDailyAlgorithmRecordForm.vue";
import InformationList from "./InformationList.vue";

// ...
import {
  createDailyAlgorithmRecord,
  deleteDailyAlgorithmRecord,
  deleteDailyAlgorithmRecordByIds,
  updateDailyAlgorithmRecord,
  findDailyAlgorithmRecord,
  getDailyAlgorithmRecordList
} from '@/api/sysDailyAlgorithmRecord'
// 全量引入格式化工具 请按需保留
import { setSelfInfo, changePassword } from '@/api/user.js'
import { reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import SelectImage from '@/components/selectImage/selectImage.vue'

export default {
  name: 'sysDailyAlgorithmRecord',
  // ...
  components: {
    calendar,
    sysDailyAlgorithmRecordForm,
    InformationList
    // ...
  },

  data(){
    return {
      showModal: false,
      link: '',
      code: '',
      codeText: "```\n语言类型（例如cpp）\n你的代码\n```",
    };
  },
  methods: {
    handleBack() {
      this.showModal = false;
    }
  }


};
// console.log(this.user_name)
// console.log(this.link)
</script>
<style scoped>
.page-container {
  background-color: #fff;
  display: flex;
  justify-content: space-between;
  max-width: 1200px;
  margin: 10% auto;
  padding: 26px 30px;
  border-radius: 2px;
  overflow: hidden;
  height: auto;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}

.left-section,
.right-section {
  width: 49%;
  padding: 0 20px;
  box-sizing: border-box;
}

.section-title {
  font-size: 24px;
  margin-bottom: 20px;
}

.data-form {
  background-color: #fff;
  padding: 20px;
border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.form-label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
}

.form-item {
  margin-bottom: 15px;
}

.textarea-wrapper {
  width: 100%;
}

.textarea-wrapper {
  max-height: 500px;
  overflow-y: auto;
}

.textarea,
.input {
  padding: 10px;
  border: 1px solid #cccccc;
  border-radius: 4px;
  width: 95%;
}

/* 添加对齐样式 */
.textarea {
  resize: none;
  /* 禁止调整大小 */
  height: 250px;
  width: 95%;
}

.submit-button ,.submit-button-center{
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}
.submit-button-center{
  /* 居中 */
  margin: 0 auto;
  display: block;
  width: 50%;
}

.submit-button:hover {
  background-color: #0056b3;
}
.modal-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  /* 半透明背景 */
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal {
  background-color: #fff;
  padding: 20px;
  border: 1px solid #ccc;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
