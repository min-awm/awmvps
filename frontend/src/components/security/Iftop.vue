<template>
  <div class="m-2 bg-white border rounded-md p-9">
    <!-- Certificate Auto-renewal Section -->
    <section class="mb-8">
      <h2 class="flex justify-between mb-4 text-xl font-bold text-gray-900">
        Theo dõi truy cập

        <OpenTerminalButton />
      </h2>
      <div class="flex flex-col gap-3 mb-4">
        <p class="text-gray-900">Cài đặt iftop:</p>
        <Code
          command="sudo apt install iftop"
          v-if="packageManager == 'not' || packageManager == 'apt'"
        />
        <Code
          command="sudo yum install iftop"
          v-if="packageManager == 'not' || packageManager == 'yum'"
        />
      </div>

      <div class="mb-4">
        <p class="mb-2 text-gray-900">Chạy iftop:</p>
        <div class="flex flex-col gap-4">
          <Code command="sudo iftop" />
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from "vue";
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import Code from "@/components/common/Code.vue";
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
