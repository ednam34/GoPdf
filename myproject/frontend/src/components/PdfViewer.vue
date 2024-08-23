<template>
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
                <VuePDF :pdf="pdf" :page="page" :scale="0.5" :key="pdfKey" />
            </div>
            <div id="input" class="input-box">
                <el-button type="warning" @click="removePage" plain>
                    <el-icon>
                        <CloseBold />
                    </el-icon>
                </el-button>
                <el-button type="warning" @click="nextPage" plain>
                    <el-icon><Select /></el-icon>
                </el-button>
                <el-button type="warning" @click="OpenFile" plain>
                    Open File
                </el-button>
            </div>

        </div>
    </div>
    <div class="deletePages">
        <el-button type="warning" @click="removeP" plain>
            Remove pages
        </el-button>
        <el-button type="warning" @click="saveFile" plain>
            Save the file
        </el-button>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { VuePDF, usePDF } from '@tato30/vue-pdf';
import { ArrowLeftBold, ArrowRightBold, Select, CloseBold } from '@element-plus/icons-vue';
import { OpenSinglePdf, PrintArr, PrintString, PrintT, RemovePages, SaveModifiedPDF } from '../../wailsjs/go/main/App';

const page = ref(1);
const pageToDelete = ref<number[]>([]);
const pdfPath = ref('src/test.pdf'); // Chemin initial du PDF
const pdfKey = ref(0); // Clé pour forcer le rechargement

const { pdf, pages } = usePDF(pdfPath);



function saveFile () {
    SaveModifiedPDF(pdfPath.value)
};

// Méthode pour aller à la page précédente
const prevPage = () => {
    if (page.value > 1) {
        page.value -= 1;
        PrintT(page.value);
    }
};

// Méthode pour aller à la page suivante
const nextPage = () => {
    if (page.value < pages.value) {
        page.value += 1;
        PrintT(page.value);
    }
};


function OpenFile() {
    OpenSinglePdf()
        .then(t => {
            if (t) {
                pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
                page.value = 1;
                pdfKey.value++;
                PrintT(t)
                pageToDelete.value = [];
            }
        })
        .catch(error => {
            console.error("Erreur lors de l'ouverture du PDF:", error);
        });
}

function removeP() {
    RemovePages(pageToDelete.value, pdfPath.value).then(path =>{
    pdfPath.value = path; // Mettre à jour le chemin du PDF
            page.value = 1;    // Réinitialiser à la première page
            pdfKey.value++;    // Incrémenter la clé pour forcer le rechargement du composant
            pageToDelete.value = []
    })
}

const removePage = () => {
    if (!pageToDelete.value.includes(page.value)) {
        pageToDelete.value.push(page.value);
    }
    PrintArr(pageToDelete.value);
    if (page.value < pages.value) {
        page.value += 1;

    }
};
</script>

<style scoped>
.pdf-container {
    margin-top: 2%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 80vh;
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
