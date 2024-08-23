<script lang="ts" setup>
import { reactive } from 'vue'
import { Greet, MergePdf, OptimizePdf, ImgToPdf } from '../../wailsjs/go/Application/App'
import { Upload } from '@element-plus/icons-vue'
const data = reactive({
  name: "",
  resultText: "Select a file to compress or files to merge",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}


function Compress() {
  OptimizePdf().then(result => {
    data.resultText = result
  })
}


function Merge() {
  MergePdf().then(result => {
    data.resultText = result
  })
}

function Img() {
  ImgToPdf().then(result => {
    data.resultText = result
  })
}
</script>

<template>
  <img id="logo" alt="Wails logo" src="../assets/images/pdf.png"/>
  <main>
    <div id="result" class="result"> <el-text class="mx-1" tag="b" type="warning" size="large">{{ data.resultText }}</el-text></div>
    <div id="input" class="input-box">
      <el-button type="warning" @click="Merge" plain>Merge<el-icon class="el-icon--right"><Upload /></el-icon></el-button> 
      <el-button type="warning"  @click="Compress" plain>Compress<el-icon class="el-icon--right"><Upload /></el-icon></el-button>
      <el-button type="warning"  @click="Img" plain>Images to pdf<el-icon class="el-icon--right"><Upload /></el-icon></el-button>

    </div>
  </main>
</template>

<style scoped>

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}


</style>
