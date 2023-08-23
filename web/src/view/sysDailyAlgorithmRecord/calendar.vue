<template>
    <div class="calendar-container">
      <table class="calendar">
        <thead>
          <tr>
            <th v-for="day in daysInWeek" :key="day">{{ day }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(week, index) in calendar" :key="index">
            <td v-for="day in week" :key="day.date">
              <span
                :class="{
                  'day-cell': true,
                  'empty': day.date === '',
                  'checked': day.submitted,
                  'crossed': !day.submitted
                }"
              >
                {{ day.date }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  
  <script>
  export default {
    data() {
      return {
        // ... 省略其他数据 ...
  
        // 初始化一个日历数组
        calendar: [],
        daysInWeek: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
      };
    },
    created() {
      this.generateCalendar();
    },
    methods: {
      generateCalendar() {
        const now = new Date();
        const year = now.getFullYear();
        const month = now.getMonth();
        const daysInMonth = new Date(year, month + 1, 0).getDate();
  
        const firstDay = new Date(year, month, 1).getDay();
  
        const calendar = [];
        let week = [];
  
        for (let i = 0; i < firstDay; i++) {
          week.push({ date: '', submitted: false });
        }
  
        for (let day = 1; day <= daysInMonth; day++) {
          const currentDate = new Date(year, month, day);
          const submitted = this.checkSubmittedStatus(currentDate); // 根据实际情况确定是否已提交
  
          week.push({ date: day, submitted });
  
          if (week.length === 7) {
            calendar.push(week);
            week = [];
          }
        }
  
        if (week.length > 0) {
          while (week.length < 7) {
            week.push({ date: '', submitted: false });
          }
          calendar.push(week);
        }
  
        this.calendar = calendar;
      },
      checkSubmittedStatus(date) {
        // 根据日期查询用户是否提交
        // 返回 true 或 false
        // 这里只是一个示例，实际情况需要根据你的后端逻辑来判断
        return Math.random() < 0.5; // 随机生成提交状态
      }
    }
  };
  </script>
  <style scoped>
  .calendar-container {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 2% auto;
  }
  
  .calendar {
    border-collapse: collapse;
    border: 1px solid #ccc;
    font-family: Arial, sans-serif;
    width: 100%;
    max-width: 80%;
    background-color: #fff;
  }
  
  .calendar th {
    background-color: #f2f2f2;
    font-size: 2vw; /* 调整星期标题字体大小 */
    text-align: center;
    padding: 0.8vw; /* 调整内边距大小 */
  }
  
  .calendar td {
    border: 1px solid #ccc;
    text-align: center;
  }
  
  .day-cell {
    position: relative;
    display: inline-block;
    width: 5vw; /* 调整单元格大小 */
    height: 5vw; /* 调整单元格大小 */
    line-height: 5vw; /* 调整行高 */
    font-size: 2.5vw; /* 调整数字的大小 */
    text-align: center;
    cursor: pointer;
  }
  
  .empty {
    background-color: transparent;
  }
  
  .checked::before,
  .crossed::before {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    opacity: 0.3;
    font-size: 3vw; /* 调整图标的大小 */
  }
  
  .checked::before {
    content: '✔';
    color: green;
  }
  
  .crossed::before {
    content: '✖';
    color: red;
  }
  
  @media (max-width: 768px) {
    .calendar th,
    .calendar td {
      padding: 0.6vw; /* 调整内边距大小 */
    }
  
    .day-cell {
      width: 12px; /* 调整单元格大小 */
      height: 12px; /* 调整单元格大小 */
      line-height: 12px; /* 调整行高 */
      font-size: 8px; /* 调整数字的大小 */
    }
  
    .checked::before,
    .crossed::before {
      font-size: 9px; /* 调整图标的大小 */
    }
  }
  </style>