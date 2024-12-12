<template>
  <div class="p-3 m-2 bg-white border rounded-md" v-if="showInstall">
    <div class="p-6">
      <h2 class="mb-4 text-2xl font-semibold">Docker</h2>

      <div class="flex items-center space-x-4">
        <div class="w-4 h-4 bg-gray-500 rounded-full"></div>
        <span class="text-lg">Chưa cài đặt</span>
      </div>

      <div class="my-8">
        <div class="flex items-center gap-2 mb-2">
          <span class="text-gray-900">Tải từ</span>
          <code class="bg-gray-200 px-2 py-0.5 rounded text-gray-800 text-sm">
            https://get.docker.com
          </code>
          <span class="text-gray-900">:</span>
        </div>
        <Code
          command="curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh"
        />
      </div>

      <OpenTerminalButton />
    </div>
  </div>

  <div class="flex min-h-screen m-2 bg-white border rounded-md" v-else>
    <div class="flex flex-shrink-0">
      <div
        class="fixed top-0 bottom-0 left-0 right-0 md:hidden"
        @click="openSideBar = false"
        v-if="openSideBar"
      ></div>
      <div
        class="fixed top-0 bottom-0 left-0 w-64 bg-gray-100 md:static md:!block"
        v-show="openSideBar"
      >
        <nav class="px-2 mt-5 space-y-1 nav">
          <RouterLink
            to="/docker/container"
            class="flex items-center px-2 py-2 text-sm font-medium text-gray-900 rounded-md hover:bg-gray-200 group"
            @click="openSideBar = false"
          >
            <Box class="w-5 h-5 mr-3 text-gray-500" />
            Containers
          </RouterLink>
          <RouterLink
            to="/docker/image"
            class="flex items-center px-2 py-2 text-sm font-medium text-gray-600 rounded-md hover:bg-gray-200 group"
            @click="openSideBar = false"
          >
            <Package class="w-5 h-5 mr-3 text-gray-500" />
            Images
          </RouterLink>
          <RouterLink
            to="/docker/volume"
            class="flex items-center px-2 py-2 text-sm font-medium text-gray-600 rounded-md hover:bg-gray-200 group"
            @click="openSideBar = false"
          >
            <Database class="w-5 h-5 mr-3 text-gray-500" />
            Volumes
          </RouterLink>
          <RouterLink
            to="/docker/setting"
            class="flex items-center px-2 py-2 text-sm font-medium text-gray-600 rounded-md hover:bg-gray-200 group"
            @click="openSideBar = false"
          >
            <Settings class="w-5 h-5 mr-3 text-gray-500" />
            Cài đặt
          </RouterLink>
        </nav>
      </div>
    </div>

    <!-- Main content -->
    <RouterView @openSideBarFn="openSideBar = true" />
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Box, Database, Package, Settings } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import Code from "@/components/common/Code.vue";

const openSideBar = ref(false);
const showInstall = ref(true);

checkStatus();

async function checkStatus() {
  try {
    const res = await axios.get(`${API.STATUS_DOCKER}`);
    if (res.success) {
      showInstall.value = false;
    }
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}
</script>

<style scoped lang="scss">
.nav {
  .router-link-active {
    background-color: #e5e7eb;
  }
}
.btn-action {
  @apply flex items-center justify-center px-4 py-2 text-white rounded-lg transition-colors duration-200;
}

.btn-action:disabled {
  @apply opacity-50 cursor-not-allowed;
}
</style>
