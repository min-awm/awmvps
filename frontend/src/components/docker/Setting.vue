<template>
  <div class="flex flex-col flex-1 overflow-hidden">
    <div class="p-4 sm:p-6">
      <div class="flex items-center gap-2 mb-6">
        <button
          class="flex items-center text-sm text-blue-600 hover:text-blue-700 md:hidden"
          @click="$emit('openSideBarFn')"
        >
          <ChevronRight class="mr-1" />
        </button>
        <h1 class="text-xl font-semibold text-gray-900">Gỡ cài đặt</h1>
      </div>

      <div class="my-8">
        <div class="flex items-center gap-2 mb-8">
          <span class="text-gray-900">Link:</span>
          <a
            href="https://docs.docker.com/engine/install/ubuntu/#uninstall-docker-engine"
            target="_blank"
            rel="noopener noreferrer"
            v-if="packageManager == 'not' || packageManager == 'apt'"
          >
            https://docs.docker.com/engine/install/ubuntu/#uninstall-docker-engine
          </a>

          <a
            href="https://docs.docker.com/engine/install/ubuntu/#uninstall-docker-engine"
            target="_blank"
            rel="noopener noreferrer"
            v-if="packageManager == 'yum'"
          >
            https://docs.docker.com/engine/install/fedora/#uninstall-docker-engine
          </a>
        </div>

        <OpenTerminalButton />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import axios from "@/axios";
import API from "@/api";

const packageManager = ref("not");
getPackageManager();

async function getPackageManager() {
  try {
    const res = await axios.get(`${API.PACKAGE_MANAGER}`);
    if (res.success) {
      packageManager.value = res.data;
    }
  } catch (error) {
    console.log(`Ssl: ${error}`);
  }
}
</script>
