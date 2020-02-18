<template>
  <div class="body">
      <div class="bookMsg">
          <div class="bookName">
              <span id="bookName">{{bookMsg.book_name}}</span>
          </div>
          <div class="bookbody">
              <div class="bookPng">
                  <img id="bookPng" :src="bookMsg.png_url" />
              </div>
              <div class="author">
                  <span id="book">作者: {{bookMsg.autor}}</span> 
              </div>
              <div class="publish">
                  <span id="book">出版社: {{bookMsg.press}}</span> 
              </div>
              <div class="publishtime">
                  <span id="book">出版时间: {{bookMsg.publish_time}}</span> 
              </div>
              <div class="pages">
                  <span id="book">页数: {{bookMsg.pages}}</span> 
              </div>
              <div class="isbn">
                  <span id="book">ISBN: {{bookMsg.isbn}}</span> 
              </div>
              <div class="dburl">
                  <span id="book">豆瓣链接: <a :href="bookMsg.dou_ban_url" style="text-decoration:none;color:red" target="_blank">豆瓣</a></span> 
              </div>
              <div class="score">
                  <span id="book">豆瓣评分: {{bookMsg.dou_ban_score}}</span> 
              </div>
          </div>
      </div>
          <div class="bookdesc">
              <span class="desctitle">
                  <i class="el-icon-notebook-2"></i>内容简介:
              </span>
              <div class="desc">
                  <span id="desc">{{bookMsg.content}}</span>
              </div>
          </div>
          <div class="send">
              <div class="sendtitle">
              <span class="desctitle">
                  <i class="el-icon-notebook-2"></i>相关书籍:
              </span>
              </div>
              <div class="sendbody">
                 <div v-for="i in recommend" id="sendbody">
                     <router-link :to="{name:'book', params: {bid:i.ID}}" style="text-decoration:none;color:black;" @click.native="reflash">
                    <div class="rbpng">
                        <img :src="i.PngUrl" id="rbpng">
                    </div>
                    <div ckass="rbn">
                      <span id="rbn">{{i.BookName}}</span>
                    </div>
                     </router-link>
                 </div>
              </div>
          </div>
          <div class="foot">
              <span class="desctitle">
                  <i class="el-icon-download"></i>下载:
              </span>
              <br>
              <span class="downdesc">文字版, 带目录, 文件格式: pdf, epub, mobi, azw3</span>
              <br>
              <div class="jd" v-if="bookMsg.buy_url">
              <a :href="bookMsg.buy_url" style="text-decoration:none;color:red" target="_blank">
              <span class="pan">购买正版(推荐)</span></a></div>
              <div class="yunpan" v-if="bookMsg.yun_pan_url">
              <router-link :to="{path:bookMsg.yun_pan_url}" style="text-decoration:none;color:red"><span class="pan">百度网盘</span></router-link>
              </div>
              <div class="jd" v-if="!bookMsg.buy_url&&!bookMsg.yun_pan_url">
                  <span class="pan">暂无资源</span></div>
                  <br>
            <span class="under">
               (如果下载链接不可用或者找不到书籍，可以加入到https://itpanda.net/进行反馈)   
              </span>
         </div>
  </div>
</template>

<script>
import {apiurl} from "../service/api.js"

export default {
  name: 'book',
  created(){
     var that = this;
     var id=that.$route.params.bid
     that.$axios({
        method:"get",
        url:apiurl.book+"/"+id,
    }).then(res=>{
        if(res.status==200){
            that.bookMsg = res.data.data.book
        }
    })
    that.$axios({
        method:"get",
        url:apiurl.recommend+id+"/recommend/"+id,
    }).then(res=>{
        if(res.status==200){
            that.recommend = res.data.data.recommend
        }
    })
  },
  data () {
    return { 
        bookMsg:null,
        downMsg:null,
        recommend:null,
    }
  },
  methods:{
       reflash(){
           this.$router.go(0);
       }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.body{
   position: relative;
   height: 1000px;
   width: 80%;
   background: white;
   margin-left: auto;
   margin-right: auto;
}
.bookMsg{
   width: 100%;
   height: 25%;

}
.bookName{
    position: relative;
    top:5%;
    left: 18%;
    height: 20px;
    width: 50%;
}
#bookName{
    font-size: 20px;
    font-weight: bolder;
    font-family: inherit;
    color: rgb(71, 69, 69)
}
.bookbody{
    position: relative;
    width: 100%;
    height: 80%;
    top:10%;
}
.bookPng{
    position: relative;
    float: left;
    left: 18%;
    top: 5px;
    height: 80%;
    width: 12%;
    border:1px solid rgba(196, 192, 192, 0.945)
}
#bookPng{
    width:100%;
	height:100%;
}
.author{
    position: relative;
    float: left;
    top: 5px;
    left: 20%;
    width: 55%;
}
#book{
    font-size: 15px;
    color: rgb(71, 69, 69);
}
.publish{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:10px;
}
.publishtime{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:15px;
}
.pages{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:20px;
}
.isbn{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:25px;
}
.dburl{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:30px;
}
.score{
    position: relative;
    float: left;
    width: 55%;
    left: 20%;
    top:35px;
}
.bookdesc{
    position: relative;
    top:1%;
    width: 100%;
    height: 30%;
}
.desctitle{
    position: relative;
    left: 18%;
    font-size: 15px;
    color: rgba(219, 26, 26, 0.938);
}
.sendtitle{
    height: 8%;
    width: 100%;
}
.sendbody{
    position: relative;
    top: 15px;
    height: 80%;
    width: 80%;
    left: 18%;
}
#sendbody{
    position: relative;
    float: left;
    width: 17%;
    height: 85%;
    margin:10px;
}
.rbpng{
    position: relative;
    width: 80%;
    height: 80%;
    padding-bottom: 5px;
    border:1px solid rgba(219, 214, 214, 0.671);
}
#rbpng{
    width: 100%;
    height: 100%;
}
.rbn{
    position: relative;
    width: 100%;
}
#rbn{
    font-size: 13px;
}

.desc{
    position: relative;
    width: 75%;
    top:8px;
    left: 18%;
    height: 70%;
}
#desc{
    font-size: 15px;
    font-family: cursive;
}
.send{
    position: relative;
    top:2%;
    width: 100%;
    height: 28%;
}
.downdesc{
    position: relative;
    top:10px;
    left: 18%;
    font-size: 15px;
    font-family: cursive;
}
.pan{
    position: relative;
    left: 18%;
    top:20px;
    font-size: 15px;
    font-family:monospace;
    font-weight: bold;
}
.foot{
    position: relative;
    top:3%;
    width: 100%;
    height: 12%;
}
.under{
    position: relative;
    float: left;
    top:30px;
    left: 18%;
    font-size: 12px;
    font-family: cursive;
}
.jd{
    position: relative;
    float: left;
    left: 16%;
}
.yunpan{
    position: relative;
    float: left;
    left: 20%;
}
</style>