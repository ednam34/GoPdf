<template>
    <main class="test2">
        <h1>Moove the pages</h1>
        <div class="pdf-pages">
            <draggable v-model="arr" tag="ul" class="pdf-grid" :key="pdfKey">
                <template #item="{ element: p }">
                    <li class="pdf-item">
                        <VuePDF :pdf="pdf" :page="p" :scale="0.3" />
                        <el-button type="danger" :icon="Close" class="delete-button" @click="RemovePgWithPosition(p)" circle />
                    </li>
                </template>
            </draggable>
        </div>

        <el-button plain @click="Save" class="saveButton">Save</el-button>
    </main>
</template>



<script lang="ts" setup>
import { VuePDF, usePDF } from '@tato30/vue-pdf';
import { PrintArr } from '../../wailsjs/go/Application/App';
import { pdfKey, pdfPath } from '../shared';
import { Reorganize, RemovePgWithPosition } from '../file';
import draggable from 'vuedraggable'
import { ref, watchEffect } from 'vue';
import {Close} from '@element-plus/icons-vue'

const { pdf, pages } = usePDF(pdfPath);

const arr = ref<number[]>([]);


watchEffect(() => {
    if (pages.value) {
        arr.value = Array.from({ length: pages.value }, (_, i) => i + 1);
    }
});


function Save() {
    PrintArr(arr.value)
    Reorganize(arr.value)
}


</script>

<style scoped>

.saveButton {
    margin-top: 20px; /* Un simple espace en haut */
    position: relative; /* Le bouton bougera avec la page */
}

.pdf-pages {
    height: 75vh; /* Ajuste la hauteur pour qu'elle prenne environ 80% de la hauteur de l'écran */
    overflow-y: scroll; /* Ajoute une barre de défilement verticale */
    padding-right: 10px; /* Optionnel : Ajoute de l'espace pour la scrollbar */
}
/* Custom Scrollbar Styles */
.pdf-pages::-webkit-scrollbar {
  width: 7px; /* Largeur de la scrollbar */
}

/* Track */
.pdf-pages::-webkit-scrollbar-track {
  background: #f1f1f1; /* Couleur de l'arrière-plan du track */
}
 
/* Handle */
.pdf-pages::-webkit-scrollbar-thumb {
  background: rgb(143, 164, 209);
}

/* Handle on hover */
.pdf-pages::-webkit-scrollbar-thumb:hover {
    background: rgb(143, 164, 209); 
}


.pdf-grid {
    display: flex;
    flex-wrap: wrap;
    gap: 16px;
    list-style-type: none;
    padding: 0;
    margin: 0;
}

.pdf-item {
    width: calc(33.33% - 16px);
    box-sizing: border-box;
    position: relative;
}

.pdf-item>div {
    display: flex;
    justify-content: center;
    align-items: center;
}

.delete-button {
    position: absolute; /* Superpose le bouton sur la page */
    top: 10px; /* Ajuste la position en haut */
    right: 30px; /* Ajuste la position à droite */
    z-index: 1; /* Assure que le bouton est au-dessus du contenu de la page */
}
</style>