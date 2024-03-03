<template>
  <div id="body">
    <div class="bg-light-500">happy 视频</div>
    <!-- <svg-icon name="vue" />
    <svg class="icon" aria-hidden="true"> -->
    <!-- <use xlink:href="#icon-heart"></use>
    </svg> -->
    <div class="video-container">
      <!-- <video :src="curSrc" class="video" autoplay controls>
        您的浏览器不支持 HTML5 video 标签。
      </video> -->
      <svg-icon name="page-svg" />
      <video ref="video" :class="{ 'video-hide': !playVideoTag, video }" controls autoplay
        @ended="onEnded('video', false)"></video>
      <video ref="video1" :class="{ 'video-hide': playVideoTag, video }" controls autoplay
        @ended="onEnded('video1', false)"></video>
      <div class="side-bar ">
        <span class="userAvatar">
          <el-avatar :size="35" :src="avarUrl" />
        </span>
        <span class="like-heart" @click="onLike">
          <svg class="icon" aria-hidden="true">
            <use xlink:href="#icon-heart"></use>
          </svg>
        </span>
        <span class="comment" @click="onComment">
          <svg class="icon iconMeaasge" aria-hidden="true">
            <use xlink:href="#icon-message"></use>
          </svg>
        </span>
        <span class="share" @click="onShare">
          <!-- <svg class="icon iconShare" aria-hidden="true">
            <use xlink:href="#icon-share"></use>
          </svg> -->
          <el-icon color="#409EFC" class="no-inherit">
            <Share />
          </el-icon>

        </span>
      </div>
      <el-button :icon="Search" circle />
    </div>
    <div class="playButton">
      <button @click="onPlay()">play</button>
    </div>
    <div class="conmmentPanel">
      <el-drawer v-model="visible" :show-close="false" class="eldrawer">
        <template #header="{ close, titleId, titleClass }">
          <h4 :id="titleId" :class="titleClass">评论区</h4>
          <el-button type="danger" @click="close">
            <el-icon class="el-icon--left">
              <CircleCloseFilled />
            </el-icon>
            Close
          </el-button>
        </template>
        <div class="commentList">
          <div v-for="(value, key) in comments " :key="key">
            {{ value.comment }}
          </div>
        </div>
      </el-drawer>
    </div>
    <el-button text @click="open">Click to open the Message Box</el-button>

  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
//import { createUser } from '@/api/test'
import { ElButton, ElDrawer, ElMessage, ElMessageBox } from 'element-plus'
import type { Action } from 'element-plus'
import {
  Search,
  Share,
  CircleCloseFilled
} from '@element-plus/icons-vue'
const open = () => {
  ElMessageBox.alert('This is a message', 'Title', {
    // if you want to disable its autofocus
    // autofocus: false,
    confirmButtonText: 'OK',
    callback: (action: Action) => {
      ElMessage({
        type: 'info',
        message: `action: ${action}`,
      })
    },
  })
}
const visible = ref(false)
let curIndex = 0
let playVideoTag = ref(true) // 当前播放的 ref 是否为 video
let canPlayNext = true   // 是否存在下一个播放地址
let canPlayPrev = false
const srcList2 = ['https://video.pearvideo.com/mp4/adshort/20200909/cont-1695119-15373301_adpkg-ad_hd.mp4',
  'https://media.w3.org/2010/05/sintel/trailer.mp4', 'https://video.pearvideo.com/mp4/adshort/20180407/cont-1316943-11829224_adpkg-ad_hd.mp4',
  'https://video.pearvideo.com/mp4/adshort/20200309/cont-1659656-14991419_adpkg-ad_hd.mp4',
  'https://video.pearvideo.com/mp4/adshort/20190201/cont-1513779-13549419_adpkg-ad_hd.mp4'
]
const video = ref()
const video1 = ref()
const srcList = ref([])
const avarUrl = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'

const onEnded = (e: any, prev: any) => {
  if (canPlayNext || canPlayPrev) {
    const video3 = e === 'video' ? video1.value : video.value
    playVideoTag.value = !playVideoTag.value
    if (prev) {
      if (curIndex >= 1) {
        video3.src = srcList.value[--curIndex]
      } else {
        return
      }
    } else {
      if (curIndex < srcList.value.length - 1) {
        video3.src = srcList.value[++curIndex]
      } else {
        return
      }
    }
    console.log('curinxex==', curIndex);
    setTimeout(() => {
      video3.play()
    }, 300)
    if (prev) {
      console.log('加载上一个视频');
      loadPrev()
    } else {
      console.log('加载下一个视频');
      loadNext()
    }
  }
}
const playVideo = (src: any) => {  // 可以接受单个 url 和 url 列表
  video.value.src = src[curIndex]
  srcList.value = src
  loadNext()
}

const loadNext = () => {
  console.log('预加载下一个视频');
  const video2 = playVideoTag.value ? video1.value : video.value
  if (srcList.value.length - 1 === curIndex) { // 当前索引是列表的最后一个
    console.log('没有新视频')
    canPlayNext = false
    canPlayPrev = true
    return
  }
  let next = curIndex + 1
  console.log('cur===', curIndex);
  video2.src = srcList.value[next]
  video2.addEventListener('canplay', e => e.target.pause(), { once: true })
}
const loadPrev = () => {
  console.log('预加载上一个视频');
  const video2 = playVideoTag.value ? video1.value : video.value
  if (0 >= curIndex) { // 当前索引是列表的第一个
    console.log('没有上一个视频')
    canPlayPrev = false
    canPlayNext = true
    return
  }
  canPlayPrev = true
  let next = curIndex - 1
  video2.src = srcList.value[next]
  video2.addEventListener('canplay', e => e.target.pause(), { once: true })
}
const onPlay = () => {
  console.log('play');

  playVideo(srcList2)
}
//点赞 评论  分享视频 
const onLike = () => {
  alert(1)
}
const onComment = () => {
  visible.value = true
}
const onShare = () => {
  alert(3)
}
document.onkeydown = function (e) {    //对整个页面监听  
  var keyNum = window.event ? e.keyCode : e.which;       //获取被按下的键值  
  //判断如果用户按下了上键（keycody=13）  
  if (keyNum == 38) {
    if (playVideoTag.value) {//如果当前是第一个video在播放
      video.value.pause()
      console.log('video1==', video.value);
      onEnded('video', true)
    } else {
      video1.value.pause()
      console.log('video2==', video1.value);
      onEnded('video1', true)
    }
  }
  if (keyNum == 40) {
    if (!playVideoTag.value) {//如果当前是第2个video在播放
      console.log('video1==', video1.value);
      video1.value.pause()
      onEnded('video1', false)
    } else {
      console.log('video==', video.value);
      video.value.pause()
      onEnded('video', false)
    }
  }
}

const comments = [
  { comment: '你好' },
  { comment: '我是好人' },
  { comment: '好搞笑' },
]
onMounted(async () => {
  video.value.src = srcList.value[0]
  //let ans = await createUser({})
  // console.log(ans);

})
</script>
<style scoped>
@import '../assets/css/iconfont.css';

#body {
  height: 100%;
  background-color: rgb(133, 127, 255);
}

.playButton {
  margin-top: 5px;

}

.playButton>button {
  background-color: antiquewhite;
}

.icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

.video-container {
  position: absolute;
  left: 0;
  right: 0;
  top: 50px;
  bottom: 0;
  margin: auto;
  width: 100%;
  height: 90%;
}

.titleClass {
  widows: 200px;
}

.side-bar {
  display: flex;
  flex-direction: column;
  background-color: rgba(159, 220, 38, 0.8);
  overflow: hidden;
  width: 50px;
  height: 300px;
  position: absolute;
  right: 20px;
  top: 20%;
}

.userAvatar {
  display: flex;
  justify-content: center;
  align-items: center;
  flex: 1;
  cursor: pointer;
}

.like-heart {
  font-size: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex: 1;
  cursor: pointer;
}

.like-heart:hover {
  font-size: 30px;
  flex: 1;
  cursor: pointer;
  animation: likee 1s infinite;
  animation-direction: alternate-reverse;
}

.conmmentPanel {
  position: absolute;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
}

::v-deep .eldrawer {
  position: absolute;
  top: 80px;
  height: 80% !important;
}

.comment {
  font-size: 30px;
  display: flex;
  ;
  flex: 1;
  cursor: pointer;
  justify-content: center;
  align-items: center;
}

.comment:hover {
  color: blue;
}

.commentList {
  background-color: rgb(233, 230, 227);
}

.commentList div {
  height: 30px;
  color: aquamarine;
  margin-bottom: 5px;
}

.share {
  z-index: 10;
  font-size: 30px;
  display: block;
  cursor: pointer;
}

.video {
  width: 90%;
  height: 90%;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
}

video {
  object-fit: contain;
  overflow-clip-margin: content-box;
  overflow: clip;
}

.video-hide {
  position: absolute;
  left: -9999px;
  top: -999px;
}

@keyframes slide {
  from {
    top: 0px;
  }

  to {
    top: -1900px;
  }
}

@keyframes likee {
  from {
    color: rgba(202, 134, 121, 0.5)
  }

  to {
    color: rgba(252, 47, 6, 1)
  }
}</style>