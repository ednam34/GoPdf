import { ref } from 'vue';

export const pdfPath = ref(''); // Chemin initial du PDF

export const page = ref(1);
export const pageToDelete = ref<number[]>([]);
export const pdfKey = ref(0); // Clé pour forcer le rechargement // Chemin initial du PDF
export const isView = ref(true)

