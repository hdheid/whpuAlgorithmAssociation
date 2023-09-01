<template>
  <div>
    <warning-bar title="注：默认按照打卡次数排序，总用户数前20%（上取整）的用户标新" />
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column type="index" align="center" :resizable="false" label="排名" width="100">
          <template #default="scope">
            <div :class="getRankClass(scope.$index)">
              <!-- 根据前三名显示不同的奖牌 -->
              <template v-if="isSpecialRank(scope.$index)">
                <i>✨</i>
              </template>
              {{ scope.$index + 1 }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="昵称" min-width="150" prop="nickName" />
        <el-table-column align="center" label="本月打卡天数" min-width="180" prop="da_count_in_mouth" />
      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>

  </div>
</template>
  
<script>


</script>
  
<script setup>

import {
  getUserList,
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref, watch } from 'vue'

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName,
          children: []
        }
        setAuthorityOptions(item.children, option.children)
        optionsData.push(option)
      } else {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName
        }
        optionsData.push(option)
      }
    })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const limitSize = ref(0)
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
//打印tableData
// console.log(tableData)
// 查询
const getTableData = async () => {
  const table = await getUserList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize

    limitSize.value = total.value / 5
    console.log(limitSize.value)
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async () => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

// 排名的函数
const getRankClass = (index) => {
  // 根据索引返回不同的样式类
  if (index < limitSize.value) {
    return 'special-rank';
  } else {
    return 'common-rank';
  }
}

const isSpecialRank = (index) => {
  return index < limitSize.value
}

const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}


const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}



</script>
<style lang="scss">
.special-rank {
  font-size: 26px;
  font-weight: bold;
  color: #ff9b54;
  /* 橙色 */
}

.common-rank {
  font-size: 26px;
  font-weight: bold;
  color: #606060;
  /* 灰色 */

}

.user-dialog {

  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }

  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

.nickName {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.pointer {
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
  