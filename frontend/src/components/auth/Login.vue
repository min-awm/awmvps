<template>
  <div
    class="flex items-center justify-center min-h-screen p-4 bg-gradient-to-br from-indigo-100 via-blue-50 to-pink-100 sm:p-6 lg:p-8"
  >
    <div
      class="w-full max-w-md p-6 space-y-8 transition-all transform bg-white shadow-2xl sm:p-8 rounded-2xl"
    >
      <div>
        <h2 class="mt-6 text-3xl font-extrabold text-center text-gray-900">
          Chào mừng trở lại
        </h2>
        <p class="mt-2 text-center text-gray-600">Đăng nhập để tiếp tục</p>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="-space-y-px rounded-md shadow-sm">
          <div>
            <label for="username" class="sr-only">Tên đăng nhập</label>
            <input
              id="username"
              name="username"
              type="text"
              required
              class="relative block w-full px-3 py-3 text-gray-900 placeholder-gray-500 transition-all duration-300 ease-in-out border border-gray-300 appearance-none rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
              placeholder="Nhập tên đăng nhập"
              v-model="username"
            />
          </div>
          <div>
            <label for="password" class="sr-only">Mật khẩu</label>
            <input
              id="password"
              name="password"
              type="password"
              required
              class="relative block w-full px-3 py-3 text-gray-900 placeholder-gray-500 transition-all duration-300 ease-in-out border border-gray-300 appearance-none rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10"
              placeholder="Mật khẩu"
              v-model="password"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="relative flex justify-center w-full px-4 py-3 font-medium text-white border border-transparent rounded-md group bg-gradient-to-r from-indigo-600 to-blue-600"
          >
            Đăng nhập
          </button>
        </div>

        <div class="text-sm text-gray-500">
          <p>Username: admin</p>
          <p>Password: admin</p>
        </div>
      </form>

      <div class="h-3"></div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useToast } from "vue-toastification";
import axios from "@/axios";
import API from "@/api";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/user";

const toast = useToast();
const router = useRouter();
const username = ref("");
const password = ref("");

async function handleSubmit() {
  try {
    const formData = new FormData();
    formData.append("username", username.value);
    formData.append("password", password.value);

    const res = await axios.post(API.LOGIN, formData);
    if (res.success) {
      const userStore = useUserStore();
      userStore.getUserInfo();

      toast.success("Đăng nhập thành công");
      localStorage.setItem("accessToken", res.token);
      router.push({ name: "home" });
    } else {
      switch (res.message) {
        case "userPassIncorrect":
          toast.error("Tài khoản, mật khẩu không đúng");
          break;

        case "serverError":
          toast.error("Lỗi server");
          break;
        default:
          break;
      }
    }
  } catch (error) {
    console.log(`Login: ${error}`);
    toast.error("Lỗi");
  }
}
</script>
