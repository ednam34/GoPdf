<script lang="ts" setup>
import { ref } from 'vue';
import type { TabsPaneContext } from 'element-plus'
import { Odometer, DocumentAdd, DocumentRemove, DocumentCopy, PictureFilled, UploadFilled } from '@element-plus/icons-vue'
import { OnFileDrop } from "../wailsjs/runtime/runtime";
import PdfView from './components/PdfView.vue';
import PdfMove from './components/PdfMove.vue';
import { Compress, ImageConv, Merge, OpenFile, SavePdf, RemovePg, OpenFileFromPath, MoovePage } from './file'
import { PrintAny, PrintString } from '../wailsjs/go/Application/App';
import { isView } from './shared';

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}

const isViewTs = (): boolean => {
  return isView.value
}


const Test = () => {
  isView.value = !isView.value
}

OnFileDrop((x, y, paths) => {
  PrintString('Position X:' + x.toString);
  PrintString('Position Y:' + y.toString);
  PrintString('Chemin du fichier déposé:' + paths[0]);
  OpenFileFromPath(paths[0])
}, true);


</script>

<template>

  <div class="common-layout" style="--wails-drop-target:drop">
    <el-container>
      <el-aside width="200px" class="side">
        <el-space direction="vertical" size="large" class="space">
          <h1>Open File :</h1>
          <el-button plain @click="OpenFile">Open<el-icon class="el-icon--right" size="large">
              <UploadFilled />
            </el-icon></el-button>

          <el-button plain @click="SavePdf">Save<el-icon class="el-icon--right" size="large">
              <DocumentAdd />
            </el-icon></el-button>

          <h1>All Tools :</h1>
          <el-button plain @click="Merge">Merge<el-icon class="el-icon--right" size="large">
              <DocumentCopy />
            </el-icon></el-button>
          <el-button plain @click="Compress">Optimize<el-icon class="el-icon--right" size="large">
              <Odometer />
            </el-icon></el-button>
          <el-button plain @click="ImageConv">Image to pdf<el-icon class="el-icon--right" size="large">
              <PictureFilled />
            </el-icon></el-button>
          <el-button plain @click="Test">Change View<el-icon class="el-icon--right" size="large">
              <DocumentRemove />
            </el-icon></el-button>
        </el-space>
      </el-aside>
      <el-main>
        <div v-if="isViewTs()">
          <PdfView />
        </div>
        <div v-else>
          <PdfMove />
        </div>
      </el-main>
    </el-container>
  </div>

</template>


<style>
.space {
  /* margin-top: 10%; */
  width: 100%;
}



.space button {
  padding-top: 0% !important;
  margin-top: -5% !important;
  height: 50pt;
  width: 200px;
}

.side {
  background-color: rgba(151, 174, 192, 0.25);
  height: 100vh;
  position: relative;
}

/* Cible les labels des onglets non actifs */
.el-tabs__header .el-tabs__item {
  color: #f7f7f7 !important;
  /* Couleur personnalisée pour les labels des onglets non actifs */
}

/* Cible le label de l'onglet actif */
.el-tabs__header .el-tabs__item.is-active {
  color: #dd9d27 !important;
  /* Couleur personnalisée pour le label de l'onglet actif */
}

/* Cible les labels des onglets lorsqu'ils sont survolés */
.el-tabs__header .el-tabs__item:hover {
  color: #f0a315 !important;
  /* Couleur personnalisée pour les labels des onglets lorsqu'ils sont survolés */
}
</style>