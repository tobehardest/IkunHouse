<template>
    <div class="login-wrap">
    <el-row class="login-container">
        <el-col :lg="4" :md="6" class="left">
            <div>
                <div>欢迎光临</div>
                <a @click="backHome">回到首页</a>
                <a class="register" @click="goRegister">注册</a>
            </div>
        </el-col>
        <el-col :lg="20" :md="18" class="right">
            <h2 class="title">欢迎回来</h2>
            <div>
                <span class="line"></span>
                <span>账号密码登录</span>
                <span class="line"></span>
            </div>
            <el-form ref="formRef" :rules="rules" :model="form" class="w-[250px]">
                <el-form-item prop="username">
                    <el-input v-model="form.username" placeholder="请输入用户名">
                        <template #prefix>
                            <el-icon><user /></el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password" v-model="form.password" placeholder="请输入密码" show-password>
                        <template #prefix>
                            <el-icon><lock /></el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button round color="#626aef" class="w-[250px]" type="primary" @click="onSubmit" :loading="loading">登 录</el-button>
                </el-form-item>
            </el-form>
            <Transition name="right-to-left">
                <register v-show="showRegister" @back="closeRister"></register>
            </Transition>

        </el-col>
    </el-row>
</div>
</template>

<script lang="ts" setup>
import { ref,reactive,onMounted,onBeforeUnmount } from 'vue'
import register from '@/components/common/register.vue'
import { useRouter } from "vue-router"
const showRegister = ref(false)
const router = useRouter()
//打开和关闭注册页面
const closeRister = ()=>{
    showRegister.value=false
}
const goRegister = ()=>{
    showRegister.value=true
}
const backHome = ()=>{
   
    router.push({
        path: '/home',
        query: {
            msg: 'hello'
        }
    })
}
// do not use same name with ref
const form = reactive({
  username:"",
  password:""
})
const rules = {
    username:[
        { 
            required: true, 
            message: '用户名不能为空', 
            trigger: 'blur' 
        },
    ],
    password:[
        { 
            required: true, 
            message: '密码不能为空', 
            trigger: 'blur' 
        },
    ]
}
const formRef = ref(null)
const loading = ref(false)
const onSubmit = () => {
    formRef.value.validate((valid)=>{
        if(!valid){
            alert('no permisson')
            return false
        }
        loading.value = true
    })
}
// 监听回车事件
function onKeyUp(e){
    if(e.key == "Enter") onSubmit()
}
// 添加键盘监听
onMounted(()=>{
    document.addEventListener("keyup",onKeyUp)
})
// 移除键盘监听
onBeforeUnmount(()=>{
    document.removeEventListener("keyup",onKeyUp)
})

</script>

<style scoped>
.login-wrap{
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    bottom: 0;
    z-index: 255;
    display: flex;
    justify-content: center;
    align-items: center;
    background-image: -webkit-gradient(linear, 0% 0%, 0% 100%, from(rgb(107, 75, 171)), to(rgb(191, 9, 219)));
}
.login-container{
    @apply min-h-screen bg-indigo-500;
}
.login-container .left,.login-container .right{
    @apply flex items-center justify-center;
}
.login-container .right{
    @apply bg-light-50 flex-col;
}
.left>div>div:first-child{
    @apply font-bold text-5xl text-light-50 mb-4;
}
.left>div>div:last-child{
    @apply text-gray-200 text-sm;
}
.right .title{
    @apply font-bold text-3xl text-gray-800;
}
.right>div{
    @apply flex items-center justify-center my-5 text-gray-300 space-x-2;
}
.right .line{
    @apply h-[1px] w-16 bg-gray-200;
}
.right-to-left-enter-active{
    animation: bounce 2s ease;
}
.right-to-left-leave-active{
    animation: bounce 2s ease reverse;
}
@keyframes bounce {
    0%{
        transform:scale(0);
    }
    50%{
        transform:scale(1.2);
    }
    100%{
        transform:scale(1);
    }
}
.register{
    cursor: pointer;
}
</style>