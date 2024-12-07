<template>
  <div
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black bg-opacity-50"
    @click="closeModal"
  >
    <!-- Modal Content -->
    <div class="w-full max-w-md bg-white rounded-lg shadow-xl" @click.stop>
      <div class="p-6">
        <h2 class="mb-4 text-2xl font-bold text-gray-800">Đổi tên</h2>
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div>
              <label
                for="currentPassword"
                class="block text-sm font-medium text-gray-700"
              >
                Tên mới
              </label>
              <input
                v-model="newName"
                type="text"
                required
                class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
          <div class="flex items-center justify-end mt-6 space-x-3">
            <button
              type="button"
              @click="closeModal"
              class="px-4 py-2 text-sm font-medium text-gray-700 border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Hủy
            </button>
            <button
              type="submit"
              class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              :disabled="isSubmitting"
            >
              <span v-if="isSubmitting" class="flex items-center">
                <Loader2Icon class="w-4 h-4 mr-2 -ml-1 animate-spin" />
                Đang xử lý...
              </span>
              <span v-else>Xác nhận</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Loader2Icon } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";

const emit = defineEmits(["close"]);
const toast = useToast();

const newName = ref("");
const isSubmitting = ref(false);

function closeModal() {
  emit("close");
  resetForm();
}

async function handleSubmit() {
  isSubmitting.value = true;

  try {
    const formData = new FormData();
    formData.append("newUsername", username.value);

    const res = await axios.post(API.CHANGE_USER_INFO, formData);

    if (res.success) {
      toast.success("Đã cập nhập thành công");
      closeModal();
      userStore.logout();
    } else {
      switch (res.message) {
        case "passIncorrect":
          toast.error("Mật khẩu không đúng");
          break;

        case "serverError":
          toast.error("Lỗi server");
          break;
        default:
          break;
      }
    }
  } catch (error) {
    console.log(`Change user info: ${error}`);
    toast.error("Lỗi");
  } finally {
    isSubmitting.value = false;
  }
}
</script>
