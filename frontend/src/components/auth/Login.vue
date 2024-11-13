<template>
  <div
    class="min-h-screen bg-gradient-to-br from-indigo-100 via-blue-50 to-pink-100 flex items-center justify-center p-4 sm:p-6 lg:p-8"
  >
    <div
      class="max-w-md w-full space-y-8 bg-white p-6 sm:p-8 rounded-2xl shadow-2xl transform transition-all"
    >
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Chào mừng trở lại
        </h2>
        <p class="mt-2 text-center text-gray-600">
          Đăng nhập để tiếp tục
        </p>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="username" class="sr-only">Tên đăng nhập</label>
            <input
              id="username"
              name="username"
              type="text"
              required
              class="appearance-none rounded-t-md relative block w-full px-3 py-3 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 transition-all duration-300 ease-in-out"
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
              class="appearance-none rounded-b-md relative block w-full px-3 py-3 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 transition-all duration-300 ease-in-out"
              placeholder="Mật khẩu"
              v-model="password"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="group relative w-full flex justify-center py-3 px-4 border border-transparent font-medium rounded-md text-white bg-gradient-to-r from-indigo-600 to-blue-600"
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

const username = ref("");
const password = ref("");

const handleSubmit = () => {
  // Here you would typically handle the login logic
  console.log("Login attempted", {
    username: username.value,
    password: password.value,
  });
};
</script>
