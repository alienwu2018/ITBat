<template>
    <div class="msg">
        <div class="box">
        <span class="link">地址:<a :href="data.p_url" style="text-decoration:none;color:red;" target="_blank">{{data.p_url}}</a></span>
        <br>
        <span class="code">提取码:{{data.pwd}}</span>
        </div>
    </div>

</template>

<script>
import {apiurl} from "../service/api.js"

export default {
    created(){
        var that = this;
        var id=that.$route.params.bid
        //首次访问,获取首页的数据
        that.$axios({
            method:"get",
            url:apiurl.yun+id+"/download/"+id,
        }).then(res=>{
            if(res.status==200){
                that.data = res.data.data.downlink
            }
        })
    },
    data(){
        return{
            data:null,
        }
    },
    methods:{
  
    }
    
}
</script>

<style scoped>
.downloaddesc{
    position: relative;
    height: 500px;
    width: 800px;
    text-align: center;
    background: red;
}
.msg{
    text-align: center;
    height: 200px;
    background: white;
}
.box{
    position: relative;
    height: 100px;
    width: 60%;
    top:15%;
    left:20%;
    background: rgba(138, 221, 138, 0.842);
}
.link{
    position: relative;
    top:10%;
}
.code{
    position: relative;
    top:15%;
}
</style>