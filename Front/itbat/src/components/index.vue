<template>
    <div class="body">
        <div class="ul1">
          <span class="title1">书籍分类</span>
          <ul>

          </ul>
        </div>
        <div class="ul2">
            <div class="head">
               <span class="title">所有书籍</span>
               <a href="#" class="score" @click="score"><i class="el-icon-top">按好评排序</i></a>
            </div>
            <div class="content">
              <ul class="ul">
                  <li v-for="i in books" class="li">
                    <div class="img">
                      <img :src="i.png_url" class="img" border="1"/>
                    </div>
                    <div class="ceng">
                      <div class="bookname">
                        <a href="#" style="text-decoration:none;">
                        <span>{{i.book_name}}</span>
                        </a>
                      </div>
                      <div class="autor">
                        <span>作者:{{i.autor}}</span>
                      </div>
                    </div>
                    <br/>
                    <br/>
                     <div class="desc">
                        <span style="font-size:12px;">{{"简述:  "+i.content.substr(0,201)+"..."}}</span>
                        <a href="#" style="text-decoration:none;color:red;font-size:10px">详情</a>
                    </div>
                  </li>
              </ul>
            </div>
            <div class="foot">
              <el-pagination
              class="fenye"
              background
              layout="prev, pager, next"
              @current-change="handleCurrentChange"
              :current-page.sync="currentPage"
              :total="pages">
              </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import {apiurl} from "../service/api.js"

export default {
  name: 'index',
  created(){
       var that = this;
        this.$axios({
            type:"get",
            url:apiurl.index,
        }).then(res=>{
            if(res.status==200){
               that.books = res.data.data.books
               that.pages = res.data.data.row
            }
        })
  },
  data () {
    return {
        books:null,
        pages : 0, 
        currentPage:1,
    }
  },
  methods:{
      handleCurrentChange(val) {
        var that = this;
        this.$axios({
            type:"get",
            url:apiurl.index+'?page='+val,
        }).then(res=>{
            if(res.status==200){
               this.books = null
               that.books = res.data.data.books
               that.pages = res.data.data.row
            }
        })
      },
      score(){
          // var that = this;
          // this.$axios({
          //     type:"post",
          //     url:apiurl.score,
          // }).then(res=>{
          //     if(res.status==200){
          //         this.books = null
          //         that.books = res.data.data.books
          //         that.pages = res.data.data.row
          //     }
          // })
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.body{
   position: relative;
   height: 1800px;
   width: 80%;
   background: white;
   margin-left: auto;
   margin-right: auto;
}
.el-aside {
  background-color:white;
  color: #333;
  text-align: center;
}
.el-header{
   background-color: rgb(253, 254, 255);
  color: #333;
  text-align: center;
  line-height: 60px;
}
.el-main {
  background-color: white;
}
.ul1{
  position: relative;
  top: 0;
  float: left;
  height: 100%;
  width: 30%;
}
.ul2{
  position: relative;
  top: 0;
  float: left;
  height: 100%;
  width: 70%;
}
.head{
  position: relative;
  top: 0;
  width: 100%;
  height: 4%;
  background: white;
}
.title{
  position: relative;
  float: left;
  top:20px;
  left:10px;
  font-family: Arial, Helvetica, sans-serif;
  color: #333;
  font-size: 25px;
  font-weight: bolder;
}
.score{
  position: relative;
  float: left;
  top:30px;
  left:100px;
  font-family: Arial, Helvetica, sans-serif;
  color: red;
  font-size: 15px;
  font-weight: bolder;
}
.title1{
  position: relative;
  top:20px;
  left:80px;
  font-family: Arial, Helvetica, sans-serif;
  color: #333;
  font-size: 25px;
  font-weight: bolder;
}
.content{
  position: relative;
  width: 100%;
  height: 92%;
}
.ul{
  padding:0;
  height: 100%;
  margin:0;
  text-align: left;
  list-style-type: none;
}
.li{
  position: relative;
  display:block;
  height: 9.7%;
}
.img{
  position: relative;
  float: left;
  height: 150px;
  width: 120px;
}
.ceng{
  position: relative;
  float: left;
}
.bookname{
  position: relative;
  left: 15px;
  font-size: 20px;
  font-family: serif;
  font-weight: bold;
}
.autor{
  position: relative;
  left: 15px;
  font-family: serif;
}
.desc{
  position: relative;
  width: 80%;
  left: 15px;
  font-size: 5px;
}
.foot{
  position: relative;
  top:0%;
  width: 100%;
  height: 4%;
}
.fenye{
  position: relative;
  left: 5%;
}
</style>
