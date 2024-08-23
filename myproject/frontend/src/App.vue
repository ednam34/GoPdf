<script lang="ts" setup>
import { ref } from 'vue';
import type { TabsPaneContext } from 'element-plus'
import HelloWorld from './components/HelloWorld.vue'
import PdfViewer from './components/PdfViewer.vue';
const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>

<template>

<el-tabs
    v-model="activeName"
    type="card"
    class="demo-tabs"
    @tab-click="handleClick"
  >
    <el-tab-pane label="PDF basict tools" name="first"><HelloWorld/></el-tab-pane>
    <el-tab-pane label="PDF viewer"  name="second"><PdfViewer/></el-tab-pane>
  </el-tabs>
  

</template>


<style>

/* Cible les labels des onglets non actifs */
.el-tabs__header .el-tabs__item {
  color: #f7f7f7 !important; /* Couleur personnalisée pour les labels des onglets non actifs */
}

/* Cible le label de l'onglet actif */
.el-tabs__header .el-tabs__item.is-active {
  color: #dd9d27 !important; /* Couleur personnalisée pour le label de l'onglet actif */
}

/* Cible les labels des onglets lorsqu'ils sont survolés */
.el-tabs__header .el-tabs__item:hover {
  color: #f0a315 !important; /* Couleur personnalisée pour les labels des onglets lorsqu'ils sont survolés */
}


</style>