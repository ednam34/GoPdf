<template>
    <main class="test2">
        <h1>Moove the pages</h1>

        <draggable v-model="arr" tag="ul" class="pdf-grid" :key="pdfKey">
            <template #item="{ element: p }">
                <li class="pdf-item">
                    <VuePDF :pdf="pdf" :page="p" :scale="0.3" />
                </li>
            </template>
        </draggable>

        <el-button plain @click="Save">Save</el-button>
    </main>
</template>



<script lang="ts" setup>
import { VuePDF, usePDF } from '@tato30/vue-pdf';
import { ArrowLeftBold, ArrowRightBold } from '@element-plus/icons-vue';
import { PrintAny, PrintArr, SaveModifiedPDF } from '../../wailsjs/go/Application/App';
import { page, pdfKey, pdfPath } from '../shared';
import { UploadFilled } from '@element-plus/icons-vue'
import { MoovePage, MoovePageBack, Reorganize } from '../file';
import draggable from 'vuedraggable'
import { ref, watchEffect } from 'vue';

const { pdf, pages } = usePDF(pdfPath);

const arr = ref<number[]>([]);


watchEffect(() => {
    if (pages.value) {
        arr.value = Array.from({ length: pages.value }, (_, i) => i + 1);
    }
});


function Save(){
    PrintArr(arr.value)
    Reorganize(arr.value)
}


</script>

<style scoped>
.pdf-grid {
    display: flex;
    flex-wrap: wrap; /* Permet de passer à la ligne suivante lorsque la largeur est atteinte */
    gap: 16px; /* Espace entre les éléments */
    list-style-type: none; /* Enlève les puces de la liste */
    padding: 0;
    margin: 0;
}

.pdf-item {
    width: calc(33.33% - 16px); /* Ajuster la largeur en fonction du nombre de colonnes désirées */
    box-sizing: border-box; /* Assure que padding et border sont inclus dans la largeur */
}

.pdf-item > div {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>