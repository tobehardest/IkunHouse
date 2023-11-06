<template>
    <div class="register-wrap">
        <span class="back" @click="back">
          <i class="iconfont icon-fenxiang"></i>
        </span>
        <div class="form-wrap">
            <el-form ref="formRef" :rules="rules" :model="form" class="w-[250px]">
                <h1 class="title">注册</h1>
                <el-form-item prop="email">
                    <el-input class="input" type="email" placeholder="输入邮箱" autocomplete="off" v-model="form.email" id="email" />
                </el-form-item>
                <el-form-item prop="password">
                    <el-input class="input" type="password" placeholder="输入密码" v-model="form.password" id="password" />
                </el-form-item>
                <!-- <code-input :email="email" @code-tip="_emitTip" ref="codeInput" @code="_code"></code-input> -->
                <el-form-item>
                    <el-button round color="#626aef" class="w-[250px]" type="primary" @click="register" :loading="loading">
                        注册</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive ,defineEmits} from 'vue'
const formRef = ref(null)
const loading = ref(false)
const form = reactive({
    email: "",
    password: ""
})
const emit = defineEmits(['back'])

const rules = {
    email: [
        {
            required: true,
            message: '邮箱不能为空',
            trigger: 'blur'
        },
    ],
    password: [
        {
            required: true,
            message: '密码不能为空',
            trigger: 'blur'
        },
    ]
}
const register = () => {
    alert('注册')
}
const back = () => {
    emit('back',0)
}
</script>
<style scoped>
@import '../../assets/css/iconfont.css';
.register-wrap {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-image: -webkit-gradient(linear, 0% 0%, 0% 100%, from(rgb(107, 75, 171)), to(rgb(191, 9, 219)));
    z-index: 255;
    display: block flex;
    justify-content: center;
    align-items: center;
}
.title{
    font-size: 40px;
}
.back{
    display: block;
    position: absolute;
    left: 10px;
    top: 10px;
    color: rgb(9, 255, 0);
}
.iconfont{
    cursor: pointer;
    font-size: 20px;
    width: 30px;
    height: 30px;
    color: pink;
}
.iconfont:hover{
    color: aliceblue;
}


</style>