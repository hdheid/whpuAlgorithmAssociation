<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="日期:" prop="date">
          <el-date-picker v-model="formData.date" type="date" :readonly="true" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="用户名:" prop="user_name">
          <el-input v-model="nickName" :clearable="true" :readonly="true" />
        </el-form-item>
        <el-form-item label="链接:" prop="link">
          <el-input v-model="formData.link" :clearable="true" :readonly="true" />
        </el-form-item>
        <el-form-item label="代码:" prop="code">
          <el-input v-model="formData.code" :clearable="true" :readonly="true" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DailyAlgorithmRecord',
  props: {
    link: {
      type: String,
    },
    code: {
      type: String,
    }
  },
  // 使用$emit方法改变父组件的showmodal
  methods: {
    back() {
      this.$emit('change-modal');
    }
  }

}
</script>

<script setup>
import {
  createDailyAlgorithmRecord,
  updateDailyAlgorithmRecord,
  findDailyAlgorithmRecord,
  coverDailyAlgorithmRecord
} from '@/api/sysDailyAlgorithmRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { onMounted, watch, toRefs } from 'vue'

import { useUserStore } from '@/pinia/modules/user'
import SelectImage from '@/components/selectImage/selectImage.vue'
import SysDailyAlgorithmRecord from './sysDailyAlgorithmRecord.vue'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
// Props
const props = defineProps({
  link: String,
  code: String
})

// Convert props to refs
const { link, code } = toRefs(props)

const type = ref('')
const formData = ref({
  date: new Date(),
  user_name: '',
  link: '',
  code: '',
})
//在提交信息框现实的是nickname，提交给数据库的是uuid
const nickName = ref('')
// Initialize formData with props values
onMounted(() => {
  formData.value.link = link.value
  formData.value.code = code.value
  formData.value.user_name = userStore.userInfo.uuid
  nickName.value = userStore.userInfo.nickName
})
// 验证规则
const rule = reactive({
  date: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  user_name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  link: [{
    required: true,
    message: '题目链接为必填',
    trigger: ['input', 'blur'],
  }],
  code: [{
    required: true,
    message: '题目代码为必填',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()
// 上一次点击save按钮的时间,和这次的时间差
var lastClickTime = ref(0)
var currentTime = ref(0);

// 初始化方法
const init = async () => {
  lastClickTime = Date.now();

  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findDailyAlgorithmRecord({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reDAR
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async () => {
    currentTime = Date.now();

  // await 关键字用于等待 coverDailyAlgorithmRecord 函数返回一个 Promise 对象的结果。这意味着在接收到结果之前，代码将暂停执行，并等待 Promise 对象的解决（即成功或失败）。
  // console.log(formData.value)
  const resCoverRecord = await coverDailyAlgorithmRecord({ date: formData.value.date, user_name: formData.value.user_name })
  // console.log(resCoverRecord)
  if ("reDAR" in resCoverRecord.data) {
    resCoverRecord.data.reDAR.link = formData.value.link
    resCoverRecord.data.reDAR.code = formData.value.code
    type.value = 'update'
  } else {
    type.value = 'create'
  }
  // console.log(type.value)

  // 时间间隔不超过3秒,拒绝执行
  if (currentTime - lastClickTime < 2000) {
    type.value = 'tooFast'
  }
  // console.log(currentTime)


  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'tooFast':
        ElMessage({
          type: 'warning',
          message: '操作过快'
        })
        return
      case 'create':
        res = await createDailyAlgorithmRecord(formData.value)
        break
      case 'update':
        res = await updateDailyAlgorithmRecord(resCoverRecord.data.reDAR)
        break
      default:
        res = await createDailyAlgorithmRecord(formData.value)
        break
    }
    if (res.code === 0) {
      switch (type.value) {
        case 'create':
          ElMessage({
            type: 'success',
            message: '上传每日一题成功'
          })
          break
        case 'update':
          ElMessage({
            type: 'success',
            message: '更新每日一题记录成功'
          })
          break
        default:
          ElMessage({
            type: 'success',
            message: '上传每日一题成功'
          })
          break
      }
    }
  })
  lastClickTime = currentTime
}

// 返回按钮

</script>

<style></style>
