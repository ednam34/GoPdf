

import { page, pageToDelete, pdfKey, pdfPath, isView } from './shared';
import { ImgToPdf, MergePdf, MoovePdfPage, OpenSinglePdf, OpenSinglePdfFromPath, OptimizePdf, PrintAny, RemovePages, SaveModifiedPDF, ReorderPdf } from '../wailsjs/go/Application/App';


export function OpenFile() {
    OpenSinglePdf()
        .then(t => {
            if (t) {
                pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
                page.value = 1;
                pdfKey.value++;
                PrintAny(t)
                pageToDelete.value = [];
            }
        })
        .catch(error => {
            console.error("Erreur lors de l'ouverture du PDF:", error);
        });
}

export function OpenFileFromPath(path:string) {
    OpenSinglePdfFromPath(path)
        .then(t => {
            if (t) {
                pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
                page.value = 1;
                pdfKey.value++;
                PrintAny(t)
                pageToDelete.value = [];
            }
        })
        .catch(error => {
            console.error("Erreur lors de l'ouverture du PDF:", error);
        });
}

export function Compress() {
    OptimizePdf().then(t => {
        if (t) {
            pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
            page.value = 1;
            pdfKey.value++;
            PrintAny(t)
            pageToDelete.value = [];
        }
    }).catch(error => {
        console.error("Erreur lors de l'ouverture du PDF:", error);
    });
}


export function Merge() {
    MergePdf().then(t => {
        if (t) {
            pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
            page.value = 1;
            pdfKey.value++;
            PrintAny(t)
            pageToDelete.value = [];
        }
    })
        .catch(error => {
            console.error("Erreur lors de l'ouverture du PDF:", error);
        });
}

export function SavePdf() {
    SaveModifiedPDF(pdfPath.value)
}


export function ImageConv() {
    ImgToPdf().then(t => {
        if (t) {
            pdfPath.value = `${t}`; // Mettre à jour le chemin du PDF
            page.value = 1;
            pdfKey.value++;
            PrintAny(t)
            pageToDelete.value = [];
        }
    }).catch(error => {
        console.error("Erreur lors de l'ouverture du PDF:", error);
    });
}


export function RemovePg() {
    
    RemovePages([page.value], pdfPath.value).then(path => {
        pdfPath.value = path; // Mettre à jour le chemin du PDF
        pdfKey.value++;    // Incrémenter la clé pour forcer le rechargement du composant
        pageToDelete.value = []
    })

}

export function Reorganize(order:number[]){
    ReorderPdf(pdfPath.value,order)
}

export function MoovePage(){
    MoovePdfPage(pdfPath.value,page.value).then(path => {
        pdfPath.value = path; // Mettre à jour le chemin du PDF
        pdfKey.value++;    // Incrémenter la clé pour forcer le rechargement du composant
        pageToDelete.value = []
    })
}


export function MoovePageBack(){
    MoovePdfPage(pdfPath.value,page.value-1).then(path => {
        pdfPath.value = path; // Mettre à jour le chemin du PDF
        pdfKey.value++;    // Incrémenter la clé pour forcer le rechargement du composant
        pageToDelete.value = []
        page.value--;

    })
}
