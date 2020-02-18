<template>
  <div class="body">
    <div class="ul1">
      <span class="title1">书籍分类</span>
      <el-tree class="tree" :data="tree" :props="defaultProps" @node-click="handleNodeClick"></el-tree>
    </div>
    <div class="ul2">
      <div class="head">
        <span class="title">所有书籍</span>
        <a href="javascript:void(0);" class="score">
          <i class="el-icon-top">按好评排序</i></a>
      </div>
      <div class="content">
        <ul class="ul">
          <li v-for="i in books" class="li">
            <div class="imgdiv">
              <img :src="i.png_url" class="img" border="1" />
            </div>
            <div class="ceng">
              <div class="bookname">
               <router-link :to="{name:'book', params: {bid:i.id}}" style="text-decoration:none;color:black" target="_blank">
                  <span>{{i.book_name}}</span></router-link>
              </div>
              <div class="autor">
                <span>作者:{{i.autor}}</span></div>
                          <div class="desc">
              <span style="font-size:12px;">{{"简述: "+i.content.substr(0,201)+"..."}}</span>
              <router-link :to="{name:'book', params: {bid:i.id}}" style="text-decoration:none;color:red;font-size:10px" target="_blank">
                 详情</router-link></div>
            </div>
          </li>
        </ul>
      </div>
      <div class="foot">
        <el-pagination class="fenye" background layout="prev, pager, next" @current-change="handleCurrentChange" :current-page.sync="currentPage" :total="pages"></el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import {apiurl} from "../service/api.js"

export default {
  name: 'search',
  created(){
        var that = this;
        var keyword=that.$route.params.BigCategory
        // // //首次访问,获取类别的数据
        that.$axios({
            type:"get",
            url:apiurl.category+keyword,
        }).then(res=>{
            if(res.status==200){
                that.books = res.data.data.books
                that.pages = res.data.data.row
            }
        })
        //获取书籍的总类别渲染树形结构
        that.$axios({
            method:"get",
            url:apiurl.docategories,
        }).then(res=>{
            if(res.status==200){
              that.mapper = res.data.data.relation
            }
        })
  },
  data () {
    return {
        books:null,     //书籍的数据
        pages : 0,      //总页数
        currentPage:1,  //当前页数
        isscore : false,  //是否按评分排序
        mapper:null,      //树形结构中文名字对应的接口参数映射

        tree:[{label:'前端开发',children:[{label:'JavaScript'},{label:'Node.js'},{label:'HTML&CSS'},{label:'框架和库'},{label:'其它'}]},
              {label:'后端开发',children:[{label:'C'},{label:'C++'},{label:'Python'},{label:'Java'},{label:'Go'},{label:'PHP'},{label:'C#'},{label:'Rust'},{label:'Erlang'},{label:'其它'}]},
              {label:'移动开发',children:[{label:'IOS'},{label:'Andriod'},{label:'其它'}]},
              {label:'机器学习'},{label:'数据挖掘'},{label:'大数据与云计算'},{label:'网络编程'},{label:'算法与数据结构'},
              {label:'数据库'},{label:'操作系统'},{label:'编译原理'},{label:'计算机科学'},{label:'软件开发与软件工程'},
              {label:'网络安全'},{label:'运维管理'},{label:'游戏开发与游戏设计'},{label:'硬件'},{label:'搜索引擎'},
              {label:'区块链'},{label:'消息队列'},{label:'交互设计'},{label:'互联网与科普'},{label:'其他'}]
    }
  },
  methods:{
      //操作树形结构的事件
      handleNodeClick(data,node) {
          var father = this.mapper[node.parent['label']]
          var son = this.mapper[data['label']]
          if(father==undefined){
             this.$router.push({ name: 'bigcategory', params: {BigCategory:son}})
          }else{
             this.$router.push({ name: 'smallcategory', params: {BigCategory:father,SmallCategory:son}})
          }
      },
      //分页组件的事件
      handleCurrentChange(val) {
          var that = this
          var keyword=that.$route.params.BigCategory
          that.$axios({
          method:"get",
          url:apiurl.category+keyword,
          params:{
            "page":val,
          }
          }).then(res=>{
              if(res.status==200){
                that.books = res.data.data.books
                // that.pages = res.data.data.row
              }
          })
      },
  },
  watch: {
      $route(){  //监听路由参数变化
        var that = this
        that.currentPage=1
        var father = that.$route.params.BigCategory
        var son = that.$route.params.SmallCategory
        if (son==undefined){
            that.$axios({
            type:"get",
            url:apiurl.category+father,
            }).then(res=>{
                if(res.status==200){
                    that.books = res.data.data.books
                    that.pages = res.data.data.row
                }
            })
        }else{
           this.$router.push({ name: 'smallcategory', params: {BigCategory:father,SmallCategory:son}})
        }
      },
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
.tree{
  position: relative;
  top:50px;
  left: 50px;
  width: 60%;
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
.imgdiv{
  position: relative;
  float: left;
  height: 150px;
  width: 15%;
}
.img{
  width:100%;
	height:100%;
}
.ceng{
  position: relative;
  left: 2%;
  float: left;
  width: 75%;
}
.bookname{
  /* position: relative; */
  width: 100%;
  left: 15px;
  font-size: 20px;
  font-family: serif;
  font-weight: bold;
}
.autor{
  /* position: relative; */
  width: 100%;
  left: 15px;
  font-family: serif;
}
.desc{
  /* position: relative; */
  width: 100%;
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