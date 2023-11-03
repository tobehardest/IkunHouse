<template>
  <div id="body">
    <div @click="onPlay()">chen 视频</div>

    <div class="video-container">
      <!-- <video :src="curSrc" class="video" autoplay controls>
        您的浏览器不支持 HTML5 video 标签。
      </video> -->
      <video ref="video" :class="{ 'video-hide': !playVideoTag, video }" controls autoplay
        @ended="onEnded('video',false)"></video>
      <video ref="video1" :class="{ 'video-hide': playVideoTag, video }" controls autoplay
        @ended="onEnded('video1',false)"></video>
    </div>
    <div class="gg">
      <button @click="onPlay()">play1</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
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


const onEnded = (e: any,prev:any) => {
  if (canPlayNext || canPlayPrev) {
    const video3 = e === 'video' ? video1.value : video.value
    playVideoTag.value = !playVideoTag.value
    if(prev){
      if(curIndex>=1){
        video3.src = srcList.value[--curIndex]
      }else{
        return 
      }
    }else{
      if(curIndex<srcList.value.length-1){
        video3.src = srcList.value[++curIndex]
      }else{
        return
      }
    }
    console.log('curinxex==',curIndex);
    setTimeout(() => {
      video3.play()
    }, 300)
    if(prev){
      console.log('加载上一个视频');
      loadPrev()
    }else{
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
    canPlayPrev=true
    return
  }
  let next = curIndex+1
  console.log('cur===',curIndex);
  video2.src = srcList.value[next]
  video2.addEventListener('canplay', e => e.target.pause(), { once: true })
}
const loadPrev = () => {
  console.log('预加载上一个视频');
  const video2 = playVideoTag.value ? video1.value : video.value
  if (0 >= curIndex) { // 当前索引是列表的第一个
    console.log('没有上一个视频')
    canPlayPrev = false
    canPlayNext=true
    return
  }
  canPlayPrev=true
  let next = curIndex-1
  video2.src = srcList.value[next]
  video2.addEventListener('canplay', e => e.target.pause(), { once: true })
}
const onPlay = () => {
  console.log('play');
  playVideo(srcList2)
}
document.onkeydown = function (e) {    //对整个页面监听  
  var keyNum = window.event ? e.keyCode : e.which;       //获取被按下的键值  
  //判断如果用户按下了上键（keycody=13）  
  if (keyNum == 38) {
    if(playVideoTag.value){//如果当前是第一个video在播放
      video.value.pause()
      console.log('video1==',video.value);
      onEnded('video',true)
    }else{
      video1.value.pause()
      console.log('video2==',video1.value);
      onEnded('video1',true)
    }
  }
  if (keyNum == 40) {
    if(!playVideoTag.value){//如果当前是第2个video在播放
      console.log('video1==',video1.value);
      video1.value.pause()
      onEnded('video1',false)
    }else{
      console.log('video==',video.value);
      video.value.pause()
      onEnded('video',false)
    }
  }
}
onMounted(()=>{
  video.value.src = srcList.value[0]
    
})
</script>
<style>
#body {
  width: 100%;
  height: 100%;
  background-color: aquamarine;
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
  animation: slide 0.5s;
}
@keyframes slide{
  from{
    top:0px;
  }
  to{
    top:-1900px;
  }
}
</style>