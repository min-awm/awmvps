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
        <h1 class="text-xl font-semibold text-gray-900">Volumes</h1>
      </div>

      <!-- Container Table -->
      <div class="px-2 overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Driver
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-3 py-3 text-xs font-medium tracking-wider text-right text-gray-500 uppercase"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(volume, i) in volumes" :key="`volume-${i}`">
              <td class="px-3 py-4">
                <div class="flex items-center">
                  {{ volume.Driver }}
                </div>
              </td>
              <td class="px-3 py-4 text-sm text-gray-500">
                {{ volume.Name }}
              </td>
              <td
                class="flex items-center justify-end gap-3 px-3 py-4 space-x-2 text-right"
              >
                <button
                  class="text-gray-400 hover:text-gray-500"
                  @click="removeVolume(volume.Name)"
                >
                  <Trash2 class="w-4 h-4" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="p-6">
      <h2 class="mb-4 text-2xl font-semibold">Logs</h2>
      <div
        class="h-48 p-4 overflow-y-auto font-mono text-sm text-green-400 bg-gray-900 rounded-lg"
      >
        <p v-for="(log, i) in dataLogs" :key="`volume-log-${i}`">
          {{ log }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { Trash2, ChevronRight } from "lucide-vue-next";
import axios from "@/axios";
import API from "@/api";

const dataLogs = ref([]);
const volumes = ref([]);
getListVolume();

async function getListVolume() {
  try {
    const res = await axios.get(`${API.LIST_VOLUME_DOCKER}`);
    if (res.success) {
      if (res.message == "") {
        volumes.value = [];
      } else {
        const jsonStrings = res.message
          .trim()
          .split("\n")
          .map((line) => line.trim())
          .map((line) => line.slice(1, -1));

        const parsedData = jsonStrings.map(JSON.parse);
        volumes.value = parsedData;
      }
    } else {
      dataLogs.value.push(
        `[${new Date().toLocaleTimeString()}] ${res.message}`
      );
    }
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}

async function removeVolume(id) {
  dataLogs.value.push(
    `[${new Date().toLocaleTimeString()}] Đang xoá volume ${id}`
  );
  try {
    const formData = new FormData();
    formData.append("id", id);
    const res = await axios.post(`${API.REMOVE_VOLUME_DOCKER}`, formData);
    if (res.success) {
      getListVolume();
    }
    dataLogs.value.push(`[${new Date().toLocaleTimeString()}] ${res.message}`);
  } catch (error) {
    console.log(`Docker: ${error}`);
  }
}
</script>
