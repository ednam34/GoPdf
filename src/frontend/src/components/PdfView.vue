<template>
    <main class="test2">
        <div class="pdf-viewer">
            <div class="pdf-container">
                <div class="pdf-controls">
                    <el-button type="warning" @click="prevPage" text>
                        <el-icon>
                            <ArrowLeftBold />
                        </el-icon>
                    </el-button>
                    <span>{{ page }} / {{ pages }}</span>
                    <el-button type="warning" @click="nextPage" text>
                        <el-icon>
                            <ArrowRightBold />
                        </el-icon>
                    </el-button>
                </div>
                <div class="pdf">
                    <!-- Utiliser une clé pour forcer le rechargement du composant -->
                    <VuePDF :pdf="pdf" :page="page" :scale="0.6" :key="pdfKey" />
                </div>
            </div>
        </div>
    </main>
</template>



<script lang="ts" setup>
import { ref } from 'vue';
import { VuePDF, usePDF } from '@tato30/vue-pdf';
import { ArrowLeftBold, ArrowRightBold, Select, CloseBold } from '@element-plus/icons-vue';
import { OpenSinglePdf, PrintArr, PrintString, PrintAny, RemovePages, SaveModifiedPDF } from '../../wailsjs/go/Application/App';
import { page,pageToDelete,pdfKey,pdfPath } from '../shared';


const { pdf, pages } = usePDF(pdfPath);


const prevPage = () => {
    if (page.value > 1) {
        page.value -= 1;
        PrintAny(page.value);
    }
};

// Méthode pour aller à la page suivante
const nextPage = () => {
    if (page.value < pages.value) {
        page.value += 1;
        PrintAny(page.value);
    }
}
</script>


<style scoped>

.test2{
    /* background-color: aliceblue; */
}
.pdf-container {
    margin-top: 2%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 90vh;
    overflow: auto;
}

.pdf {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
}

.pdf-controls {
    margin-bottom: 20px;
}


.deletePages {
    margin-top: 20px;
}
</style>
