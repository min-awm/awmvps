<template>
  <div
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black bg-opacity-50"
    @click="closeModal"
  >
    <!-- Modal Content -->
    <div class="w-full max-w-md bg-white rounded-lg shadow-xl" @click.stop>
      <div class="p-6">
        <h2 class="mb-4 text-2xl font-bold text-gray-800">
          Đổi thông tin tài khoản
        </h2>
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div>
              <label
                for="currentPassword"
                class="block text-sm font-medium text-gray-700"
              >
                Mật khẩu hiện tại
              </label>
              <input
                id="currentPassword"
                v-model="currentPassword"
                type="password"
                required
                class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.currentPassword }"
              />
              <p
                v-if="errors.currentPassword"
                class="mt-1 text-sm text-red-600"
              >
                {{ errors.currentPassword }}
              </p>
            </div>
            <div>
              <label
                for="username"
                class="block text-sm font-medium text-gray-700"
              >
                Tên đăng nhập mới
              </label>
              <input
                id="username"
                v-model="username"
                type="text"
                required
                class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.username }"
              />
              <p v-if="errors.username" class="mt-1 text-sm text-red-600">
                {{ errors.username }}
              </p>
            </div>

            <div>
              <label
                for="newPassword"
                class="block text-sm font-medium text-gray-700"
              >
                Mật khẩu mới
              </label>
              <input
                id="newPassword"
                v-model="newPassword"
                type="password"
                required
                class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.newPassword }"
              />
              <p v-if="errors.newPassword" class="mt-1 text-sm text-red-600">
                {{ errors.newPassword }}
              </p>
            </div>
            <div>
              <label
                for="confirmPassword"
                class="block text-sm font-medium text-gray-700"
              >
                Xác nhận mật khẩu mới
              </label>
              <input
                id="confirmPassword"
                v-model="confirmPassword"
                type="password"
                required
                class="block w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                :class="{ 'border-red-500': errors.confirmPassword }"
              />
              <p
                v-if="errors.confirmPassword"
                class="mt-1 text-sm text-red-600"
              >
                {{ errors.confirmPassword }}
              </p>
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
import { ref, reactive } from "vue";
import { Loader2Icon } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import { useToast } from "vue-toastification";
import { useUserStore } from "@/store/user";

const emit = defineEmits(["close"]);
const toast = useToast();
const userStore = useUserStore();

const username = ref("");
const currentPassword = ref("");
const newPassword = ref("");
const confirmPassword = ref("");
const isSubmitting = ref(false);
const errors = reactive({
  username: "",
  currentPassword: "",
  newPassword: "",
  confirmPassword: "",
});

function closeModal() {
  emit("close");
  resetForm();
}

function resetForm() {
  username.value = "";
  currentPassword.value = "";
  newPassword.value = "";
  confirmPassword.value = "";
  Object.keys(errors).forEach((key) => (errors[key] = ""));
}

function validateForm() {
  let isValid = true;

  if (!username.value) {
    errors.username = "Vui lòng nhập tên đăng nhập";
    isValid = false;
  } else {
    errors.username = "";
  }

  if (!currentPassword.value) {
    errors.currentPassword = "Vui lòng nhập mật khẩu hiện tại";
    isValid = false;
  } else {
    errors.currentPassword = "";
  }

  if (!newPassword.value) {
    errors.newPassword = "Vui lòng nhập mật khẩu mới";
    isValid = false;
  } else {
    errors.newPassword = "";
  }

  if (!confirmPassword.value) {
    errors.confirmPassword = "Vui lòng xác nhận mật khẩu mới";
    isValid = false;
  } else if (confirmPassword.value !== newPassword.value) {
    errors.confirmPassword = "Mật khẩu xác nhận không khớp";
    isValid = false;
  } else {
    errors.confirmPassword = "";
  }

  return isValid;
}

async function handleSubmit() {
  if (!validateForm()) return;

  isSubmitting.value = true;

  try {
    const formData = new FormData();
    formData.append("newUsername", username.value);
    formData.append("currentPassword", currentPassword.value);
    formData.append("newPassword", newPassword.value);

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
