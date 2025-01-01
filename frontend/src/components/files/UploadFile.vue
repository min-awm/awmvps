<template>
  <label class="p-2 rounded-md hover:bg-gray-100">
    <CloudUpload class="w-5 h-5 text-gray-600" />
    <input type="file" @change="handleFileChange" hidden />
  </label>
</template>

<script setup>
import { CloudUpload } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";

const props = defineProps(["path"]);
const emit = defineEmits(['uploadDone'])
const toast = useToast();

async function handleFileChange(event) {
  try {
    const formData = new FormData();
    formData.append("file", event.target.files[0]);
    formData.append("path", props.path);
    const res = await axios.post(API.UPLOAD_FILE, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
    
    if (res.success) {
        toast.success(res.message)
        emit("uploadDone")
    } else {
        toast.error(res.message)
    }
  } catch (error) {
    console.log(`Upload file: ${error}`);
  }
}
</script>
