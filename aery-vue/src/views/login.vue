<template>
  <nav>
    <h1>Eary登录系统</h1>
  </nav>
  <div class="section_main">
    <section class="main">
      <!-- 公告区 -->
      <div class="display">
        <h2>重要通告</h2>
        <p>这是一个示例通告内容，请根据实际情况发布您的通告。</p>
      </div>

      <!-- 登录页面区域 -->
      <div class="login">
        <router-view></router-view>
      </div>
    </section>
  </div>

  <footer>
    <p>&copy; 盛森版权所有，禁止盗用</p>
    <a href="#">微信</a>
    <a href="#">QQ</a>
  </footer>
</template>

<script setup lang="ts">
</script>

<style scoped>
/* 全局样式 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* 页面背景和字体 */
body {
  font-family: Arial, sans-serif;
  background-color: #f5f5f5;
  color: #333;
}

/* 导航栏样式 */
nav {
  background-color: #3b5998;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

nav h1 {
  color: #fff;
  font-size: 32px;
  letter-spacing: 2px;
}

/* 主区域的Flex布局 */
.main {
  display: flex;
  justify-content: space-between;
  max-width: 1200px;
  margin: 20px auto;
}

/* 公告区样式 */
.display {
  background-color: #f9f9f9;
  border-left: 6px solid #3b5998;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  width: 60%; /* 公告区占60% */
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.display h2 {
  font-size: 28px;
  color: #333;
  margin-bottom: 10px;
}

.section_main{
  width: 100%;
  height: 65vh;
  margin-top: 60px;
}

.display p {
  font-size: 18px;
  color: #555;
  line-height: 1.6;
}

/* 登录区域样式 */
.login {
  border-radius: 8px;
  width: 35%; /* 登录页面占40%，留点空隙 */
}

/* 页脚样式 */
footer {
  background-color: #333;
  color: #fff;
  text-align: center;
  box-shadow: 0 -4px 6px rgba(0, 0, 0, 0.1);
  padding: 20px;
  margin-top: 44px;
}

footer p {
  margin-bottom: 10px;
}

footer a {
  color: #3b5998;
  margin: 0 10px;
  text-decoration: none;
  font-weight: bold;
}

footer a:hover {
  text-decoration: underline;
}

/* 调整大屏幕布局 */
@media (min-width: 1024px) {
  nav {
    padding: 30px;
  }

  nav h1 {
    font-size: 36px;
  }

  footer {
    padding: 30px;
  }
}
</style>
