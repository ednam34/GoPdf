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
                <div v-if="pdf" class="pdf">
                    <!-- Utiliser une clé pour forcer le rechargement du composant -->
                    <VuePDF :pdf="pdf" :page="page" :scale="0.6" :key="pdfKey" />

                    <!-- <VuePDF  v-for="p in pages" :pdf="pdf" :page="p" :scale="0.1" :key="pdfKey" /> -->
                </div>
                <div v-else class="pdf">
                    <div class="pdf-upload-content">
                        <el-icon class="el-icon--upload" size="50">
                            <upload-filled />
                        </el-icon>
                        <div class="el-upload__text">
                            <h1>Drop file here or click on the <em>open</em> button</h1>
                        </div>
                    </div>
                </div>
                <div class="pdf-controls">
                <el-button plain @click="MoovePageBack">Moove this pages back</el-button>
                <el-button plain @click="MoovePage">Moove this pages front</el-button>
                </div>
            </div>
        </div>
         
    </main>
</template>



<script lang="ts" setup>
import { VuePDF, usePDF } from '@tato30/vue-pdf';
import { ArrowLeftBold, ArrowRightBold } from '@element-plus/icons-vue';
import { PrintAny } from '../../wailsjs/go/Application/App';
import { page, pdfKey, pdfPath } from '../shared';
import { UploadFilled } from '@element-plus/icons-vue'
import { MoovePage,MoovePageBack } from '../file';

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

.pdf-upload-content {
  display: flex;
  flex-direction: column; /* Place les éléments en colonne */
  align-items: center;    /* Centre les éléments horizontalement */
  justify-content: center; /* Centre les éléments verticalement */
  text-align: center;
}

.el-icon--upload {
  margin-bottom: 10px; /* Ajoute un espace entre l'icône et le texte */
}

.el-upload__text h1 {
  font-size: 16px;
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