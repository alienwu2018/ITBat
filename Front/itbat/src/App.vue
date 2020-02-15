<template>
  <div id="app">
    <el-container class="container">
      <el-header>
        <div class="title">
          <a href="/" style="text-decoration:none;">
            <span id="title">ITBat</span>
          </a>
        </div>
        <div class="logo">
            <el-avatar :size="60" :src="logo" @error="errorHandler">
            <img :src="logo"/>
          </el-avatar>
        </div>
        <div class="search">
            <el-input v-model="text" placeholder="搜索" @keyup.enter.native="search"></el-input>
        </div>
      </el-header>
      <el-main>
          <router-view/>
      </el-main>
      <el-footer>
      </el-footer>
      </el-container>
  </div>
</template>

<script>
import {apiurl} from "../src/service/api.js"

export default {
  name: 'App',
  data(){
    return {
        logo : require('../src/assets/baticon.png'),
        text: '',
    }
  },
  methods :{
      search(){
        var that = this
        var keyword = that.text;
        that.$router.push({ path: '/search', query: {query:keyword} })
        this.$router.go(0)
      }
  }
}
</script>

<style scoped>
.container{
  height: 2000px;
}
.el-header {
  background-color: rgb(253, 254, 255);
  color: #333;
  text-align: center;
  line-height: 60px;
}
.el-footer{
  background-color: rgba(34, 84, 134, 0.603);
  color: #333;
  text-align: center;
  line-height: 60px;
}

.title{
  position: relative;
  float: left;
  left: 50px;
}
#title{
  font-family: monospace;
  font-size: 30px;
  color:cadetblue;
  font-weight: bolder;
}
.logo{
  position: relative;
  float: left;
  left: 100px;
}
.search{
  position: relative;
  float: left;
  left : 65%;
}
.el-main {
  background-color: #E9EEF3;
  color: #333;
  /* text-align: center; */

}

</style>
